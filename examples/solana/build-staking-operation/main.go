package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	networkID = client.NETWORKIDENTIFIER_SOLANA_DEVNET
	amount    = big.NewFloat(0.1)
	rpcURL    = "https://api.devnet.solana.com"
)

/*
 * This example code stakes SOL on the devnet network.
 * Run the code with 'go run examples/solana/build-staking-operation/main.go <api_key_file_path> <wallet_address> <wallet_private_key>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
		coinbase.WithTimeout(20*time.Second),
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

	stakingOperation, err = client.Wait(ctx, stakingOperation, coinbase.WithWaitTimeoutSeconds(60))
	if err != nil {
		log.Fatalf("error waiting for staking operation: %v", err)
	}

	for _, transaction := range stakingOperation.Transactions() {
		log.Printf("Unsigned tx payload: %s\n\n", transaction.UnsignedPayload())

		signedTx, err := signSolTransaction(transaction.UnsignedPayload(), []string{os.Args[3]})
		if err != nil {
			log.Fatalf("error signing transaction: %v", err)
		}

		log.Printf("Signed tx: %s\n\n", signedTx)

		broadcastedTx, err := broadcastSolTransaction(ctx, signedTx)
		if err != nil {
			log.Fatalf("error broadcasting transaction: %v", err)
		}

		log.Printf("Broadcasted tx: %s\n\n", getTxLink(stakingOperation.NetworkID(), broadcastedTx))
	}
}

func signSolTransaction(unsignedTx string, privateKeys []string) (string, error) {
	if len(privateKeys) == 0 {
		return "", fmt.Errorf("need to pass at least one private key")
	}

	signers := make([]solana.PrivateKey, 0, len(privateKeys))

	for _, privateKey := range privateKeys {
		signer, err := solana.PrivateKeyFromBase58(privateKey)
		if err != nil {
			return "", fmt.Errorf("error getting private key: %w", err)
		}

		signers = append(signers, signer)
	}

	data := base58.Decode(unsignedTx)

	// parse transaction
	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(data))
	if err != nil {
		return "", err
	}

	// clear signatures: https://github.com/gagliardetto/solana-go/issues/186
	tx.Signatures = nil

	if _, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		for _, candidate := range signers {
			if candidate.PublicKey().Equals(key) {
				return &candidate
			}
		}

		return nil
	}); err != nil {
		return "", fmt.Errorf("error signing transaction: %w", err)
	}

	marshaledTx, err := tx.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("error marshaling transaction: %w", err)
	}

	base58EncodedSignedTx := base58.Encode(marshaledTx)

	return base58EncodedSignedTx, nil
}

func broadcastSolTransaction(ctx context.Context, signedTx string) (string, error) {
	var (
		sig solana.Signature
		err error
	)

	cluster := rpc.Cluster{
		Name: "solana-staking-demo-rpc",
		RPC:  rpcURL,
	}

	rpcClient := rpc.New(cluster.RPC)

	data := base58.Decode(signedTx)

	// parse transaction
	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(data))
	if err != nil {
		return "", err
	}

	opts := rpc.TransactionOpts{
		SkipPreflight:       false,
		PreflightCommitment: rpc.CommitmentFinalized,
	}

	fmt.Println("Sending transaction...")

	maxRetries := 5

	var sig solana.Signature
	for i := 0; i < maxRetries; i++ {
		fmt.Printf("Sending transaction attempt [%d]...\n", i)

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

func getTxLink(networkID, signature string) string {
	if networkID == "solana-mainnet" {
		return fmt.Sprintf("https://explorer.solana.com/tx/%s", signature)
	} else if networkID == "solana-devnet" {
		return fmt.Sprintf("https://explorer.solana.com/tx/%s?cluster=devnet", signature)
	}

	return ""
}
