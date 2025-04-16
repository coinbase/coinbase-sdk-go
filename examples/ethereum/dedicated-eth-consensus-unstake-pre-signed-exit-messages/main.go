package main

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

var (
	networkID = client.NETWORKIDENTIFIER_ETHEREUM_HOODI
	assetID   = coinbase.Eth
)

/*
 * This example code does a consensus layer unstake on the Hoodi network via Dedicated ETH Staking.
 * Run the code with 'go run examples/ethereum/dedicated-eth-consensus-unstake-pre-signed-exit-messages/main.go <api_key_file_path> <withdrawal_address> <rpc_url>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)

	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewExternalAddress(string(networkID), os.Args[2])

	unstakeableBalance, err := client.GetUnstakeableBalance(ctx, assetID, address, coinbase.WithNativeStakingBalanceMode())
	if err != nil {
		log.Fatalf("error getting unstakeableBalance balance: %v", err)
	}

	log.Printf("unstakeable balance: %s\n", unstakeableBalance.Amount().Text('f', 18))

	builder := coinbase.NewConsensusLayerExitOptionBuilder()
	builder.AddValidator("0x95b3c4521b329c7bb00e972752ef4a888bb046c36137275b0850190e5b5f72b316ec4802fa0d2c0bed943cb47ade9cfe")
	builder.AddValidator("0x8a1487c353f0765c4d6686558720505f119b7de4a18a07b7d2245842ffe17cd3e0fd04e8744f3d0b18f84128a8a81f7e")

	options := []coinbase.StakingOperationOption{
		coinbase.WithNativeStakingOperationMode(),
		coinbase.WithConsensusLayerExit(builder),
		//coinbase.WithImmediateUnstake(), // Enable this if you don't want pre-signed exit messages and want to unstake immediately.
	}

	unstakeOperation, err := client.BuildUnstakeOperation(
		ctx,
		big.NewFloat(64),
		assetID,
		address,
		options...,
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", unstakeOperation.ID())

	if err := client.Wait(ctx, unstakeOperation, coinbase.WithWaitTimeoutSeconds(600)); err != nil {
		log.Fatalf("error waiting for staking operation: %v", err)
	}

	exitMessages, err := unstakeOperation.GetSignedVoluntaryExitMessages()
	if err != nil {
		log.Fatalf("error getting signed voluntary exit messages: %v", err)
	}

	// These exit messages can be broadcasted to the beacon chain whenever you want to voluntarily exit a validator.
	for _, exitMessage := range exitMessages {
		log.Printf(" exit message: %s\n", exitMessage)
	}
}
