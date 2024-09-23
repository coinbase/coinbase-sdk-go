package coinbase

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// dummyEthereumUnsignedPayload creates a new Ethereum dummy transaction with minimal inputs.
// This should be deterministic and not rely on any external state.
func dummyEthereumUnsignedPayload(t *testing.T, nonce uint64) string {
	t.Helper()

	// Create a simple Ethereum transaction
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &common.Address{},
		Value:    big.NewInt(0),
		Gas:      21000,
		GasPrice: big.NewInt(1),
		Data:     nil,
	})

	// Encode the transaction to hex
	txBytes, err := tx.MarshalJSON()
	if err != nil {
		require.NoError(t, err)
	}

	return hex.EncodeToString(txBytes)
}
