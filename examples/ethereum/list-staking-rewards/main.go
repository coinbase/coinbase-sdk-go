package main

import (
	"context"
	"log"
	"os"
	"time"

	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

/*
 * This example code list historical staking rewards for any ETH validator or Shared ETH Staking wallet address.
 * For example: the addresses:
 *   * Shared ETH Staking wallet address: 0xddb00798137e9e7cc89f1e9679e6ce6ea580b8f9
 *   * ETH Validator: 0xa1d1ad0714035353258038e964ae9675dc0252ee22cea896825c01458e1807bfad2f9969338798548d9858a571f7425c
 * Run the code with 'go run examples/ethereum/list-staking-rewards/main.go <api_key_file_path> <wallet_address>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewExternalAddress(
		"ethereum-mainnet",
		os.Args[1],
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
