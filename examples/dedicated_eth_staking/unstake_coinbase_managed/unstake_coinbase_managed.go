package main

import (
	"context"
	"log"
	"math/big"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

var (
	networkID = "ethereum-holesky"
	addressID = "0x87Bf57c3d7B211a100ee4d00dee08435130A62fA"
	assetID   = coinbase.Eth
)

func main() {
	ctx := context.Background()
	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON("api-key.json"),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewExternalAddress(networkID, addressID)

	unstakeableBalance, err := client.GetUnstakeableBalance(
		ctx,
		assetID,
		address,
		coinbase.WithStakingBalanceMode(coinbase.StakingOperationModeNative),
	)
	if err != nil {
		log.Fatalf("error getting unstakeableBalance balance: %v", err)
	}

	log.Printf("unstakeable balance: %s\n", unstakeableBalance.Amount().String())

	stakingOperation, err := client.BuildUnstakeOperation(
		ctx,
		big.NewFloat(32.00),
		coinbase.Eth,
		address,
		coinbase.WithStakingOperationMode(coinbase.StakingOperationModeNative),
		coinbase.WithStakingOperationOption("immediate", "true"),
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	if err := client.Wait(ctx, stakingOperation); err != nil {
		log.Fatalf("error waiting for staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", stakingOperation.ID())
	log.Printf("staking operation status: %s\n", stakingOperation.Status())
}
