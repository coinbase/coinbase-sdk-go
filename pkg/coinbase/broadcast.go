package coinbase

import (
	"context"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil/base58"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// This is a helper function that takes a signed transaction payload and broadcasts it to the Solana network.
// It returns the transaction signature if the transaction was successfully broadcasted.
func BroadcastSolanaTransaction(ctx context.Context, signedTx string, rpcURL string) (string, error) {
	rpcClient := rpc.New(rpcURL)

	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(base58.Decode(signedTx)))
	if err != nil {
		return "", err
	}

	opts := rpc.TransactionOpts{
		SkipPreflight:       false,
		PreflightCommitment: rpc.CommitmentFinalized,
	}

	maxRetries := 5

	var sig solana.Signature
	for i := 0; i < maxRetries; i++ {
		sig, err = rpcClient.SendTransactionWithOpts(ctx, tx, opts)
		if err != nil {
			fmt.Printf("Error sending transaction: %v.", err)
			fmt.Printf("Retrying in 3 seconds...\n")

			time.Sleep(3 * time.Second)

			i++
		}

		break
	}

	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	return sig.String(), nil
}
