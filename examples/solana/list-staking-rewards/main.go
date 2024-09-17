package main

import (
	"context"
	"log"
	"os"
	"time"

	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

var (
	networkID = "solana-mainnet"
)

/*
 * This example code list historical staking rewards of any delegator on the Solana blockchain
 * Run the code with 'go run examples/solana/list-staking-rewards/main.go <api_key_file_path> <wallet_address>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	// Create a new external address on the solana-mainnet-beta network for which you want to view staking rewards.
	address := coinbase.NewExternalAddress(networkID, os.Args[2])

	// Get the rewards earned from staking in the last 10 days.
	// Note that it can take several hours for new rewards to show up.
	rewards, err := client.ListStakingRewards(
		ctx,
		coinbase.Sol,
		[]coinbase.Address{*address},
		time.Now().Add(-10*24*time.Hour),
		time.Now(),
		api.STAKINGREWARDFORMAT_USD,
	)
	if err != nil {
		log.Fatalf("error fetching staking rewards: %v", err)
	}

	// Loop through the rewards and print each staking reward.
	for _, reward := range rewards {
		log.Printf("Staking reward: %s", reward.ToString())
	}
}
