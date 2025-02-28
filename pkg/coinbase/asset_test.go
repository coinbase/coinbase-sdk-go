package coinbase

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromAtomicAmount(t *testing.T) {
	asset := Asset{
		networkId:       "testnet",
		assetId:         "test-asset",
		contractAddress: "0x123",
		decimals:        18,
	}

	tests := []struct {
		name                               string
		amountString                       string
		expectedAmountStringUpto18Decimals string
	}{
		{
			name:                               "Test with 0 eth",
			amountString:                       "0",
			expectedAmountStringUpto18Decimals: "0.000000000000000000",
		},
		{
			name:                               "Test with negative value",
			amountString:                       "-1000000000000000000",
			expectedAmountStringUpto18Decimals: "-1.000000000000000000",
		},
		{
			name:                               "Test with 1 eth",
			amountString:                       "1000000000000000000",
			expectedAmountStringUpto18Decimals: "1.000000000000000000",
		},
		{
			name:                               "Test with very small fractional value",
			amountString:                       "1",
			expectedAmountStringUpto18Decimals: "0.000000000000000001",
		},
		{
			name:                               "Test with maximum uint128 value",
			amountString:                       "340282366920938463463374607431768211456",
			expectedAmountStringUpto18Decimals: "340282366920938463463.374607431768211456",
		},
		{
			name:                               "Test with large fractional value",
			amountString:                       "123456789012345678901234567890",
			expectedAmountStringUpto18Decimals: "123456789012.345678901234567890",
		},
		{
			name:                               "Test with random fractional eth",
			amountString:                       "32006467556000000000",
			expectedAmountStringUpto18Decimals: "32.006467556000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := asset.FromAtomicAmount(getTestBigInt(t, tt.amountString))
			assert.Equal(t, tt.expectedAmountStringUpto18Decimals, result.Text('f', 18))
		})
	}
}

func TestToAtomicAmount(t *testing.T) {
	asset := Asset{
		networkId:       "testnet",
		assetId:         "test-asset",
		contractAddress: "0x123",
		decimals:        18,
	}

	tests := []struct {
		name                               string
		amountString                       string
		expectedAmountStringUpto18Decimals string
	}{
		{
			name:                               "Test with 0 eth",
			amountString:                       "0.000000000000000000",
			expectedAmountStringUpto18Decimals: "0",
		},
		{
			name:                               "Test with negative value",
			amountString:                       "-1.000000000000000000",
			expectedAmountStringUpto18Decimals: "-1000000000000000000",
		},
		{
			name:                               "Test with 1 eth",
			amountString:                       "1.000000000000000000",
			expectedAmountStringUpto18Decimals: "1000000000000000000",
		},
		{
			name:                               "Test with very small fractional value",
			amountString:                       "0.000000000000000001",
			expectedAmountStringUpto18Decimals: "1",
		},
		{
			name:                               "Test with maximum uint128 value",
			amountString:                       "340282366920938463463.374607431768211456",
			expectedAmountStringUpto18Decimals: "340282366920938463463374607431768211456",
		},
		{
			name:                               "Test with large fractional value",
			amountString:                       "123456789012.345678901234567890",
			expectedAmountStringUpto18Decimals: "123456789012345678901234567890",
		},
		{
			name:                               "Test with random fractional eth",
			amountString:                       "32.006467556000000000",
			expectedAmountStringUpto18Decimals: "32006467556000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := asset.ToAtomicAmount(getTestBigFloat(t, tt.amountString))
			assert.Equal(t, tt.expectedAmountStringUpto18Decimals, result.String())
		})
	}
}

func getTestBigInt(t *testing.T, input string) *big.Int {
	t.Helper()

	result, ok := new(big.Int).SetString(input, 10)
	assert.True(t, ok)

	return result
}

func getTestBigFloat(t *testing.T, input string) *big.Float {
	t.Helper()

	result, _, err := big.ParseFloat(input, 10, 128, big.ToNearestEven)
	assert.NoError(t, err)

	return result
}
