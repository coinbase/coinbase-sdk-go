package coinbase

import (
	"context"
	"crypto"
	"encoding/base64"
	"fmt"
	"math/big"
	"time"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/errors"
)

// StakingOperationOption allows for the passing of custom options to
// the staking operation, like `mode` or `withdrawal_address`.
type StakingOperationOption func(*client.BuildStakingOperationRequest)

// WithStakingOperationMode allows for the setting of the mode of
// the staking operation (i.e. `default`, `partial`, or `native`)
func WithStakingOperationMode(mode string) StakingOperationOption {
	return WithStakingOperationOption("mode", mode)
}

// WithIntegratorContractAddress allows for the setting of the integrator contract address for Shared ETH staking.
func WithIntegratorContractAddress(integratorContractAddress string) StakingOperationOption {
	return WithStakingOperationOption("integrator_contract_address", integratorContractAddress)
}

// WithStakingOperationOption allows for the passing of custom options
// to the staking operation, like `mode` or `withdrawal_address`.
func WithStakingOperationOption(optionKey, optionValue string) StakingOperationOption {
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
			"mode":   StakingOperationModeDefault,
			"amount": asset.toAtomicAmount(amount).String(),
		},
	}
	for _, f := range o {
		f(&req)
	}
	op, httpResp, err := c.client.StakeAPI.BuildStakingOperation(ctx).BuildStakingOperationRequest(req).Execute()
	if err != nil {
		return nil, errors.MapToUserFacing(err, httpResp)
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

// Sign will sign each transaction using the supplied key if it isn't already signed.
// This will halt and return an error if any of the transactions fail to sign.
func (s *StakingOperation) Sign(k crypto.Signer) error {
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
		return nil, nil
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

func (c *Client) Wait(ctx context.Context, stakingOperation *StakingOperation, o ...WaitOption) error {
	options := &waitOptions{
		intervalSeconds: 5,
		timeoutSeconds:  3600,
	}

	for _, f := range o {
		f(options)
	}

	startTime := time.Now()

	for time.Since(startTime).Seconds() < float64(options.timeoutSeconds) {
		if err := c.ReloadStakingOperation(ctx, stakingOperation); err != nil {
			return err
		}

		if stakingOperation.IsTerminalState() {
			return nil
		}

		if time.Since(startTime).Seconds() > float64(options.timeoutSeconds) {
			return fmt.Errorf("staking operation timed out")
		}

		time.Sleep(time.Duration(options.intervalSeconds) * time.Second)
	}

	return fmt.Errorf("staking operation timed out")
}

// ReloadStakingOperation reloads a staking operation from the backend with the latest state.
// It ensures only newly constructed transactions are added to the staking operation and any existing transactions are untouched.
func (c *Client) ReloadStakingOperation(ctx context.Context, stakingOperation *StakingOperation) error {
	so, httpResp, err := c.client.StakeAPI.GetExternalStakingOperation(
		ctx,
		stakingOperation.NetworkID(),
		stakingOperation.AddressID(),
		stakingOperation.ID(),
	).Execute()
	if err != nil {
		return errors.MapToUserFacing(err, httpResp)
	}

	stakingOperation.model = so

	for _, tx := range so.Transactions {
		if !stakingOperation.hasTransactionByUnsignedPayload(tx.UnsignedPayload) {
			newTx, err := newTransactionFromModel(&tx)
			if err != nil {
				return err
			}

			stakingOperation.transactions = append(stakingOperation.transactions, newTx)
		}
	}
	return nil
}

// FetchStakingOperation fetches a staking operation from the backend given a networkID, addressID, and stakingOperationID.
func (c *Client) FetchStakingOperation(ctx context.Context, networkID, addressID, stakingOperationID string) (*StakingOperation, error) {
	so, httpResp, err := c.client.StakeAPI.GetExternalStakingOperation(ctx, networkID, addressID, stakingOperationID).Execute()
	if err != nil {
		return nil, errors.MapToUserFacing(err, httpResp)
	}

	return newStakingOperationFromModel(so)
}

func (s *StakingOperation) hasTransactionByUnsignedPayload(unsignedPayload string) bool {
	for _, tx := range s.Transactions() {
		if tx.UnsignedPayload() == unsignedPayload {
			return true
		}
	}
	return false
}

func (s *StakingOperation) IsTerminalState() bool {
	return s.Status() == "complete" || s.Status() == "failed"
}

func (s *StakingOperation) IsFailedState() bool {
	return s.Status() == "failed"
}

func (s *StakingOperation) IsCompleteState() bool {
	return s.Status() == "complete"
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
