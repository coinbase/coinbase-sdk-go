package coinbase

import "strings"

const (
	Eth          = "eth"
	Wei          = "wei"
	Gwei         = "gwei"
	GweiDecimals = 9

	StakingOperationModePartial = "partial"
	StakingOperationModeDefault = "default"
	StakingOperationModeNative  = "native"
)

func normalizeNetwork(network string) string {
	return strings.ReplaceAll(network, "_", "-")
}

func toAssetId(assetId string) string {
	return strings.ReplaceAll(assetId, "-", "_")
}
