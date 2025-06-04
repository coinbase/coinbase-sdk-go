package coinbase

import (
	"strings"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

const (
	Eth          = "eth"
	Wei          = "wei"
	Gwei         = "gwei"
	GweiDecimals = 9
	Sol          = "sol"

	StakingOperationModePartial = "partial"
	StakingOperationModeDefault = "default"
	StakingOperationModeNative  = "native"

	EthereumHoodi   = string(client.NETWORKIDENTIFIER_ETHEREUM_HOODI)
	EthereumMainnet = string(client.NETWORKIDENTIFIER_ETHEREUM_MAINNET)

	SolanaDevnet  = string(client.NETWORKIDENTIFIER_SOLANA_DEVNET)
	SolanaMainnet = string(client.NETWORKIDENTIFIER_SOLANA_MAINNET)
)

func normalizeNetwork(network string) string {
	return strings.ReplaceAll(network, "_", "-")
}

func toAssetId(assetId string) string {
	return strings.ReplaceAll(assetId, "-", "_")
}
