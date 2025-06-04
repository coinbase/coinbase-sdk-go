package main

import (
	"context"
	"log"
	"os"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

/*
 * This example code stakes ETH on the Hoodi network via Dedicated ETH Staking.
 * Run the code with 'go run examples/ethereum/dedicated-eth-list-validators/main.go <api_key_file_path>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)

	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	listMyValidators(ctx, client, coinbase.EthereumHoodi, coinbase.Eth, coinbase.ValidatorStatusUnknown)
}

func listMyValidators(ctx context.Context, client *coinbase.Client, networkID string, assetID string, status coinbase.ValidatorStatus) {
	var opts []coinbase.ListValidatorsOption

	if status != coinbase.ValidatorStatusUnknown {
		opts = append(opts, coinbase.WithListValidatorsStatusOption(status))
	}

	validators, err := client.ListValidators(ctx, networkID, assetID, opts...)
	if err != nil {
		log.Fatalf("error listing validators: %v", err)
	}

	for _, validator := range validators {
		log.Printf("%s %s, %s, %s", validator.WithdrawalCredentials(), validator.WithdrawalAddress(), validator.PublicKey(), validator.Status())
	}
}
