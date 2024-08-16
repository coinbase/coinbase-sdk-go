package main

import (
	"context"
	"log"
	"math/big"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

func main() {
	ctx := context.Background()
	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON("api_key.json"),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewAddress("ethereum-holesky", "0x57a063e1df096aaA6b2068C3C7FE6Ac4BC3c4F58")
	op, err := client.BuildStakeOperation(ctx, address, "eth", big.NewFloat(0.0001))
	log.Printf("staking operation ID: %s\n", op.ID())
	log.Printf("staking operation Transactions: %+v\n", op.Transactions())

}
