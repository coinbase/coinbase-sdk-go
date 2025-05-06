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
 * This example code does an execution layer orchestrated unstaking of an ethereum validator on the Hoodi network.
 * Run the code with 'go run examples/ethereum/dedicated-eth-execution-unstake/main.go <api_key_file_path> <withdrawal_address> <private_key> <rpc_url>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)

	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	withdrawalAddress := coinbase.NewExternalAddress(coinbase.EthereumHoodi, os.Args[2])

	unstakeableBalance, err := client.GetUnstakeableBalance(ctx, coinbase.Eth, withdrawalAddress, coinbase.WithNativeStakingBalanceMode())
	if err != nil {
		log.Fatalf("error getting unstakeableBalance balance: %v", err)
	}

	log.Printf("unstakeable balance: %s\n", unstakeableBalance.Amount().Text('f', 18))

	builder, err := coinbase.NewExecutionLayerWithdrawalsOptionBuilder(ctx, client, withdrawalAddress)
	if err != nil {
		log.Fatalf("error creating execution layer withdrawals option builder: %v", err)
	}

	// Add a validator withdrawal to the builder. Use amount 0 if it needs to be fully unstaked/exited.
	if err := builder.AddValidatorWithdrawal("0xb8b2de0a6179248c19c75ecbd3b37359e32e388dfeb20e4f54106ebb736167b55582bef0f740c63dc69414f8dec9885f", big.NewFloat(1.0)); err != nil {
		log.Fatalf("error adding validator withdrawal: %v", err)
	}

	options := []coinbase.StakingOperationOption{
		coinbase.WithNativeStakingOperationMode(),
		coinbase.WithExecutionLayerWithdrawals(builder),
	}

	unstakeOperation, err := client.BuildUnstakeOperation(
		ctx,
		big.NewFloat(0), // Amount here doesn't matter.
		coinbase.Eth,
		withdrawalAddress,
		options...,
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", unstakeOperation.ID())

	if err := client.Wait(ctx, unstakeOperation, coinbase.WithWaitTimeoutSeconds(600)); err != nil {
		log.Fatalf("error waiting for staking operation: %v", err)
	}
	// Load your wallet's private key from which you initiated the above stake operation.
	key, err := crypto.HexToECDSA(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Signing unstake operation ...")

	// Sign the transactions within staking operation resource with your private key.
	if err := unstakeOperation.Sign(key); err != nil {
		log.Fatal(err)
	}

	// For Hoodi, publicly available RPC URL's can be found here https://chainlist.org/chain/560048
	ethClient, err := ethclient.Dial(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}

	// Broadcast each of the signed transactions to the network.
	for index, transaction := range unstakeOperation.Transactions() {
		log.Printf("Partial Withdrawal tx %d: %s\n", index+1, transaction.UnsignedPayload())

		rawTx, ok := transaction.Raw().(*types.Transaction)
		if !ok {
			log.Fatal("Failed to cast to *types.Transaction")
		}

		if err := ethClient.SendTransaction(context.Background(), rawTx); err != nil {
			log.Fatal(err)
		}

		log.Printf("Broadcasted transaction hash: %s", rawTx.Hash().Hex())
	}

}
