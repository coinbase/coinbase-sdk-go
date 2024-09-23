package main

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
 * This example code stakes ETH on the Holesky network via Shared ETH Staking.
 * Run the code with 'go run examples/ethereum/stake/main.go <api_key_file_path> <wallet_address> <wallet_private_key>'
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

	stakeOperation, err := client.BuildStakeOperation(
		ctx,
		big.NewFloat(0.0001),
		coinbase.Eth,
		address,
		coinbase.WithStakingOperationMode(coinbase.StakingOperationModePartial),
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", stakeOperation.ID())

	key, err := crypto.HexToECDSA(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	// Sign the transactions within staking operation resource with your private key.
	err = stakeOperation.Sign(key)
	if err != nil {
		log.Fatal(err)
	}

	// For Holesky, publicly available RPC URL's can be found here https://chainlist.org/chain/17000
	ethClient, err := ethclient.Dial("<RPC_NODE_URL>")
	if err != nil {
		log.Fatal(err)
	}

	// Broadcast each of the signed transactions to the network.
	for _, transaction := range stakeOperation.Transactions() {
		rawTx := transaction.Raw()
		ethTx, ok := rawTx.(*types.Transaction)
		if !ok {
			log.Fatal("failed to cast transaction to Ethereum transaction")
		}

		if err := ethClient.SendTransaction(context.Background(), ethTx); err != nil {
			log.Fatal(err)
		}

		log.Printf("Broadcasted transaction hash: %s", ethTx.Hash().Hex())
	}
}
