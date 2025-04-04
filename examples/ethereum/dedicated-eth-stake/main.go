package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	networkID = client.NETWORKIDENTIFIER_ETHEREUM_HOODI
	assetID   = coinbase.Eth
)

/*
 * This example code stakes ETH on the Holesky network via Dedicated ETH Staking.
 * Run the code with 'go run examples/ethereum/dedicated-eth-stake/main.go <api_key_file_path> <funding_wallet_address> <private_key> <rpc_url>'
 */

func main() {
	ctx := context.Background()

	client, err := coinbase.NewClient(
		coinbase.WithAPIKeyFromJSON(os.Args[1]),
	)

	if err != nil {
		log.Fatalf("error creating coinbase client: %v", err)
	}

	address := coinbase.NewExternalAddress(string(networkID), os.Args[2])

	stakeableBalance, err := client.GetStakeableBalance(ctx, assetID, address, coinbase.WithStakingBalanceMode(coinbase.StakingOperationModeNative))
	if err != nil {
		log.Fatalf("error getting stakeable balance: %v", err)
	}

	log.Printf("stakeable balance: %s\n", stakeableBalance.Amount().Text('f', 18))

	unstakeableBalance, err := client.GetUnstakeableBalance(ctx, assetID, address, coinbase.WithStakingBalanceMode(coinbase.StakingOperationModeNative))
	if err != nil {
		log.Fatalf("error getting unstakeableBalance balance: %v", err)
	}

	log.Printf("unstakeableBalance balance: %s\n", unstakeableBalance.Amount().Text('f', 18))

	listMyValidators(ctx, client, string(networkID), assetID)

	options := []coinbase.StakingOperationOption{
		coinbase.WithStakingOperationMode(coinbase.StakingOperationModeNative),
		coinbase.With0x02WithdrawalCredentialType(),
	}

	stakeOperation, err := client.BuildStakeOperation(
		ctx,
		big.NewFloat(64),
		assetID,
		address,
		options...,
	)
	if err != nil {
		log.Fatalf("error building staking operation: %v", err)
	}

	if err := client.Wait(ctx, stakeOperation, coinbase.WithWaitTimeoutSeconds(600)); err != nil {
		log.Fatalf("error waiting for staking operation: %v", err)
	}

	log.Printf("staking operation ID: %s\n", stakeOperation.ID())

	// Load your wallet's private key from which you initiated the above stake operation.
	key, err := crypto.HexToECDSA(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	// Sign the transactions within staking operation resource with your private key.
	err = stakeOperation.Sign(key)
	if err != nil {
		log.Fatal(err)
	}

	// For Hoodi, publicly available RPC URL's can be found here https://chainlist.org/chain/560048
	ethClient, err := ethclient.Dial(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}

	// Broadcast each of the signed transactions to the network.
	for index, transaction := range stakeOperation.Transactions() {
		log.Printf("Deposit tx %d: %s\n", index+1, transaction.UnsignedPayload())

		rawTx, ok := transaction.Raw().(*types.Transaction)
		if !ok {
			log.Fatal("Failed to cast to *types.Transaction")
		}

		if err := ethClient.SendTransaction(context.Background(), rawTx); err != nil {
			log.Fatal(err)
		}

		println(fmt.Sprintf("Broadcasted transaction hash: %s", rawTx.Hash().Hex()))
	}

	listMyValidators(ctx, client, string(networkID), assetID)
}

func listMyValidators(ctx context.Context, client *coinbase.Client, networkID string, assetID string) {
	validators, err := client.ListValidators(ctx, networkID, assetID)
	if err != nil {
		log.Fatalf("error listing validators: %v", err)
	}

	for _, validator := range validators {
		log.Printf("%s %s, %s, %s", validator.WithdrawalCredentials(), validator.WithdrawalAddress(), validator.PublicKey(), validator.Status())
	}
}
