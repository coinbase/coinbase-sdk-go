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
	networkID = client.NETWORKIDENTIFIER_ETHEREUM_HOLESKY
	assetID   = coinbase.Eth
)

/*
 * This example code stakes ETH on the Holesky network via Dedicated ETH Staking.
 * Run the code with 'go run examples/ethereum/dedicated-eth-stake/main.go <api_key_file_path> <funding_wallet_address>'
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

	stakeableBalance, err := client.GetStakeableBalance(ctx, assetID, address, coinbase.WithStakingBalanceMode(coinbase.StakingOperationModeNative))
	if err != nil {
		log.Fatalf("error getting stakeable balance: %v", err)
	}

	log.Printf("stakeable balance: %s\n", stakeableBalance)

	listMyValidators(ctx, client, string(networkID), assetID)

	stakeOperation, err := client.BuildStakeOperation(
		ctx,
		big.NewFloat(64),
		assetID,
		address,
		coinbase.WithStakingOperationMode(coinbase.StakingOperationModeNative),
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	stakeOperation, err = client.Wait(ctx, stakeOperation, coinbase.WithWaitTimeoutSeconds(600))
	if err != nil {
		log.Fatalf("error waiting for staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", stakeOperation.ID())

	for index, tx := range stakeOperation.Transactions() {
		log.Printf("Deposit tx %d: %s\n", index+1, tx.UnsignedPayload())
	}

	listMyValidators(ctx, client, string(networkID), assetID)
}

func listMyValidators(ctx context.Context, client *coinbase.Client, networkID string, assetID string) {
	validators, err := client.ListValidators(ctx, networkID, assetID)
	if err != nil {
		log.Fatalf("error listing validators: %v", err)
	}

	for _, validator := range validators {
		log.Printf(validator.ToString())
	}
}
