package main

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	networkID = coinbase.SolanaDevnet
	amount    = big.NewFloat(0.01)
	rpcURL    = "https://api.devnet.solana.com"
)

/*
 * This example code stakes SOL on the devnet network.
 * Run the code with 'go run examples/solana/stake/main.go <api_key_file_path> <wallet_address> <wallet_private_key>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
		coinbase.WithTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewExternalAddress(string(networkID), os.Args[2])

	balance, err := client.GetStakeableBalance(ctx, coinbase.Sol, address)
	if err != nil {
		log.Fatalf("error getting balance: %v", err)
	}

	log.Printf("Stakeable balance: %s\n\n", balance.Amount().String())

	stakingOperation, err := client.BuildStakeOperation(ctx, amount, coinbase.Sol, address)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	log.Printf("Staking operation ID: %s\n\n", stakingOperation.ID())

	err = client.Wait(ctx, stakingOperation, coinbase.WithWaitTimeoutSeconds(60))
	if err != nil {
		log.Fatalf("error waiting for staking operation: %v", err)
	}

	privateKey, err := decodePrivateKey(os.Args[3])

	rpcClient := rpc.New(rpcURL)

	maxRetries := uint(5)
	opts := rpc.TransactionOpts{
		SkipPreflight: false,
		MaxRetries:    &maxRetries,
		// NOTE: In production, consider using rpc.CommitmentFinalized instead to ensure the block is included.
		PreflightCommitment: rpc.CommitmentProcessed,
	}

	for i, stakingTransactionDetail := range stakingOperation.StakingTransactionDetails() {
		log.Printf("metadata: %s", stakingTransactionDetail.DedicatedEthStakeTransactionDetail.Metadata.EthereumValidatorDepositData)

		transaction := stakingOperation.Transaction(i)

		err = stakingOperation.Sign(privateKey)
		if err != nil {
			log.Fatalf("error signing transaction: %v", err)
		}

		unsignedTx := transaction.UnsignedPayload()
		signedTx := transaction.SignedPayload()

		log.Printf("Unsigned tx payload: %s\n\n", unsignedTx)
		log.Printf("Signed tx payload: %s\n\n", signedTx)

		rawTx := transaction.Raw()
		solanaTx, ok := rawTx.(*solana.Transaction)
		if !ok {
			log.Fatal("failed to cast raw transaction to solana.Transaction")
		}

		sig, err := rpcClient.SendTransactionWithOpts(ctx, solanaTx, opts)
		if err != nil {
			log.Fatalf("failed to send transaction: %v", err)
		}

		log.Printf("Broadcasted tx: %s\n\n", getTxLink(stakingOperation.NetworkID(), sig.String()))
	}
}

func decodePrivateKey(privateKeyString string) (*ed25519.PrivateKey, error) {
	// Decode the base58 encoded private key
	privateKeyBytes := base58.Decode(privateKeyString)
	if len(privateKeyBytes) != ed25519.PrivateKeySize {
		log.Fatalf("invalid private key length: expected %d bytes, got %d bytes", ed25519.PrivateKeySize, len(privateKeyBytes))
	}

	// Convert the byte slice to an ed25519 private key
	privateKey := ed25519.PrivateKey(privateKeyBytes)

	return &privateKey, nil
}

func getTxLink(networkID, signature string) string {
	if networkID == coinbase.SolanaMainnet {
		return fmt.Sprintf("https://explorer.solana.com/tx/%s", signature)
	} else if networkID == coinbase.SolanaDevnet {
		return fmt.Sprintf("https://explorer.solana.com/tx/%s?cluster=devnet", signature)
	}

	return ""
}
