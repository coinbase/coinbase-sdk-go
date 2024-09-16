package coinbase

import (
	"context"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/mocks"
)

func TestListHistoricalStakingBalances(t *testing.T) {
	parsedTime := time.Now().Truncate(time.Hour * 24).UTC()
	tests := []struct {
		name string
		want *StakingBalance
	}{
		{
			"Test ListHistoricalStakingBalances",
			&StakingBalance{
				model: &client.StakingBalance{
					ParticipantType: "participantType",
				},
				bondedStake: &Balance{
					amount: big.NewFloat(123),
				},
				unbondedBalance: &Balance{
					amount: big.NewFloat(1234),
				},
				parsedDate: parsedTime,
			},
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		assetId := "assetId"
		address := &Address{}
		startTime := time.Now()
		endTime := time.Now()
		decimals := int32(0)
		t.Run(tt.name, func(t *testing.T) {
			stakeAPI := mocks.NewStakeAPI(t)
			c := &Client{
				client: &client.APIClient{
					StakeAPI: stakeAPI,
				},
			}
			stakeAPI.On("FetchHistoricalStakingBalances", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
				api.ApiFetchHistoricalStakingBalancesRequest{ApiService: stakeAPI},
			).Once()
			stakeAPI.On("FetchHistoricalStakingBalancesExecute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
				&client.FetchHistoricalStakingBalances200Response{
					Data: []client.StakingBalance{
						{
							ParticipantType: "participantType",
							Date:            parsedTime.Format(timestampFormat),
							BondedStake: client.Balance{
								Amount: "123",
								Asset: client.Asset{
									Decimals: &decimals,
								},
							},
							UnbondedBalance: client.Balance{
								Amount: "1234",
								Asset: client.Asset{
									Decimals: &decimals,
								},
							},
						},
					},
				},
				&http.Response{StatusCode: 200},
				nil,
			).Once()
			resp, err := c.ListHistoricalStakingBalances(ctx, assetId, address, startTime, endTime)
			require.NoError(t, err)
			require.Equal(t, 1, len(resp))
			assert.Equal(t, tt.want.ParticipantType(), resp[0].ParticipantType())
			assert.Equal(t, tt.want.Date(), resp[0].Date())
			wantBondedStake, _ := tt.want.BondedStake().Amount().Float64()
			haveBondedStake, _ := resp[0].BondedStake().Amount().Float64()
			assert.InDelta(t, wantBondedStake, haveBondedStake, 0.0001)
			wantUnbondedBalance, _ := tt.want.UnbondedBalance().Amount().Float64()
			haveUnbondedBalance, _ := resp[0].UnbondedBalance().Amount().Float64()
			assert.InDelta(t, wantUnbondedBalance, haveUnbondedBalance, 0.0001)
		})
	}
}
