package coinbase

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

type StakingReward struct {
	model  client.StakingReward
	asset  Asset
	format client.StakingRewardFormat
}

func NewStakingReward(model client.StakingReward, asset Asset, format client.StakingRewardFormat) StakingReward {
	return StakingReward{
		model:  model,
		asset:  asset,
		format: format,
	}
}

func (s StakingReward) Amount() (*big.Int, error) {
	amountBigInt, ok := new(big.Int).SetString(s.model.GetAmount(), 10)
	if !ok {
		return nil, fmt.Errorf("invalid amount found: %s", s.model.GetAmount())
	}

	if s.format == client.STAKINGREWARDFORMAT_USD {
		return amountBigInt.Div(amountBigInt, big.NewInt(100)), nil
	}
	return s.asset.fromAtomicAmount(amountBigInt), nil
}

func (s StakingReward) Date() (time.Time, error) {

	parsedDate, err := time.Parse(time.DateOnly, s.model.GetDate())
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date found: %s", s.model.GetDate())
	}

	return parsedDate, nil
}

func (s StakingReward) AddressId() string {
	return s.model.GetAddressId()
}

func (s StakingReward) ToString() string {
	return fmt.Sprintf(
		"StakingReward { date: '%s' address: '%s' amount: '%s' }",
		s.model.GetDate(),
		s.model.GetAddressId(),
		s.model.GetAmount(),
	)
}

func (c *Client) ListStakingRewards(
	ctx context.Context,
	addresses []Address,
	startTime time.Time,
	endTime time.Time,
	format client.StakingRewardFormat,
) ([]StakingReward, error) {
	var (
		addressIds     = make([]string, 0, len(addresses))
		stakingRewards = make([]StakingReward, 0)
		page           = "first"
	)

	// No addresses to fetch rewards for.
	if len(addresses) == 0 {
		return []StakingReward{}, nil
	}

	// We're assuming the addresses will be all for the same networkId and assetId.
	networkId, assetId := addresses[0].NetworkID(), addresses[0].ID()

	// Get the asset.
	asset, err := c.fetchAsset(ctx, networkId, assetId)
	if err != nil {
		return []StakingReward{}, err
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

	for page != "" {

		if page == "first" {
			page = ""
		}

		// Get the rewards for the address ids.
		resp, err := c.getRewards(ctx, page, rewardsReq)
		if err != nil {
			return []StakingReward{}, err
		}

		for _, stakingReward := range resp.GetData() {
			stakingRewards = append(stakingRewards, NewStakingReward(stakingReward, asset, format))
		}

		if resp.GetHasMore() {
			if resp.GetNextPage() != "" {
				page = resp.GetNextPage()
			}
		}

	}

	return stakingRewards, nil
}

// getRewards calls the backend to retrieve rewards for a list of address ids. This is a helper function
// to help determine whether to use page in the call or not.
func (c *Client) getRewards(
	ctx context.Context,
	page string,
	req client.FetchStakingRewardsRequest,
) (*client.FetchStakingRewards200Response, error) {
	var (
		resp    *client.FetchStakingRewards200Response
		httpRes *http.Response
		err     error
	)

	if page == "" {
		resp, httpRes, err = c.client.StakeAPI.
			FetchStakingRewards(ctx).
			FetchStakingRewardsRequest(req).
			Limit(100).
			Execute()
	} else {
		resp, httpRes, err = c.client.StakeAPI.
			FetchStakingRewards(ctx).
			FetchStakingRewardsRequest(req).
			Limit(100).
			Page(page).
			Execute()
	}

	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode != 200 {
		return nil, err
	}

	return resp, nil
}
