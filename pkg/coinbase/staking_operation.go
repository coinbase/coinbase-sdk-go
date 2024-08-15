package coinbase

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

// StakeingOperationOption allows for the passing of custom options to
// the staking operation, like `mode` or `withdrawal_address`.
type StakingOperationOption func(*client.BuildStakingOperationRequest)

// WithStakingOperationMode allows for the setting of the mode of
// the staking operation (ie. `default`, `partial`, or `native`)
func WithStakingOperationMode(mode string) StakingOperationOption {
	return WithStakingOperationOption("mode", mode)
}

// WithStakingOperationOption allows for the passing of custom options
// to the staking operation, like `mode` or `withdrawal_address`.
func WithStakingOperationOption(optionKey string, optionValue string) StakingOperationOption {
	return func(op *client.BuildStakingOperationRequest) {
		op.Options[optionKey] = optionValue
	}
}

// BuildStakingOperation will build an ephemeral staking operation based on
// the passed address, assetID, action, and amount.
func (c *Client) BuildStakingOperation(
	ctx context.Context,
	address *Address,
	assetID string,
	action string,
	amount *big.Float,
	o ...StakingOperationOption,
) (*StakingOperation, error) {
	req := client.BuildStakingOperationRequest{
		NetworkId: address.NetworkID,
		AssetId:   assetID,
		AddressId: address.ID,
		Action:    action,
		Options: map[string]string{
			"mode":   "default",
			"amount": amount.String(),
		},
	}
	for _, f := range o {
		f(&req)
	}
	op, _, err := c.client.StakeAPI.BuildStakingOperation(ctx).BuildStakingOperationRequest(req).Execute()
	if err != nil {
		return nil, err
	}

	return newStakingOperationFromModel(op), nil
}

// BuildStakeOperation will build an ephemeral staking operation using the
// stake action
func (c *Client) BuildStakeOperation(
	ctx context.Context,
	address *Address,
	assetID string,
	amount *big.Float,
	o ...StakingOperationOption,
) (*StakingOperation, error) {
	return c.BuildStakingOperation(ctx, address, assetID, "stake", amount, o...)
}

// BuildStakeOperation will build an ephemeral staking operation using the
// unstake action
func (c *Client) BuildUnstakeOperation(
	ctx context.Context,
	address *Address,
	assetID string,
	amount *big.Float,
	o ...StakingOperationOption,
) (*StakingOperation, error) {
	return c.BuildStakingOperation(ctx, address, assetID, "unstake", amount, o...)
}

// BuildStakeOperation will build an ephemeral staking operation using the
// claim_stake action
func (c *Client) BuildClaimStakeOperation(
	ctx context.Context,
	address *Address,
	assetID string,
	amount *big.Float,
	o ...StakingOperationOption,
) (*StakingOperation, error) {
	return c.BuildStakingOperation(ctx, address, assetID, "claim_stake", amount, o...)
}

// FetchExternalStakingOperation loads a staking operation from the API associated
// with an address.
func (c *Client) FetchExternalStakingOperation(ctx context.Context, address *Address, id string) (*StakingOperation, error) {
	op, _, err := c.client.StakeAPI.GetExternalStakingOperation(ctx, address.NetworkID(), address.ID(), id).Execute()
	if err != nil {
		return nil, err
	}
	return newStakingOpertionFromModel(op), nil
}

// StakingOperation represents a staking operation for
// a given action, asset, and amount.
type StakingOperation struct {
	model        *client.StakingOperation
	transactions []*Transaction
}

// ID returns the StakingOperation ID
func (o *StakingOperation) ID() string {
	return o.model.Id
}

// NetworkID returns the StakingOperation network id
func (o *StakingOperation) NetworkID() string {
	return o.model.NetworkId
}

// AddressID returns the StakingOperation address id
func (o *StakingOperation) AddressID() string {
	return o.model.AddressId
}

// Status returns the StakingOperation status
func (o *StakingOperation) Status() string {
	return o.model.Status
}

// Transactions returns the transactions associated with
// the StakingOperation
func (o *StakingOperation) Transaction() []*Transaction {
	return o.transactions
}

// Sign will sign each transaction using the supplied key
func (o *StakingOperation) Sign(k *ecdsa.PrivateKey) error {
	for _, tx := range o.Transactions() {
		if !tx.IsSigned() {
			tx.Sign(k)
		}
	}
}

func newStakingOperationFromModel(m *client.StakingOperation) *StakingOperation {
	if m == nil {
		return nil
	}

	transactions := make([]*Transaction, len(m.Transactions))
	for i, tx := range m.Transactions {
		transactions[i] = newTransactionFromModel(&tx)
	}
	return &StakingOperation{
		model:        m,
		transactions: transactions,
	}
}
