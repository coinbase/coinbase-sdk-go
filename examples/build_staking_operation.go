package main

import (
	"context"
	"log"
	"math/big"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
)

func main() {
	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON("/Users/deangalvin/Downloads/cdp_api_key (3).json"),
	)
	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}
	address := client.NewAddress("ethereum-holesky", "0x57a063e1df096aaA6b2068C3C7FE6Ac4BC3c4F58")
	op, err := address.BuildStakeOperaiton(context.Background(), "eth", big.NewFloat(0.0001))
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", op.ID())
	log.Printf("staking operation Transactions: %+v\n", op.Transactions())
	//
	address := coinbase.NewAddress("ethereum-holesky", "0x57a063e1df096aaA6b2068C3C7FE6Ac4BC3c4F58")
	//req := address
//	.BuildStakeOperaitonRequest(context.Background(), "eth", big.NewFloat(0.0001))
	op, err := client.BuildStakeOperation(
		context.Background(),
		&coinbase.BuildStakeOperationRequest{
			Address: address,
			AssetId: "eth",
			Amount: big.NewFloat(0.0001),
			Options: map[string]string{
				"mode": "default",
			}
		}
	)
	rewardsIter, err := client.ListStakingRewards(
		ctx, []coinbase.Address{address}, "eth", time.Now().Add(-time.Hour*24*7), time.Now(),
		coinbase.WithStakingRewardsLimit(100),
		coinbase.WithStakingRewardsPageOffset(2),
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", op.ID())
	log.Printf("staking operation Transactions: %+v\n", op.Transactions())

}
