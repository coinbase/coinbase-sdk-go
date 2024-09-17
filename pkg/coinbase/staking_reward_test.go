package coinbase

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ListRewardsRequest struct {
	assetId   string
	address   []Address
	startTime time.Time
	endTime   time.Time
	format    api.StakingRewardFormat
}

type mockController struct {
	stakeAPI  *mocks.StakeAPI
	assetsAPI *mocks.AssetsAPI
}

func TestListStakingRewards_Success(t *testing.T) {
	tests := map[string]struct {
		req             ListRewardsRequest
		expectedRespLen int
		setup           func(*mockController)
	}{
		"happy path with no pages": {
			req: ListRewardsRequest{
				assetId: "test-asset-id",
				address: []Address{
					{
						id: "test-address-id",
					},
				},
				startTime: time.Now(),
				endTime:   time.Now(),
				format:    api.STAKINGREWARDFORMAT_USD,
			},
			setup: func(c *mockController) {
				mockFetchAsset(t, c.assetsAPI, http.StatusOK)
				mockFetchStakingRewards(t, c.stakeAPI, http.StatusOK)
			},
			expectedRespLen: 1,
		},
		"happy path with a page": {
			req: ListRewardsRequest{
				assetId: "test-asset-id",
				address: []Address{
					{
						id: "test-address-id",
					},
				},
				startTime: time.Now(),
				endTime:   time.Now(),
				format:    api.STAKINGREWARDFORMAT_USD,
			},
			setup: func(c *mockController) {
				mockFetchAsset(t, c.assetsAPI, http.StatusOK)
				mockFetchStakingRewardsWithPage(t, c.stakeAPI)
			},
			expectedRespLen: 2,
		},
		"happy path with no addresses": {
			req: ListRewardsRequest{
				assetId:   "test-asset-id",
				startTime: time.Now(),
				endTime:   time.Now(),
				format:    api.STAKINGREWARDFORMAT_USD,
			},
			setup: func(c *mockController) {
			},
			expectedRespLen: 0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mc := &mockController{
				stakeAPI:  mocks.NewStakeAPI(t),
				assetsAPI: mocks.NewAssetsAPI(t),
			}

			tc.setup(mc)

			c := &Client{
				client: &api.APIClient{
					StakeAPI:  mc.stakeAPI,
					AssetsAPI: mc.assetsAPI,
				},
			}

			resp, err := c.ListStakingRewards(
				context.Background(),
				tc.req.assetId,
				tc.req.address,
				tc.req.startTime,
				tc.req.endTime,
				tc.req.format,
			)
			assert.NoError(t, err, "error should be nil")
			assert.Equal(t, tc.expectedRespLen, len(resp), "response length should be 0")
		})
	}
}

func TestListStakingRewards_Failures(t *testing.T) {
	tests := map[string]struct {
		req           ListRewardsRequest
		setup         func(*mockController)
		expectedError string
	}{
		"fail to fetch asset": {
			req: ListRewardsRequest{
				assetId: "test-asset-id",
				address: []Address{
					{
						id: "test-address-id",
					},
				},
				startTime: time.Now(),
				endTime:   time.Now(),
				format:    api.STAKINGREWARDFORMAT_USD,
			},
			setup: func(c *mockController) {
				c.assetsAPI.On("GetAsset", mock.Anything, mock.Anything, mock.Anything).Return(
					api.ApiGetAssetRequest{
						ApiService: c.assetsAPI,
					}).Once()
				c.assetsAPI.On("GetAssetExecute", mock.Anything).Return(
					nil,
					nil,
					fmt.Errorf("failed to fetch asset"),
				).Once()
			},
			expectedError: "failed to fetch asset",
		},
		"fail to fetch staking rewards": {
			req: ListRewardsRequest{
				assetId: "test-asset-id",
				address: []Address{
					{
						id: "test-address-id",
					},
				},
				startTime: time.Now(),
				endTime:   time.Now(),
				format:    api.STAKINGREWARDFORMAT_USD,
			},
			setup: func(c *mockController) {
				mockFetchAsset(t, c.assetsAPI, http.StatusOK)
				c.stakeAPI.On("FetchStakingRewards", mock.Anything, mock.Anything).Return(
					api.ApiFetchStakingRewardsRequest{
						ApiService: c.stakeAPI,
					}).Once()
				c.stakeAPI.On("FetchStakingRewardsExecute", mock.Anything, mock.Anything).Return(
					nil,
					nil,
					fmt.Errorf("failed to fetch staking rewards"),
				).Once()
			},
			expectedError: "failed to fetch staking rewards",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mc := &mockController{
				stakeAPI:  mocks.NewStakeAPI(t),
				assetsAPI: mocks.NewAssetsAPI(t),
			}

			tc.setup(mc)
			c := &Client{
				client: &api.APIClient{
					StakeAPI:  mc.stakeAPI,
					AssetsAPI: mc.assetsAPI,
				},
			}

			_, err := c.ListStakingRewards(
				context.Background(),
				tc.req.assetId,
				tc.req.address,
				tc.req.startTime,
				tc.req.endTime,
				tc.req.format,
			)
			assert.Error(t, err, "error should not be nil")
			assert.Contains(t, err.Error(), tc.expectedError)
		})
	}
}

func TestAddressId(t *testing.T) {
	stakingReward := StakingReward{
		model: api.StakingReward{
			AddressId: "test-address-id",
		},
	}

	assert.Equal(t, "test-address-id", stakingReward.AddressId())
}

func TestAmount(t *testing.T) {
	tests := map[string]struct {
		stakingReward StakingReward
		expected      *big.Float
		expectedError string
	}{
		"happy path usd format": {
			stakingReward: StakingReward{
				model: api.StakingReward{
					Amount: "100",
					Format: api.STAKINGREWARDFORMAT_USD,
				},
				format: api.STAKINGREWARDFORMAT_USD,
				asset: Asset{
					decimals: 18,
				},
			},
			expected: big.NewFloat(1),
		},
		"happy path crypto format": {
			stakingReward: StakingReward{
				model: api.StakingReward{
					Amount: "9837374174",
					Format: api.STAKINGREWARDFORMAT_NATIVE,
				},
				format: api.STAKINGREWARDFORMAT_NATIVE,
				asset: Asset{
					decimals: 18,
				},
			},
			expected: big.NewFloat(0.000000009837374174),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			resp, err := tc.stakingReward.Amount()
			if tc.expectedError != "" {
				assert.Error(t, err, "error should not be nil")
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err, "error should be nil")
				assert.Equal(t, tc.expected.Text('f', -1), resp.Text('f', -1))
			}
		})
	}
}

func TestDate(t *testing.T) {
	stakingReward := StakingReward{
		model: api.StakingReward{
			Date: time.Date(2024, 8, 10, 0, 0, 00, 0, time.UTC),
		},
	}

	resp, err := stakingReward.Date()
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, "2024-08-10 00:00:00 +0000 UTC", resp.String())
}

func mockFetchAsset(t *testing.T, assetsAPI *mocks.AssetsAPI, statusCode int) {
	t.Helper()
	decimals := int32(18)
	assetsAPI.On("GetAsset", mock.Anything, mock.Anything, mock.Anything).Return(
		api.ApiGetAssetRequest{
			ApiService: assetsAPI,
		})
	assetsAPI.On("GetAssetExecute", mock.Anything).Return(
		&api.Asset{
			NetworkId: "test-network-id",
			AssetId:   "test-asset-id",
			Decimals:  &decimals,
		},
		&http.Response{
			StatusCode: statusCode,
		},
		nil,
	)
}

func mockFetchStakingRewards(t *testing.T, stakeAPI *mocks.StakeAPI, statusCode int) {
	t.Helper()
	stakeAPI.On("FetchStakingRewards", mock.Anything, mock.Anything).Return(
		api.ApiFetchStakingRewardsRequest{
			ApiService: stakeAPI,
		})
	stakeAPI.On("FetchStakingRewardsExecute", mock.Anything, mock.Anything).Return(
		&api.FetchStakingRewards200Response{
			Data: []api.StakingReward{
				{
					AddressId: "test-address-id",
					Date:      time.Date(2024, 8, 10, 0, 0, 00, 0, time.UTC),
					Amount:    "1",
					State:     "pending",
					Format:    "usd",
					UsdValue: api.StakingRewardUSDValue{
						Amount:          "1",
						ConversionPrice: "2613.800049",
						ConversionTime:  time.Now(),
					},
				},
			},
			HasMore:  false,
			NextPage: "",
		},
		&http.Response{
			StatusCode: statusCode,
		},
		nil,
	)
}

func mockFetchStakingRewardsWithPage(t *testing.T, stakeAPI *mocks.StakeAPI) {
	t.Helper()
	stakeAPI.On("FetchStakingRewards", mock.Anything, mock.Anything).Return(
		api.ApiFetchStakingRewardsRequest{
			ApiService: stakeAPI,
		}).Once()
	stakeAPI.On("FetchStakingRewardsExecute", mock.Anything, mock.Anything).Return(
		&api.FetchStakingRewards200Response{
			Data: []api.StakingReward{
				{
					AddressId: "test-address-id",
					Date:      time.Date(2024, 8, 10, 0, 0, 00, 0, time.UTC),
					Amount:    "1",
					State:     "pending",
					Format:    "usd",
					UsdValue: api.StakingRewardUSDValue{
						Amount:          "1",
						ConversionPrice: "2613.800049",
						ConversionTime:  time.Now(),
					},
				},
			},
			HasMore:  true,
			NextPage: "test-next-page",
		},
		&http.Response{
			StatusCode: http.StatusOK,
		},
		nil,
	).Once()
	stakeAPI.On("FetchStakingRewards", mock.Anything, mock.Anything).Return(
		api.ApiFetchStakingRewardsRequest{
			ApiService: stakeAPI,
		}).Once()
	stakeAPI.On("FetchStakingRewardsExecute", mock.Anything, mock.Anything).Return(
		&api.FetchStakingRewards200Response{
			Data: []api.StakingReward{
				{
					AddressId: "test-address-id",
					Date:      time.Date(2024, 8, 10, 0, 0, 00, 0, time.UTC),
					Amount:    "1",
					State:     "pending",
					Format:    "usd",
					UsdValue: api.StakingRewardUSDValue{
						Amount:          "1",
						ConversionPrice: "2613.800049",
						ConversionTime:  time.Now(),
					},
				},
			},
			HasMore:  false,
			NextPage: "",
		},
		&http.Response{
			StatusCode: http.StatusOK,
		},
		nil,
	).Once()
}
