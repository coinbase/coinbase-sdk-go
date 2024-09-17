package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

var (
	networkID = "solana-mainnet"
)

/*
 * This example code lists historical staking balances of any delegator on the Solana blockchain
 * Run the code with 'go run examples/solana/list-staking-balances/main.go <api_key_file_path> <wallet_address>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	// Create a new external address on the solana-mainnet-beta network for which you want to view staking balances.
	address := coinbase.NewExternalAddress(networkID, os.Args[2])

	// Get the balances earned from staking in the last 10 days.
	// Note that it can take several hours for new balances to show up.
	balances, err := client.ListHistoricalStakingBalances(
		ctx,
		coinbase.Sol,
		address,
		time.Now().Add(-10*24*time.Hour),
		time.Now(),
	)
	if err != nil {
		log.Fatalf("error fetching staking balances: %v", err)
	}

	// Loop through the balances and print each staking balance.
	for _, balance := range balances {
		log.Printf("Staking balance: %s", balance.String())
	}
}
