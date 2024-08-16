package coinbase

import (
	"context"
	"fmt"
	"math/big"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

type Asset struct {
	networkId       string
	assetId         string
	contractAddress string
	decimals        int32
}

func NewAsset(
	networkId string,
	assetId string,
	contractAddress string,
	decimals int32,
) Asset {
	return Asset{
		networkId:       networkId,
		assetId:         assetId,
		contractAddress: contractAddress,
		decimals:        decimals,
	}
}

func (a Asset) fromAtomicAmount(wholeAmount *big.Int) *big.Int {
	// Calculate 10^decimals.
	pow := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(a.decimals)), nil)

	// Multiply wholeAmount by 10^decimals.
	atomicAmount := wholeAmount.Div(wholeAmount, pow)

	return atomicAmount
}

func (a Asset) AssetId() string {
	return a.assetId
}

func (a Asset) ToString() string {
	return fmt.Sprintf(
		"Asset { networkId: '%s' assetId: '%s' contractAddress: '%s' decimals: '%s' }",
		a.networkId,
		a.assetId,
		a.contractAddress,
		string(a.decimals),
	)
}

func primaryDenomination(assetId string) string {
	if assetId == Gwei || assetId == Wei {
		return Eth
	}
	return assetId
}

func fromAssetModel(model *client.Asset, assetId string) (Asset, error) {

	decimals := model.GetDecimals()

	if toAssetId(model.GetAssetId()) != toAssetId(assetId) {
		switch assetId {
		case "gwei":
			decimals = GweiDecimals
		case "wei":
			decimals = 0
		case "eth":
		default:
			return Asset{}, fmt.Errorf("invalid asset ID: %s", assetId)
		}
	}

	return Asset{
		networkId:       model.GetNetworkId(),
		assetId:         model.GetAssetId(),
		contractAddress: model.GetContractAddress(),
		decimals:        decimals,
	}, nil
}

func (c *Client) fetchAsset(ctx context.Context, networkId string, assetId string) (Asset, error) {
	// Get the Asset from the backend.
	asset, httpRes, err := c.client.AssetsAPI.GetAsset(
		ctx,
		normalizeNetwork(networkId),
		primaryDenomination(assetId),
	).Execute()
	if err != nil {
		return Asset{}, err
	}

	if httpRes.StatusCode != 200 {
		return Asset{}, err
	}

	return fromAssetModel(asset, assetId)
}
