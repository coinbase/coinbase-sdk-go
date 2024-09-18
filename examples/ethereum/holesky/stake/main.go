package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"time"

	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

/*
 * This example code stakes ETH on the Holesky network.
 * Run the code with 'go run examples/ethereum/stake/main.go <api_key_file_path> <wallet_address>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewExternalAddress("ethereum-holesky", os.Args[2])

	stakeableBalance, err := client.GetStakeableBalance(ctx, coinbase.Eth, address, coinbase.WithStakingBalanceMode(coinbase.StakingOperationModePartial))
	if err != nil {
		log.Fatalf("error getting stakeable balance: %v", err)
	}

	log.Printf("stakeable balance: %s\n", stakeableBalance)

	op, err := client.BuildStakeOperation(
		ctx,
		big.NewFloat(0.0001),
		coinbase.Eth,
		address,
		coinbase.WithStakingOperationMode(coinbase.StakingOperationModePartial),
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", op.ID())

	for _, transaction := range op.Transactions() {
		log.Printf("staking operation Transaction: %+v\n", transaction)
	}

	address = coinbase.NewExternalAddress(
		"ethereum-mainnet",
		"0xddb00798137e9e7cc89f1e9679e6ce6ea580b8f9",
	)

	rewards, err := client.ListStakingRewards(
		ctx,
		coinbase.Eth,
		[]coinbase.Address{*address},
		time.Now().Add(-7*24*time.Hour),
		time.Now(),
		api.STAKINGREWARDFORMAT_USD,
	)
	if err != nil {
		log.Fatalf("error listing rewards: %v", err)
	}

	for _, reward := range rewards {
		println(reward.ToJSON())
	}

	address = coinbase.NewExternalAddress(
		"ethereum-mainnet",
		"0xadbf3776d60b3f9dd30cb3257b50583898745deb40cb6cb842120753bf055f6c3863e0f5bdb5c403d9aa5a275ce165e8",
	)
	balances, err := client.ListHistoricalStakingBalances(
		ctx,
		coinbase.Eth,
		address,
		time.Now().Add(-7*24*time.Hour),
		time.Now(),
	)
	if err != nil {
		log.Fatalf("error listing balances: %v", err)
	}

	for _, balance := range balances {
		println(balance.String())
	}
}
