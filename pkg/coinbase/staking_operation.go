package coinbase

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"math/big"
	"time"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

// StakingOperationOption allows for the passing of custom options to
// the staking operation, like `mode` or `withdrawal_address`.
type StakingOperationOption func(*client.BuildStakingOperationRequest)

// WithStakingOperationMode allows for the setting of the mode of
// the staking operation (ie. `default`, `partial`, or `native`)
func WithStakingOperationMode(mode string) StakingOperationOption {
	return WithStakingOperationOption("mode", mode)
}

// WithStakingOperationImmediate allows for the setting of the immediate flag
// specifically for Dedicated ETH Staking whether to immediate unstake or not. (i.e. `true` or `false`)
func WithStakingOperationImmediate(immediate string) StakingOperationOption {
	return WithStakingOperationOption("immediate", immediate)
}

// WithStakingOperationOption allows for the passing of custom options
// to the staking operation, like `mode` or `withdrawal_address`.
func WithStakingOperationOption(optionKey string, optionValue string) StakingOperationOption {
	return func(op *client.BuildStakingOperationRequest) {
		op.Options[optionKey] = optionValue
	}
}

type waitOptions struct {
	intervalSeconds int
	timeoutSeconds  int
}

// WaitOption allows for the passing of custom options to the wait function.
type WaitOption func(*waitOptions)

// WithWaitIntervalSeconds sets the interval in seconds to wait between
// polling the staking operation.
func WithWaitIntervalSeconds(intervalSeconds int) WaitOption {
	return func(o *waitOptions) {
		o.intervalSeconds = intervalSeconds
	}
}

// WithWaitTimeoutSeconds sets the timeout in seconds to wait for the
// staking operation to complete.
func WithWaitTimeoutSeconds(timeoutSeconds int) WaitOption {
	return func(o *waitOptions) {
		o.timeoutSeconds = timeoutSeconds
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

	asset, err := c.fetchAsset(ctx, address.NetworkID(), assetID)
	if err != nil {
		return nil, err
	}

	req := client.BuildStakingOperationRequest{
		NetworkId: normalizeNetwork(address.NetworkID()),
		AssetId:   assetID,
		AddressId: address.ID(),
		Action:    action,
		Options: map[string]string{
			"amount": asset.toAtomicAmount(amount).String(),
		},
	}
	for _, f := range o {
		f(&req)
	}
	op, _, err := c.client.StakeAPI.BuildStakingOperation(ctx).BuildStakingOperationRequest(req).Execute()
	if err != nil {
		return nil, err
	}

	return newStakingOperationFromModel(op)
}

// BuildStakeOperation will build an ephemeral staking operation using the
// stake action
func (c *Client) BuildStakeOperation(
	ctx context.Context,
	amount *big.Float,
	assetID string,
	address *Address,
	o ...StakingOperationOption,
) (*StakingOperation, error) {
	return c.BuildStakingOperation(ctx, address, assetID, "stake", amount, o...)
}

// BuildUnstakeOperation will build an ephemeral staking operation using the
// unstake action
func (c *Client) BuildUnstakeOperation(
	ctx context.Context,
	amount *big.Float,
	assetID string,
	address *Address,
	o ...StakingOperationOption,
) (*StakingOperation, error) {
	return c.BuildStakingOperation(ctx, address, assetID, "unstake", amount, o...)
}

// BuildClaimStakeOperation will build an ephemeral staking operation using the
// claim_stake action
func (c *Client) BuildClaimStakeOperation(
	ctx context.Context,
	amount *big.Float,
	assetID string,
	address *Address,
	o ...StakingOperationOption,
) (*StakingOperation, error) {
	return c.BuildStakingOperation(ctx, address, assetID, "claim_stake", amount, o...)
}

// StakingOperation represents a staking operation for
// a given action, asset, and amount.
type StakingOperation struct {
	model        *client.StakingOperation
	transactions []*Transaction
}

// ID returns the StakingOperation ID
func (s *StakingOperation) ID() string {
	return s.model.Id
}

// NetworkID returns the StakingOperation network id
func (s *StakingOperation) NetworkID() string {
	return s.model.NetworkId
}

// AddressID returns the StakingOperation address id
func (s *StakingOperation) AddressID() string {
	return s.model.AddressId
}

// Status returns the StakingOperation status
func (s *StakingOperation) Status() string {
	return s.model.Status
}

// Transactions returns the transactions associated with
// the StakingOperation
func (s *StakingOperation) Transactions() []*Transaction {
	return s.transactions
}

// Sign will sign each transaction using the supplied key
// This will halt and return an error if any of the transactions
// fail to sign.
func (s *StakingOperation) Sign(k *ecdsa.PrivateKey) error {
	for _, tx := range s.Transactions() {
		if !tx.IsSigned() {
			err := tx.Sign(k)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *StakingOperation) GetSignedVoluntaryExitMessages() ([]string, error) {
	var signedVoluntaryExitMessages []string

	stakingOperationMetadata := s.model.GetMetadata().ArrayOfSignedVoluntaryExitMessageMetadata

	if s.model.Metadata == nil {
		return []string{}, nil
	}

	for _, metadata := range *stakingOperationMetadata {
		decodedMessage, err := base64.StdEncoding.DecodeString(metadata.SignedVoluntaryExit)
		if err != nil {
			return nil, fmt.Errorf("failed to decode signed voluntary exit message: %s", err)
		}
		signedVoluntaryExitMessages = append(signedVoluntaryExitMessages, string(decodedMessage))
	}

	return signedVoluntaryExitMessages, nil
}

func (c *Client) Wait(ctx context.Context, stakingOperation *StakingOperation, o ...WaitOption) (*StakingOperation, error) {

	options := &waitOptions{
		intervalSeconds: 5,
		timeoutSeconds:  3600,
	}

	for _, f := range o {
		f(options)
	}

	startTime := time.Now()

	for time.Since(startTime).Seconds() < float64(options.timeoutSeconds) {
		so, err := c.fetchExternalStakingOperation(ctx, stakingOperation)
		if err != nil {
			return stakingOperation, err
		}
		stakingOperation = so

		if stakingOperation.isTerminalState() {
			return stakingOperation, nil
		}

		if time.Since(startTime).Seconds() > float64(options.timeoutSeconds) {
			return stakingOperation, fmt.Errorf("staking operation timed out")
		}

		time.Sleep(time.Duration(options.intervalSeconds) * time.Second)
	}

	return stakingOperation, fmt.Errorf("staking operation timed out")
}

// FetchExternalStakingOperation loads a staking operation from the API associated
// with an address.
func (c *Client) fetchExternalStakingOperation(ctx context.Context, stakingOperation *StakingOperation) (*StakingOperation, error) {
	so, httpRes, err := c.client.StakeAPI.GetExternalStakingOperation(
		ctx,
		stakingOperation.NetworkID(),
		stakingOperation.AddressID(),
		stakingOperation.ID(),
	).Execute()
	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch staking operation: %s", httpRes.Status)
	}

	return newStakingOperationFromModel(so)
}

func (s *StakingOperation) isTerminalState() bool {
	return s.Status() == "complete" || s.Status() == "failed"
}

func newStakingOperationFromModel(m *client.StakingOperation) (*StakingOperation, error) {
	if m == nil {
		return nil, fmt.Errorf("staking operation model is nil")
	}

	transactions := make([]*Transaction, len(m.Transactions))
	for i, tx := range m.Transactions {
		newTx, err := newTransactionFromModel(&tx)
		if err != nil {
			return nil, err
		}
		transactions[i] = newTx
	}
	return &StakingOperation{
		model:        m,
		transactions: transactions,
	}, nil
}
