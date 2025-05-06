package main

import (
	"context"
	"log"
	"os"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
 * This example code does an execution layer orchestrated consolidation b/w 2 ethereum validators on the Hoodi network.
 * Run the code with 'go run examples/ethereum/dedicated-eth-consolidate/main.go <api_key_file_path> <withdrawal_address> <source-val-pubkey> <target-val-pubkey> <private_key> <rpc_url>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)

	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	withdrawalAddress := coinbase.NewExternalAddress(coinbase.EthereumHoodi, os.Args[2])

	options := []coinbase.StakingOperationOption{
		coinbase.WithSourceValidatorPublicKey(os.Args[3]),
		coinbase.WithTargetValidatorPublicKey(os.Args[4]),
	}

	consolidateOperation, err := client.BuildValidatorConsolidationOperation(
		ctx,
		withdrawalAddress,
		options...,
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", consolidateOperation.ID())

	if err := client.Wait(ctx, consolidateOperation, coinbase.WithWaitTimeoutSeconds(600)); err != nil {
		log.Fatalf("error waiting for staking operation: %v", err)
	}

	// Load your wallet's private key from which you initiated the above stake operation.
	key, err := crypto.HexToECDSA(os.Args[5])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Signing consolidate operation ...")

	// Sign the transactions within staking operation resource with your private key.
	if err := consolidateOperation.Sign(key); err != nil {
		log.Fatal(err)
	}

	// For Hoodi, publicly available RPC URL's can be found here https://chainlist.org/chain/560048
	ethClient, err := ethclient.Dial(os.Args[6])
	if err != nil {
		log.Fatal(err)
	}

	// Broadcast each of the signed transactions to the network.
	for index, transaction := range consolidateOperation.Transactions() {
		log.Printf("Consolidate tx %d: %s\n", index+1, transaction.UnsignedPayload())

		rawTx, ok := transaction.Raw().(*types.Transaction)
		if !ok {
			log.Fatal("Failed to cast to *types.Transaction")
		}

		if err := ethClient.SendTransaction(context.Background(), rawTx); err != nil {
			log.Fatal(err)
		}

		log.Printf("Broadcasted transaction hash: %s", rawTx.Hash().Hex())
	}
}
