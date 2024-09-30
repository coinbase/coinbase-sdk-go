package coinbase

import "strings"

const (
	Eth          = "eth"
	Wei          = "wei"
	Gwei         = "gwei"
	GweiDecimals = 9
	Sol          = "sol"

	StakingOperationModePartial = "partial"
	StakingOperationModeDefault = "default"
	StakingOperationModeNative  = "native"

	EthereumHolesky = "ethereum-holesky"
	EthereumMainnet = "ethereum-mainnet"

	SolanaDevnet  = "solana-devnet"
	SolanaMainnet = "solana-mainnet"
)

func normalizeNetwork(network string) string {
	return strings.ReplaceAll(network, "_", "-")
}

func toAssetId(assetId string) string {
	return strings.ReplaceAll(assetId, "-", "_")
}
