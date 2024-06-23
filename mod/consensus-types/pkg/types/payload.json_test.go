package types_test

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestExecutableDataDeneb_MarshalUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   types.ExecutableDataDeneb
		wantErr bool
	}{
		{
			name: "valid data",
			input: types.ExecutableDataDeneb{
				ParentHash:    ethcommon.BytesToHash([]byte("parentHash")),
				FeeRecipient:  ethcommon.BytesToAddress([]byte("feeRecipient")),
				StateRoot:     common.Bytes32(ethcommon.BytesToHash([]byte("stateRoot"))),
				ReceiptsRoot:  common.Bytes32(ethcommon.BytesToHash([]byte("receiptsRoot"))),
				LogsBloom:     make([]byte, 256),
				Random:        common.Bytes32(ethcommon.BytesToHash([]byte("random"))),
				Number:        math.U64(1),
				GasLimit:      math.U64(1000000),
				GasUsed:       math.U64(500000),
				Timestamp:     math.U64(1622740000),
				ExtraData:     []byte("extraData"),
				BaseFeePerGas: math.U256L(ethcommon.BytesToHash([]byte("1000000000"))),
				BlockHash:     ethcommon.BytesToHash([]byte("blockHash")),
				Transactions:  [][]byte{[]byte("tx1"), []byte("tx2")},
				Withdrawals: []*engineprimitives.Withdrawal{
					{
						Index:     math.U64(0),
						Validator: math.U64(1),
						Address:   ethcommon.BytesToAddress([]byte("withdrawalAddress")),
						Amount:    math.U64(1000),
					},
				},
				BlobGasUsed:   math.U64(2000000),
				ExcessBlobGas: math.U64(500000),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal the input object to JSON
			data, err := tt.input.MarshalJSON()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// Unmarshal the JSON back into an ExecutableDataDeneb object
			var unmarshaled types.ExecutableDataDeneb
			err = unmarshaled.UnmarshalJSON(data)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// Verify that the input and unmarshaled objects are equal
			require.Equal(t, tt.input.ParentHash, unmarshaled.ParentHash)
			require.Equal(t, tt.input.FeeRecipient, unmarshaled.FeeRecipient)
			require.Equal(t, common.Bytes32(ethcommon.BytesToHash(tt.input.StateRoot[:])), unmarshaled.StateRoot)
			require.Equal(t, tt.input.ReceiptsRoot, unmarshaled.ReceiptsRoot)
			require.Equal(t, tt.input.LogsBloom, unmarshaled.LogsBloom)
			require.Equal(t, tt.input.Random, unmarshaled.Random)
			require.Equal(t, tt.input.Number, unmarshaled.Number)
			require.Equal(t, tt.input.GasLimit, unmarshaled.GasLimit)
			require.Equal(t, tt.input.GasUsed, unmarshaled.GasUsed)
			require.Equal(t, tt.input.Timestamp, unmarshaled.Timestamp)
			require.Equal(t, tt.input.ExtraData, unmarshaled.ExtraData)
			require.Equal(t, tt.input.BaseFeePerGas, unmarshaled.BaseFeePerGas)
			require.Equal(t, tt.input.BlockHash, unmarshaled.BlockHash)
			require.Equal(t, tt.input.Transactions, unmarshaled.Transactions)
			require.Equal(t, tt.input.Withdrawals, unmarshaled.Withdrawals)
			require.Equal(t, tt.input.BlobGasUsed, unmarshaled.BlobGasUsed)
			require.Equal(t, tt.input.ExcessBlobGas, unmarshaled.ExcessBlobGas)
		})
	}
}
