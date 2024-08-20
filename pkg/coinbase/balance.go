package coinbase

import (
	"fmt"
	"math/big"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

// Balance represents a struct that holds a balance data.
type Balance struct {
	model  *client.Balance
	asset  Asset
	amount *big.Float
}

// Asset returns the asset.
func (b *Balance) Asset() Asset {
	return b.asset
}

// Amount returns the amount.
func (b *Balance) Amount() *big.Float {
	return b.amount
}

// String returns the string representation of the balance.
func (b *Balance) String() string {
	return fmt.Sprintf(
		"Balance { amount: '%s' asset: '%s' }",
		b.Amount().String(),
		b.asset.ToString(),
	)
}

func newBalanceFromModel(model *client.Balance) (*Balance, error) {
	asset, err := fromAssetModel(&model.Asset, model.Asset.AssetId)
	if err != nil {
		return nil, err
	}
	amount, ok := new(big.Int).SetString(model.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse amount: %s", model.Amount)
	}

	return &Balance{
		model:  model,
		asset:  asset,
		amount: asset.FromAtomicAmount(amount),
	}, nil
}
