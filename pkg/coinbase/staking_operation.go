package coinbase

import (
	"context"
	"crypto"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
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

// WithNativeStakingOperationMode allows for the setting of the staking operation mode to "native".
func WithNativeStakingOperationMode() StakingOperationOption {
	return WithStakingOperationMode(StakingOperationModeNative)
}

// WithPartialStakingOperationMode allows for the setting of the staking operation mode to "partial".
func WithPartialStakingOperationMode() StakingOperationOption {
	return WithStakingOperationMode(StakingOperationModePartial)
}

// WithWithdrawalCredentialType allows for the setting of the withdrawal credential type.
func WithWithdrawalCredentialType(credentialType string) StakingOperationOption {
	return WithStakingOperationOption("withdrawal_credential_type", credentialType)
}

// With0x01WithdrawalCredentialType allows for the setting of the withdrawal credential type to "0x01".
func With0x01WithdrawalCredentialType() StakingOperationOption {
	return WithWithdrawalCredentialType("0x01")
}

// With0x02WithdrawalCredentialType allows for the setting of the withdrawal credential type to "0x02".
func With0x02WithdrawalCredentialType() StakingOperationOption {
	return WithWithdrawalCredentialType("0x02")
}

// WithUnstakeType allows for the setting of the unstake type.
func WithUnstakeType(unstakeType string) StakingOperationOption {
	return WithStakingOperationOption("unstake_type", unstakeType)
}

// WithUnstakeTypeConsensus allows for the setting of the unstake type to "consensus".
func WithUnstakeTypeConsensus() StakingOperationOption {
	return WithUnstakeType("consensus")
}

// WithUnstakeTypeExecution allows for the setting of the unstake type to "execution".
func WithUnstakeTypeExecution() StakingOperationOption {
	return WithUnstakeType("execution")
}

// WithIntegratorContractAddress allows for the setting of the integrator contract address for Shared ETH staking.
func WithIntegratorContractAddress(integratorContractAddress string) StakingOperationOption {
	return WithStakingOperationOption("integrator_contract_address", integratorContractAddress)
}

// WithImmediateUnstake allows for the setting of the immediate unstake option which results in immediately exiting the Ethereum validator.
// Don't use this option if you want to generate pre-signed exit messages only.
func WithImmediateUnstake() StakingOperationOption {
	return WithStakingOperationOption("immediate", "true")
}

// WithFeeRecipientAddress allows for the setting of the fee recipient address.
func WithFeeRecipientAddress(feeRecipientAddress string) StakingOperationOption {
	return WithStakingOperationOption("fee_recipient_address", feeRecipientAddress)
}

// WithRewardSplitterPlanID allows for the setting of the reward splitter plan ID.
func WithRewardSplitterPlanID(rewardSplitterPlanID string) StakingOperationOption {
	return WithStakingOperationOption("reward_splitter_plan_id", rewardSplitterPlanID)
}

// WithSourceValidatorPublicKey allows for the setting of the source validator public key.
func WithSourceValidatorPublicKey(sourceValidatorPublicKey string) StakingOperationOption {
	return WithStakingOperationOption("source_validator_pubkey", sourceValidatorPublicKey)
}

// WithTargetValidatorPublicKey allows for the setting of the target validator public key.
func WithTargetValidatorPublicKey(targetValidatorPublicKey string) StakingOperationOption {
	return WithStakingOperationOption("target_validator_pubkey", targetValidatorPublicKey)
}

type ConsensusLayerExitOptionBuilder struct {
	validatorPubKeys []string
}

func NewConsensusLayerExitOptionBuilder() *ConsensusLayerExitOptionBuilder {
	return &ConsensusLayerExitOptionBuilder{
		validatorPubKeys: make([]string, 0),
	}
}

func (c *ConsensusLayerExitOptionBuilder) AddValidator(publicKey string) {
	for _, existingKey := range c.validatorPubKeys {
		if existingKey == publicKey {
			return // Skip adding duplicate public keys
		}
	}

	c.validatorPubKeys = append(c.validatorPubKeys, publicKey)
}

func WithConsensusLayerExit(builder *ConsensusLayerExitOptionBuilder) StakingOperationOption {
	return func(op *client.BuildStakingOperationRequest) {
		op.Options["unstake_type"] = "consensus"
		op.Options["validator_pub_keys"] = strings.Join(builder.validatorPubKeys, ",")
	}
}

// ExecutionLayerWithdrawalsOptionBuilder builds the options for Native ETH execution layer validator withdrawals as defined in https://eips.ethereum.org/EIPS/eip-7002.
type ExecutionLayerWithdrawalsOptionBuilder struct {
	assetUnitConverter func(*big.Float) string
	validatorAmounts   map[string]string
	buffer             string
}

// NewExecutionLayerWithdrawalsOptionBuilder creates a new builder for Native ETH execution layer validator withdrawals.
func NewExecutionLayerWithdrawalsOptionBuilder(
	ctx context.Context,
	c *Client,
	address *Address,
) (*ExecutionLayerWithdrawalsOptionBuilder, error) {
	asset, err := c.fetchAsset(ctx, address.NetworkID(), Eth)
	if err != nil {
		return nil, err
	}

	return &ExecutionLayerWithdrawalsOptionBuilder{
		assetUnitConverter: func(amount *big.Float) string { return asset.ToAtomicAmount(amount).String() },
		validatorAmounts:   make(map[string]string),
		buffer:             "{}",
	}, nil
}

// AddValidatorWithdrawal adds a Native ETH validator public key and amount to the list of validators to generate withdrawal transactions for.
// An amount of "0" indicates a full withdrawal.
func (b *ExecutionLayerWithdrawalsOptionBuilder) AddValidatorWithdrawal(publicKey string, amount *big.Float) error {
	if amount == nil {
		return nil
	}

	b.validatorAmounts[publicKey] = b.assetUnitConverter(amount)

	buf, err := json.Marshal(b.validatorAmounts)
	if err != nil {
		return err
	}

	b.buffer = string(buf)

	return nil
}

// WithExecutionLayerWithdrawals allows for Native ETH execution layer withdrawals as defined in https://eips.ethereum.org/EIPS/eip-7002.
// Selected validators must have been upgraded to withdrawal credentials of type "0x02".
func WithExecutionLayerWithdrawals(builder *ExecutionLayerWithdrawalsOptionBuilder) StakingOperationOption {
	return func(op *client.BuildStakingOperationRequest) {
		op.Options["unstake_type"] = "execution"
		op.Options["validator_unstake_amounts"] = builder.buffer

		// Top-level "amount" is excluded for execution layer withdrawals until we support amount based picking
		// up of validators to unstake for execution layer withdrawals.
		delete(op.Options, "amount")
	}
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
			"mode": StakingOperationModeDefault,
		},
	}

	if amount != nil {
		req.Options["amount"] = asset.ToAtomicAmount(amount).String()
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

// BuildValidatorConsolidationOperation will build an ephemeral staking operation using the consolidate action
// to help consolidate validators post Pectra.
func (c *Client) BuildValidatorConsolidationOperation(
	ctx context.Context,
	address *Address,
	o ...StakingOperationOption,
) (*StakingOperation, error) {
	o = append(o, WithNativeStakingOperationMode())
	return c.BuildStakingOperation(ctx, address, Eth, "consolidate", nil, o...)
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
