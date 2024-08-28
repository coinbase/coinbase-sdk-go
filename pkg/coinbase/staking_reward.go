package coinbase

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/errors"
)

// StakingReward represents a struct that holds a staking reward data.
type StakingReward struct {
	model  client.StakingReward
	asset  Asset
	format client.StakingRewardFormat
}

// NewStakingReward creates a new staking rewards object.
func NewStakingReward(model client.StakingReward, asset Asset, format client.StakingRewardFormat) StakingReward {
	return StakingReward{
		model:  model,
		asset:  asset,
		format: format,
	}
}

// ListStakingRewards list the staking rewards for a given set of addresses.
func (c *Client) ListStakingRewards(
	ctx context.Context,
	assetId string,
	addresses []Address,
	startTime time.Time,
	endTime time.Time,
	format client.StakingRewardFormat,
) ([]StakingReward, error) {
	var (
		addressIds     = make([]string, 0, len(addresses))
		stakingRewards = make([]StakingReward, 0)
		page           = ""
	)

	// No addresses to fetch rewards for.
	if len(addresses) == 0 {
		return nil, nil
	}

	// We're assuming the addresses will be all for the same networkId and assetId.
	networkId := addresses[0].NetworkID()

	// Get the asset.
	asset, err := c.fetchAsset(ctx, networkId, assetId)
	if err != nil {
		return nil, err
	}

	// Create a list of address ids.
	for _, address := range addresses {
		addressIds = append(addressIds, address.ID())
	}

	rewardsReq := client.FetchStakingRewardsRequest{
		NetworkId:  normalizeNetwork(networkId),
		AssetId:    assetId,
		AddressIds: addressIds,
		StartTime:  startTime,
		EndTime:    endTime,
		Format:     format,
	}

	for {
		// Get the rewards for the address ids.
		resp, err := c.getRewards(ctx, page, rewardsReq)
		if err != nil {
			return nil, err
		}

		for _, stakingReward := range resp.GetData() {
			stakingRewards = append(stakingRewards, NewStakingReward(stakingReward, asset, format))
		}

		if resp.GetHasMore() && resp.GetNextPage() != "" {
			page = resp.GetNextPage()
		} else {
			break
		}
	}

	return stakingRewards, nil
}

// AddressId returns the address of the staking reward.
func (s StakingReward) AddressId() string {
	return s.model.GetAddressId()
}

// Amount returns the amount of the staking reward.
func (s StakingReward) Amount() (*big.Float, error) {
	if s.format == client.STAKINGREWARDFORMAT_USD {
		amountBigInt, ok := new(big.Float).SetString(s.model.GetAmount())
		if !ok {
			return nil, fmt.Errorf("invalid amount found: %s", s.model.GetAmount())
		}
		return new(big.Float).Quo(amountBigInt, big.NewFloat(100)), nil
	}

	amountBigInt, ok := new(big.Int).SetString(s.model.GetAmount(), 10)
	if !ok {
		return nil, fmt.Errorf("invalid amount found: %s", s.model.GetAmount())
	}
	return s.asset.FromAtomicAmount(amountBigInt), nil
}

// Date returns the date of the staking reward.
func (s StakingReward) Date() (time.Time, error) {
	parsedDate, err := time.Parse(time.DateOnly, s.model.GetDate())
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date found: %s", s.model.GetDate())
	}

	return parsedDate, nil
}

// USDValue returns the USD value of the staking reward.
func (s StakingReward) USDValue() (*big.Float, error) {
	amountBigInt, ok := new(big.Float).SetString(s.model.GetAmount())
	if !ok {
		return nil, fmt.Errorf("invalid amount found: %s", s.model.GetAmount())
	}
	return new(big.Float).Quo(amountBigInt, big.NewFloat(100)), nil
}

// ConversionPrice returns the conversion price of the staking reward.
func (s StakingReward) ConversionPrice() (*big.Float, error) {
	amountBigInt, ok := new(big.Float).SetString(s.model.UsdValue.ConversionPrice)
	if !ok {
		return nil, fmt.Errorf("invalid conversion price found: %s", s.model.UsdValue.ConversionPrice)
	}
	return amountBigInt, nil
}

// ConversionTime returns the conversion time of the staking reward.
func (s StakingReward) ConversionTime() time.Time {
	return s.model.UsdValue.ConversionTime
}

// ToString prints a simplified version of the reward.
func (s StakingReward) ToString() string {
	usdValue := s.model.GetUsdValue()

	return fmt.Sprintf(
		"StakingReward { date: '%s' address: '%s' amount: '%s' usd_value: '%s' conversion_price: '%s' conversion_time: '%s'}",
		s.model.GetDate(),
		s.model.GetAddressId(),
		s.model.GetAmount(),
		usdValue.GetAmount(),
		usdValue.GetConversionPrice(),
		usdValue.GetConversionTime().String(),
	)
}

// ToJSON prints the string representation of the reward.
func (s StakingReward) ToJSON() string {
	jsonData, err := json.MarshalIndent(s.model, "", "  ")
	if err != nil {
		fmt.Println("Error converting struct to JSON:", err)
		return ""
	}
	return string(jsonData)
}

// getRewards calls the backend to retrieve rewards for a list of address ids. This is a helper function
// to help determine whether to use page in the call or not.
func (c *Client) getRewards(
	ctx context.Context,
	page string,
	req client.FetchStakingRewardsRequest,
) (*client.FetchStakingRewards200Response, error) {
	var (
		resp     *client.FetchStakingRewards200Response
		httpResp *http.Response
		err      error
	)

	if page == "" {
		resp, httpResp, err = c.client.StakeAPI.
			FetchStakingRewards(ctx).
			FetchStakingRewardsRequest(req).
			Limit(100).
			Execute()
	} else {
		resp, httpResp, err = c.client.StakeAPI.
			FetchStakingRewards(ctx).
			FetchStakingRewardsRequest(req).
			Limit(100).
			Page(page).
			Execute()
	}

	if err != nil {
		return nil, errors.MapToUserFacing(err, httpResp)
	}

	return resp, nil
}
