package coinbase

import (
	"math/big"
	"testing"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/stretchr/testify/assert"
)

func TestBalance_String(t *testing.T) {
	tests := []struct {
		name string
		b    *Balance
		want string
	}{
		{
			name: "success",
			b: &Balance{
				asset: Asset{
					networkId:       "networkId",
					assetId:         "assetId",
					contractAddress: "contractAddress",
					decimals:        18,
				},
				amount: new(big.Float).SetInt64(100),
			},
			want: "Balance { amount: '100' asset: 'Asset { networkId: 'networkId' assetId: 'assetId' contractAddress: 'contractAddress' decimals: '18' }' }",
		},
		{
			name: "from model",
			b: func() *Balance {
				dec := int32(18)
				contractAddress := "contractAddress"
				balance, err := newBalanceFromModel(&client.Balance{
					Asset: client.Asset{
						NetworkId:       "networkId",
						AssetId:         "assetId",
						ContractAddress: &contractAddress,
						Decimals:        &dec,
					},
					Amount: "100000000000000000000",
				})
				if err != nil {
					t.Fatal(err)
				}
				return balance
			}(),
			want: "Balance { amount: '100' asset: 'Asset { networkId: 'networkId' assetId: 'assetId' contractAddress: 'contractAddress' decimals: '18' }' }",
		},
		{
			name: "empty amount",
			b: func() *Balance {
				dec := int32(18)
				contractAddress := "contractAddress"
				balance, err := newBalanceFromModel(&client.Balance{
					Asset: client.Asset{
						NetworkId:       "networkId",
						AssetId:         "assetId",
						ContractAddress: &contractAddress,
						Decimals:        &dec,
					},
					Amount: "",
				})
				if err != nil {
					t.Fatal(err)
				}
				return balance
			}(),
			want: "Balance { amount: '0' asset: 'Asset { networkId: 'networkId' assetId: 'assetId' contractAddress: 'contractAddress' decimals: '18' }' }",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.b.String())
	}
}

func TestBalance_Asset(t *testing.T) {
	asset := Asset{
		networkId:       "networkId",
		assetId:         "assetId",
		contractAddress: "contractAddress",
		decimals:        18,
	}
	tests := []struct {
		name string
		b    *Balance
		want Asset
	}{
		{
			name: "success",
			b: &Balance{
				asset:  asset,
				amount: new(big.Float).SetInt64(100),
			},
			want: asset,
		},
		{
			name: "from model",
			b: func() *Balance {
				dec := int32(18)
				contractAddress := "contractAddress"
				balance, err := newBalanceFromModel(&client.Balance{
					Asset: client.Asset{
						NetworkId:       "networkId",
						AssetId:         "assetId",
						ContractAddress: &contractAddress,
						Decimals:        &dec,
					},
					Amount: "100000000000000000000",
				})
				if err != nil {
					t.Fatal(err)
				}
				return balance
			}(),
			want: asset,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.b.Asset())
	}
}

func TestBalance_Amount(t *testing.T) {
	tests := []struct {
		name string
		b    *Balance
		want *big.Float
	}{
		{
			name: "success",
			b: &Balance{
				asset:  Asset{},
				amount: new(big.Float).SetInt64(100),
			},
			want: new(big.Float).SetInt64(100),
		},
		{
			name: "from model",
			b: func() *Balance {
				dec := int32(18)
				balance, err := newBalanceFromModel(&client.Balance{
					Asset: client.Asset{
						AssetId:  "assetId",
						Decimals: &dec,
					},
					Amount: "100000000000000000000",
				})
				if err != nil {
					t.Fatal(err)
				}
				return balance
			}(),
			want: new(big.Float).SetInt64(100),
		},
	}
	for _, tt := range tests {
		want, _ := tt.want.Float64()
		have, _ := tt.b.Amount().Float64()
		assert.InDelta(t, want, have, 0.0001)
	}
}
