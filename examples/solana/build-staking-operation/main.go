package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	networkID          = "solana-devnet"
	amount             = big.NewFloat(0.1)
	rpcURL             = "https://api.devnet.solana.com"
	defaultPrivKeyPath = filepath.Join(home(), ".config/solana/id.json")
)

func main() {
	ctx := context.Background()

	walletAddress := os.Args[2]

	privKeys := []string{defaultPrivKeyPath}
	if keys := os.Getenv("SOL_PRIVATE_KEYS"); keys != "" {
		privKeys = strings.Split(keys, ",")
	}

	signers := make([]solana.PrivateKey, len(privKeys))
	for i, pk := range privKeys {
		privKey, err := solana.PrivateKeyFromSolanaKeygenFile(pk)
		if err != nil {
			log.Fatalf("private key %s does not exist or is invalid", privKey)
		}

		signers[i] = privKey
	}

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewExternalAddress(networkID, walletAddress)

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

	for _, transaction := range stakingOperation.Transactions() {
		log.Printf("Tx unsigned payload: %s\n\n", transaction.UnsignedPayload())

		signedTx, err := signSolTransaction(transaction.UnsignedPayload(), signers)
		if err != nil {
			log.Fatalf("error signing transaction: %v", err)
		}

		log.Printf("Signed tx: %s\n\n", signedTx)

		sig, err := broadcastSolTransaction(ctx, signedTx)
		if err != nil {
			log.Fatalf("error broadcasting transaction: %v", err)
		}

		log.Printf("Broadcasted tx: %s\n\n", getTxLink(stakingOperation.NetworkID(), sig))
	}
}

func signSolTransaction(unsignedTx string, signers []solana.PrivateKey) (string, error) {
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
		RPC: rpcURL,
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

	maxRetries := 20

	for maxRetries > 0 {
		fmt.Printf("Trying again [%d] Sending transaction...\n", 21-maxRetries)

		sig, err = rpcClient.SendTransactionWithOpts(ctx, tx, opts)
		if err != nil {
			time.Sleep(3 * time.Second)
			maxRetries--

			continue
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

func home() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("unable to get user homedir")
	}

	return home
}
