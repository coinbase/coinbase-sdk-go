package main

import (
	"context"
	"fmt"
	"log"

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

	// List all validators.
	log.Println("Listing all validators")
	validators, err := client.ListValidators(ctx, "ethereum-holesky", coinbase.Eth)
	if err != nil {
		log.Fatalf("error listing rewards: %v", err)
	}

	for _, validator := range validators {
		fmt.Println(validator.ToString())
	}

	// List all active validators.
	log.Println("Listing all active validators")
	validators, err = client.ListValidators(
		ctx,
		"ethereum-holesky",
		coinbase.Eth,
		coinbase.WithListValidatorsStatusOption(coinbase.ValidatorStatusActive),
	)
	if err != nil {
		log.Fatalf("error listing rewards: %v", err)
	}

	for _, validator := range validators {
		fmt.Println(validator.ToString())
	}

	// Get a specific validator.
	log.Println("Getting a specific validator")
	validator, err := client.GetValidator(ctx, "ethereum-holesky", coinbase.Eth, "0x940303549cf1ccf6e9aafbc7996974b4f48ff112a3ca41ba187d9725dfc4f4c715105ec711221448dc50c777e0108066")
	if err != nil {
		log.Fatalf("error getting validator: %v", err)
	}

	fmt.Printf("Validator id [%s] has state: %s", validator.ID(), validator.Status())
}
