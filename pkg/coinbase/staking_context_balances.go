package coinbase

import (
	"context"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/errors"
)

// StakingContextBalance represents the active stakeable balances for a given address and asset.
type StakingContextBalance struct {
	StakeableBalance        *Balance
	UnstakeableBalance      *Balance
	PendingClaimableBalance *Balance
	ClaimableBalance        *Balance
}

// StakingBalanceOption allows for the passing of custom options to the staking balance
type StakingBalanceOption func(*client.GetStakingContextRequest)

// GetStakeableBalance returns the stakeable balance.
func (c *Client) GetStakeableBalance(ctx context.Context, assetId string, address *Address, o ...StakingBalanceOption) (*Balance, error) {
	sb, err := c.fetchStakingBalances(ctx, assetId, address, o...)
	if err != nil {
		return nil, err
	}

	return sb.StakeableBalance, nil
}

// GetUnstakeableBalance returns the unstakeable balance.
func (c *Client) GetUnstakeableBalance(ctx context.Context, assetId string, address *Address, o ...StakingBalanceOption) (*Balance, error) {
	sb, err := c.fetchStakingBalances(ctx, assetId, address, o...)
	if err != nil {
		return nil, err
	}

	return sb.UnstakeableBalance, nil
}

func (c *Client) GetPendingClaimableBalance(ctx context.Context, assetId string, address *Address, o ...StakingBalanceOption) (*Balance, error) {
	sb, err := c.fetchStakingBalances(ctx, assetId, address, o...)
	if err != nil {
		return nil, err
	}

	return sb.PendingClaimableBalance, nil
}

// GetClaimableBalance returns the claimable balance.
func (c *Client) GetClaimableBalance(ctx context.Context, assetId string, address *Address, o ...StakingBalanceOption) (*Balance, error) {
	sb, err := c.fetchStakingBalances(ctx, assetId, address, o...)
	if err != nil {
		return nil, err
	}

	return sb.ClaimableBalance, nil
}

// WithStakingBalanceOption allows for the passing of custom options to the staking balance
func WithStakingBalanceOption(optionKey string, optionValue string) StakingBalanceOption {
	return func(op *client.GetStakingContextRequest) {
		op.Options[optionKey] = optionValue
	}
}

// WithStakingBalanceMode allows for the setting of the mode of the staking balance
func WithStakingBalanceMode(mode string) StakingBalanceOption {
	return WithStakingBalanceOption("mode", mode)
}

// FetchStakingBalances fetches the staking balances for a given address and asset.
func (c *Client) fetchStakingBalances(ctx context.Context, assetId string, address *Address, o ...StakingBalanceOption) (*StakingContextBalance, error) {
	req := client.GetStakingContextRequest{
		NetworkId: address.NetworkID(),
		AssetId:   assetId,
		AddressId: address.ID(),
		Options: map[string]string{
			"mode": StakingOperationModeDefault,
		},
	}
	for _, f := range o {
		f(&req)
	}

	stakingContext, httpResp, err := c.client.StakeAPI.GetStakingContext(ctx).GetStakingContextRequest(req).Execute()
	if err != nil {
		return nil, errors.MapToUserFacing(err, httpResp)
	}

	return newStakingBalancesFromModel(stakingContext)
}

func newStakingBalancesFromModel(context *client.StakingContext) (*StakingContextBalance, error) {
	stakeableBalance, err := newBalanceFromModel(&context.Context.StakeableBalance)
	if err != nil {
		return nil, err
	}

	unstakeableBalance, err := newBalanceFromModel(&context.Context.UnstakeableBalance)
	if err != nil {
		return nil, err
	}

	pendingClaimableBalance, err := newBalanceFromModel(&context.Context.PendingClaimableBalance)
	if err != nil {
		return nil, err
	}

	claimableBalance, err := newBalanceFromModel(&context.Context.ClaimableBalance)
	if err != nil {
		return nil, err
	}

	return &StakingContextBalance{
		StakeableBalance:        stakeableBalance,
		UnstakeableBalance:      unstakeableBalance,
		PendingClaimableBalance: pendingClaimableBalance,
		ClaimableBalance:        claimableBalance,
	}, nil
}
