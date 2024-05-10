// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package core

import (
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/constants"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core/state"
)

// StateProcessor is a basic Processor, which takes care of the
// main state transition for the beacon chain.
type StateProcessor[SidecarsT interface{ Len() int }] struct {
	cs     primitives.ChainSpec
	bp     BlobProcessor[SidecarsT]
	rp     RandaoProcessor
	signer crypto.BLSSigner
	logger log.Logger[any]

	// DepositProcessor
	// WithdrawalProcessor
}

// NewStateProcessor creates a new state processor.
func NewStateProcessor[SidecarsT interface{ Len() int }](
	cs primitives.ChainSpec,
	bp BlobProcessor[SidecarsT],
	rp RandaoProcessor,
	signer crypto.BLSSigner,
	logger log.Logger[any],
) *StateProcessor[SidecarsT] {
	return &StateProcessor[SidecarsT]{
		cs:     cs,
		bp:     bp,
		rp:     rp,
		signer: signer,
		logger: logger,
	}
}

// Transition is the main function for processing a state transition.
func (sp *StateProcessor[SidecarsT]) Transition(
	st state.BeaconState,
	blk types.BeaconBlock,
	/*validateSignature bool, */
	validateResult bool,
) error {
	// Process the slot.
	if err := sp.ProcessSlot(st); err != nil {
		return err
	}

	// Process the block.
	if err := sp.ProcessBlock(st, blk); err != nil {
		return err
	}

	if validateResult {
		stateRoot, err := st.HashTreeRoot()
		if err != nil {
			return err
		}

		if stateRoot != blk.GetStateRoot() {
			return ErrStateRootMismatch
		}
	}

	return nil
}

// ProcessSlot is run when a slot is missed.
func (sp *StateProcessor[SidecarsT]) ProcessSlot(
	st state.BeaconState,
) error {
	slot, err := st.GetSlot()
	if err != nil {
		return err
	}

	// Before we make any changes, we calculate the previous state root.
	prevStateRoot, err := st.HashTreeRoot()
	if err != nil {
		return err
	}

	// We update our state roots and block roots.
	if err = st.UpdateStateRootAtIndex(
		uint64(slot)%sp.cs.SlotsPerHistoricalRoot(),
		prevStateRoot,
	); err != nil {
		return err
	}

	// We get the latest block header, this will not have
	// a state root on it.
	latestHeader, err := st.GetLatestBlockHeader()
	if err != nil {
		return err
	}

	// We set the "rawHeader" in the StateProcessor, but cannot fill in
	// the StateRoot until the following block.
	if (latestHeader.StateRoot == primitives.Root{}) {
		latestHeader.StateRoot = prevStateRoot
		if err = st.SetLatestBlockHeader(latestHeader); err != nil {
			return err
		}
	}

	// We update the block root.
	var prevBlockRoot primitives.Root
	prevBlockRoot, err = latestHeader.HashTreeRoot()
	if err != nil {
		return err
	}

	if err = st.UpdateBlockRootAtIndex(
		uint64(slot)%sp.cs.SlotsPerHistoricalRoot(), prevBlockRoot,
	); err != nil {
		return err
	}

	// Process the Epoch Boundary.
	if uint64(slot+1)%sp.cs.SlotsPerEpoch() == 0 {
		if err = sp.processEpoch(st); err != nil {
			return err
		}
		sp.logger.Info(
			"processed epoch transition ⏰ ",
			"old", uint64(slot)/sp.cs.SlotsPerEpoch(),
			"new", uint64(slot+1)/sp.cs.SlotsPerEpoch(),
		)
	}

	return st.SetSlot(slot + 1)
}

// ProcessBlobs processes the blobs and ensures they match the local state.
func (sp *StateProcessor[SidecarsT]) ProcessBlobs(
	st state.BeaconState,
	avs AvailabilityStore[types.ReadOnlyBeaconBlockBody, SidecarsT],
	sidecars SidecarsT,
) error {
	slot, err := st.GetSlot()
	if err != nil {
		return err
	}
	return sp.bp.ProcessBlobs(slot, avs, sidecars)
}

// ProcessBlock processes the block and ensures it matches the local state.
func (sp *StateProcessor[SidecarsT]) ProcessBlock(
	st state.BeaconState,
	blk types.BeaconBlock,
) error {
	// process the freshly created header.
	if err := sp.processHeader(st, blk); err != nil {
		return err
	}

	// process the withdrawals.
	body := blk.GetBody()
	if err := sp.processWithdrawals(
		st, body.GetExecutionPayload(),
	); err != nil {
		return err
	}

	// phase0.ProcessProposerSlashings
	// phase0.ProcessAttesterSlashings

	// process the randao reveal.
	if err := sp.processRandaoReveal(st, blk); err != nil {
		return err
	}

	// phase0.ProcessEth1Vote ? forkchoice?

	// TODO: LOOK HERE
	//
	// process the deposits and ensure they match the local state.
	if err := sp.processOperations(st, body); err != nil {
		return err
	}

	// ProcessVoluntaryExits

	return nil
}

// processEpoch processes the epoch and ensures it matches the local state.
func (sp *StateProcessor[SidecarsT]) processEpoch(st state.BeaconState) error {
	var err error
	if err = sp.processRewardsAndPenalties(st); err != nil {
		return err
	}
	if err = sp.processSlashingsReset(st); err != nil {
		return err
	}
	if err = sp.processRandaoMixesReset(st); err != nil {
		return err
	}
	return nil
}

// processHeader processes the header and ensures it matches the local state.
func (sp *StateProcessor[SidecarsT]) processHeader(
	st state.BeaconState,
	blk types.BeaconBlock,
) error {
	// TODO: this function is really confusing, can probably just
	// be removed and the logic put in the ProcessBlock function.
	header := blk.GetHeader()
	if header == nil {
		return ErrNilBlockHeader
	}

	// Store as the new latest block
	headerRaw := &types.BeaconBlockHeader{
		BeaconBlockHeaderBase: types.BeaconBlockHeaderBase{
			Slot:            header.Slot,
			ProposerIndex:   header.ProposerIndex,
			ParentBlockRoot: header.ParentBlockRoot,
			// state_root is zeroed and overwritten in the next `process_slot`
			// call.
			// with BlockHeaderState.UpdateStateRoot(), once the post state is
			// available.
			StateRoot: [32]byte{},
		},
		BodyRoot: header.BodyRoot,
	}
	return st.SetLatestBlockHeader(headerRaw)
}

// processRandaoReveal processes the randao reveal and
// ensures it matches the local state.
func (sp *StateProcessor[SidecarsT]) processRandaoReveal(
	st state.BeaconState,
	blk types.BeaconBlock,
) error {
	return sp.rp.ProcessRandao(st, blk)
}

// processRandaoMixesReset as defined in the Ethereum 2.0 specification.
// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#randao-mixes-updates
//
//nolint:lll
func (sp *StateProcessor[SidecarsT]) processRandaoMixesReset(
	st state.BeaconState,
) error {
	return sp.rp.ProcessRandaoMixesReset(st)
}

// getAttestationDeltas as defined in the Ethereum 2.0 specification.
// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#get_attestation_deltas
//
//nolint:lll
func (sp *StateProcessor[SidecarsT]) getAttestationDeltas(
	st state.BeaconState,
) ([]math.Gwei, []math.Gwei, error) {
	// TODO: implement this function forreal
	validators, err := st.GetValidators()
	if err != nil {
		return nil, nil, err
	}
	placeholder := make([]math.Gwei, len(validators))
	return placeholder, placeholder, nil
}

// processRewardsAndPenalties as defined in the Ethereum 2.0 specification.
// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#process_rewards_and_penalties
//
//nolint:lll
func (sp *StateProcessor[SidecarsT]) processRewardsAndPenalties(
	st state.BeaconState,
) error {
	slot, err := st.GetSlot()
	if err != nil {
		return err
	}

	if sp.cs.SlotToEpoch(slot) == math.U64(constants.GenesisEpoch) {
		return nil
	}

	rewards, penalties, err := sp.getAttestationDeltas(st)
	if err != nil {
		return err
	}
	validators, err := st.GetValidators()
	if err != nil {
		return err
	}
	if len(validators) != len(rewards) || len(validators) != len(penalties) {
		return errors.Newf(
			"mismatched rewards and penalties lengths: %d, %d, %d",
			len(validators), len(rewards), len(penalties),
		)
	}
	for i := range validators {
		// Increase the balance of the validator.
		if err = st.IncreaseBalance(
			math.ValidatorIndex(i),
			rewards[i],
		); err != nil {
			return err
		}

		// Decrease the balance of the validator.
		if err = st.DecreaseBalance(
			math.ValidatorIndex(i),
			penalties[i],
		); err != nil {
			return err
		}
	}
	return nil
}