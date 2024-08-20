package coinbase

import (
	"context"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

// StakingBalances represents the active stakeable balances for a given address and asset.
type StakingBalances struct {
	StakeableBalance   *Balance
	UnstakeableBalance *Balance
	ClaimableBalance   *Balance
}

// StakingBalanceOption allows for the passing of custom options to the staking balance
type StakingBalanceOption func(*client.GetStakingContextRequest)

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
func (c *Client) FetchStakingBalances(ctx context.Context, address *Address, assetId string, o ...StakingBalanceOption) (*StakingBalances, error) {
	req := client.GetStakingContextRequest{
		NetworkId: address.NetworkID(),
		AssetId:   assetId,
		AddressId: address.ID(),
		Options: map[string]string{
			"mode": "default",
		},
	}
	for _, f := range o {
		f(&req)
	}
	context, _, err := c.client.StakeAPI.GetStakingContext(ctx).GetStakingContextRequest(req).Execute()
	if err != nil {
		return nil, err
	}

	return newStakingBalancesFromModel(context)
}

// GetStakeableBalance returns the stakeable balance.
func (b *StakingBalances) GetStakeableBalance() *Balance {
	return b.StakeableBalance
}

// GetUnstakeableBalance returns the unstakeable balance.
func (b *StakingBalances) GetUnstakeableBalance() *Balance {
	return b.UnstakeableBalance
}

// GetClaimableBalance returns the claimable balance.
func (b *StakingBalances) GetClaimableBalance() *Balance {
	return b.ClaimableBalance
}

// String returns a string representation of the staking balances.
func (b *StakingBalances) String() string {
	return "Stakeable: " + b.StakeableBalance.String() + ", Unstakeable: " + b.UnstakeableBalance.String() + ", Claimable: " + b.ClaimableBalance.String()
}

func newStakingBalancesFromModel(context *client.StakingContext) (*StakingBalances, error) {
	stakeableBalance, err := newBalanceFromModel(&context.Context.StakeableBalance)
	if err != nil {
		return nil, err
	}

	unstakeableBalance, err := newBalanceFromModel(&context.Context.UnstakeableBalance)
	if err != nil {
		return nil, err
	}

	claimableBalance, err := newBalanceFromModel(&context.Context.ClaimableBalance)
	if err != nil {
		return nil, err
	}

	return &StakingBalances{
		StakeableBalance:   stakeableBalance,
		UnstakeableBalance: unstakeableBalance,
		ClaimableBalance:   claimableBalance,
	}, nil
}
