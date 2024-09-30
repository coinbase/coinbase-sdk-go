package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

/*
 * This example code list historical staking balances for any ETH validator or Shared ETH Staking wallet address.
 * For example: the addresses:
 *   * Shared ETH Staking wallet address: 0xddb00798137e9e7cc89f1e9679e6ce6ea580b8f9
 *   * ETH Validator: 0xadbf3776d60b3f9dd30cb3257b50583898745deb40cb6cb842120753bf055f6c3863e0f5bdb5c403d9aa5a275ce165e8
 * Run the code with 'go run examples/ethereum/list-staking-balances/main.go <api_key_file_path> <wallet_address>'
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
		coinbase.EthereumMainnet,
		os.Args[1],
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
		println(balance)
	}
}
