package coinbase

import (
	"context"
	"fmt"
	"time"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/errors"
)

// StakingBalance represents a struct that holds a staking balance data.
type StakingBalance struct {
	model           *client.StakingBalance
	parsedDate      time.Time
	bondedStake     *Balance
	unbondedBalance *Balance
}

// Date returns the date of the staking balance.
func (sb *StakingBalance) Date() time.Time {
	return sb.parsedDate
}

// BondedStake returns the bonded stake of the staking balance.
func (sb *StakingBalance) BondedStake() *Balance {
	return sb.bondedStake
}

// UnbondedBalance returns the unbonded balance of the staking balance.
func (sb *StakingBalance) UnbondedBalance() *Balance {
	return sb.unbondedBalance
}

// ParticipantType returns the participant type of the staking balance.
func (sb *StakingBalance) ParticipantType() string {
	return sb.model.ParticipantType
}

// String returns the string representation of the staking balance.
func (sb *StakingBalance) String() string {
	return fmt.Sprintf(
		"StakingBalance { date: '%s' bondedStake: '%s' unbondedBalance: '%s' participantType: '%s' }",
		sb.Date().String(),
		sb.BondedStake().String(),
		sb.UnbondedBalance().String(),
		sb.ParticipantType(),
	)
}

// ListHistoricalStakingBalances fetches the historical staking balances for a given address and asset, and timeframe.
func (c *Client) ListHistoricalStakingBalances(
	ctx context.Context,
	assetId string,
	address *Address,
	startTime time.Time,
	endTime time.Time,
) ([]*StakingBalance, error) {
	req := c.client.StakeAPI.FetchHistoricalStakingBalances(ctx, address.NetworkID(), address.ID())
	req = req.AssetId(assetId)
	req = req.StartTime(startTime)
	req = req.EndTime(endTime)
	resp, httpResp, err := req.Execute()
	if err != nil {
		return nil, errors.MapToUserFacing(err, httpResp)
	}

	balances := make([]*StakingBalance, len(resp.Data))
	for i, model := range resp.Data {
		balance, err := newStakingBalanceFromModel(&model)
		if err != nil {
			return nil, err
		}
		balances[i] = balance
	}

	return balances, nil
}

func newStakingBalanceFromModel(m *client.StakingBalance) (*StakingBalance, error) {
	date, err := time.Parse(timestampFormat, m.Date)
	if err != nil {
		return nil, err
	}

	bondedStake, err := newBalanceFromModel(&m.BondedStake)
	if err != nil {
		return nil, err
	}

	unbondedBalance, err := newBalanceFromModel(&m.UnbondedBalance)
	if err != nil {
		return nil, err
	}

	return &StakingBalance{
		model:           m,
		parsedDate:      date,
		bondedStake:     bondedStake,
		unbondedBalance: unbondedBalance,
	}, nil
}
