package main

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
 * This example code tops up an existing validator on the Hoodi network with some ETH.
 * Run the code with 'go run examples/ethereum/dedicated-eth-top-up/main.go <api_key_file_path> <funding_wallet_address> <top-up-validator-pubkey> <private_key> <rpc_url>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)

	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewExternalAddress(coinbase.EthereumHoodi, os.Args[2])

	options := []coinbase.StakingOperationOption{
		coinbase.WithNativeStakingOperationMode(),
		coinbase.WithTopUpValidatorPublicKey(os.Args[3]),
	}

	topUpOperation, err := client.BuildStakeOperation(
		ctx,
		big.NewFloat(1), // Top Up 1 ETH
		coinbase.Eth,
		address,
		options...,
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", topUpOperation.ID())

	if err := client.Wait(ctx, topUpOperation, coinbase.WithWaitTimeoutSeconds(600)); err != nil {
		log.Fatalf("error waiting for staking operation: %v", err)
	}

	// Load your wallet's private key from which you initiated the above top up operation.
	key, err := crypto.HexToECDSA(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Signing top up operation ...")

	// Sign the transactions within staking operation resource with your private key.
	if err := topUpOperation.Sign(key); err != nil {
		log.Fatal(err)
	}

	// For Hoodi, publicly available RPC URL's can be found here https://chainlist.org/chain/560048
	ethClient, err := ethclient.Dial(os.Args[5])
	if err != nil {
		log.Fatal(err)
	}

	// Broadcast each of the signed transactions to the network.
	for index, transaction := range topUpOperation.Transactions() {
		log.Printf("Top up tx %d: %s\n", index+1, transaction.SignedPayload())

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
