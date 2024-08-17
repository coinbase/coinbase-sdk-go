package main

import (
	"context"
	"log"
	"math/big"
	"time"

	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

func main() {
	ctx := context.Background()
	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON("api-key.json"),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewAddress("ethereum-holesky", "0x57a063e1df096aaA6b2068C3C7FE6Ac4BC3c4F58")
	op, err := client.BuildStakeOperation(ctx, address, "eth", big.NewFloat(0.0001))
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}
	log.Printf("staking operation ID: %s\n", op.ID())
	for _, transaction := range op.Transactions() {
		log.Printf("staking operation Transaction: %+v\n", transaction)
	}

	address = coinbase.NewAddress(
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

}
