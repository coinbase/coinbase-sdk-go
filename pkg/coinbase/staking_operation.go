package coinbase

import (
	"context"
	"crypto"
	"encoding/base64"
	"encoding/json"
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
	model           *client.StakingOperation
	transactionType client.StakingTransactionDetailType
	txDetails       []*StakingTransactionDetail
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

// WalletID returns the StakingOperation wallet id
func (s *StakingOperation) WalletID() string {
	if s.model.WalletId == nil {
		return ""
	}

	return *s.model.WalletId
}

// Status returns the StakingOperation status
func (s *StakingOperation) Status() string {
	return s.model.Status
}

// StakingTransactionDetails returns all the transactions and their associated metadata.
func (s *StakingOperation) StakingTransactionDetails() []*StakingTransactionDetail {
	return s.txDetails
}

func (s *StakingOperation) Transaction(index int) *Transaction {
	if index < 0 || index >= len(s.StakingTransactionDetails()) {
		return nil
	}

	detail := s.StakingTransactionDetails()[index]

	switch s.transactionType {
	case client.STAKINGTRANSACTIONDETAILTYPE_PARTIAL_ETH:
		return detail.PartialEthStakingTransactionDetail.Transaction
	case client.STAKINGTRANSACTIONDETAILTYPE_DEDICATED_ETH_STAKE:
		return detail.DedicatedEthStakeTransactionDetail.Transaction
	case client.STAKINGTRANSACTIONDETAILTYPE_SOLANA:
		return detail.SolanaStakingTransactionDetail.Transaction
	}

	return nil
}

func (s *StakingOperation) TransactionType() client.StakingTransactionDetailType {
	return s.transactionType
}

// Sign will sign each transaction using the supplied key if it isn't already signed.
// This will halt and return an error if any of the transactions fail to sign.
func (s *StakingOperation) Sign(k crypto.Signer) error {
	for _, stakingTransactionDetail := range s.StakingTransactionDetails() {
		var tx *Transaction
		switch s.transactionType {
		case client.STAKINGTRANSACTIONDETAILTYPE_PARTIAL_ETH:
			tx = stakingTransactionDetail.PartialEthStakingTransactionDetail.Transaction
		case client.STAKINGTRANSACTIONDETAILTYPE_DEDICATED_ETH_STAKE:
			tx = stakingTransactionDetail.DedicatedEthStakeTransactionDetail.Transaction
		case client.STAKINGTRANSACTIONDETAILTYPE_SOLANA:
			tx = stakingTransactionDetail.SolanaStakingTransactionDetail.Transaction
		default:
			return fmt.Errorf("unknown staking transaction type: %s", s.transactionType)
		}

		if !tx.IsSigned() {
			if err := tx.Sign(k); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *StakingOperation) SignTx(index int, k crypto.Signer) error {
	if index < 0 || index >= len(s.StakingTransactionDetails()) {
		return fmt.Errorf("index out of range")
	}

	stakingTransactionDetail := s.StakingTransactionDetails()[index]
	var tx *Transaction
	switch s.transactionType {
	case client.STAKINGTRANSACTIONDETAILTYPE_PARTIAL_ETH:
		tx = stakingTransactionDetail.PartialEthStakingTransactionDetail.Transaction
	case client.STAKINGTRANSACTIONDETAILTYPE_DEDICATED_ETH_STAKE:
		tx = stakingTransactionDetail.DedicatedEthStakeTransactionDetail.Transaction
	case client.STAKINGTRANSACTIONDETAILTYPE_SOLANA:
		tx = stakingTransactionDetail.SolanaStakingTransactionDetail.Transaction
	default:
		return fmt.Errorf("unknown staking transaction type: %s", s.transactionType)
	}

	if !tx.IsSigned() {
		if err := tx.Sign(k); err != nil {
			return err
		}
	}

	return nil
}

func (s *StakingOperation) GetSignedVoluntaryExitMessages() ([]string, error) {
	var signedVoluntaryExitMessages []string

	for _, stakingTransactionDetail := range s.StakingTransactionDetails() {
		if stakingTransactionDetail.DedicatedEthUnstakeTransactionDetail == nil {
			return signedVoluntaryExitMessages, nil
		}

		decodedMessage, err := base64.StdEncoding.DecodeString(
			stakingTransactionDetail.DedicatedEthUnstakeTransactionDetail.Metadata.VoluntaryExitMessage.SignedVoluntaryExit,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to decode signed voluntary exit message: %s", err)
		}

		signedVoluntaryExitMessages = append(signedVoluntaryExitMessages, string(decodedMessage))
	}

	return signedVoluntaryExitMessages, nil
}

type PartialEthStakingTransactionMetadata struct {
	IntegratorContractAddress string `json:"integrator_contract_address,omitempty"`
}

func (p *PartialEthStakingTransactionMetadata) MarshalJSON() ([]byte, error) {
	type alias PartialEthStakingTransactionMetadata
	return json.Marshal((*alias)(p))
}

type PartialEthStakingTransactionDetail struct {
	Transaction *Transaction
	Metadata    PartialEthStakingTransactionMetadata
}

type DedicatedEthStakeTransactionMetadata struct {
	EthereumValidatorDepositData client.EthereumValidatorDepositData `json:"ethereum_validator_deposit_data"`
}

func (p *DedicatedEthStakeTransactionMetadata) MarshalJSON() ([]byte, error) {
	type alias DedicatedEthStakeTransactionMetadata
	return json.Marshal((*alias)(p))
}

type DedicatedEthStakeTransactionDetail struct {
	Transaction *Transaction
	Metadata    DedicatedEthStakeTransactionMetadata
}

type DedicatedEthUnstakeTransactionMetadata struct {
	VoluntaryExitMessage client.SignedVoluntaryExitMessageMetadata `json:"voluntary_exit_message"`
}

func (p *DedicatedEthUnstakeTransactionMetadata) MarshalJSON() ([]byte, error) {
	type alias DedicatedEthUnstakeTransactionMetadata
	return json.Marshal((*alias)(p))
}

type DedicatedEthUnstakeTransactionDetail struct {
	Metadata DedicatedEthUnstakeTransactionMetadata
}

type SolanaStakingTransactionMetadata struct {
	ValidatorAddress string `json:"validator_address,omitempty"`
}

func (p *SolanaStakingTransactionMetadata) MarshalJSON() ([]byte, error) {
	type alias SolanaStakingTransactionMetadata
	return json.Marshal((*alias)(p))
}

type SolanaStakingTransactionDetail struct {
	Transaction *Transaction
	Metadata    SolanaStakingTransactionMetadata
}

type StakingTransactionDetail struct {
	PartialEthStakingTransactionDetail   *PartialEthStakingTransactionDetail
	DedicatedEthStakeTransactionDetail   *DedicatedEthStakeTransactionDetail
	DedicatedEthUnstakeTransactionDetail *DedicatedEthUnstakeTransactionDetail
	SolanaStakingTransactionDetail       *SolanaStakingTransactionDetail
}

func (s *StakingTransactionDetail) GetTransaction() *Transaction {
	switch {
	case s.PartialEthStakingTransactionDetail != nil:
		return s.PartialEthStakingTransactionDetail.Transaction
	case s.DedicatedEthStakeTransactionDetail != nil:
		return s.DedicatedEthStakeTransactionDetail.Transaction
	case s.DedicatedEthUnstakeTransactionDetail != nil:
		return nil
	case s.SolanaStakingTransactionDetail != nil:
		return s.SolanaStakingTransactionDetail.Transaction
	}

	return nil
}

func (s *StakingTransactionDetail) GetMetadata() interface{} {
	switch {
	case s.PartialEthStakingTransactionDetail != nil:
		return s.PartialEthStakingTransactionDetail.Metadata
	case s.DedicatedEthStakeTransactionDetail != nil:
		return s.DedicatedEthStakeTransactionDetail.Metadata
	case s.DedicatedEthUnstakeTransactionDetail != nil:
		return s.DedicatedEthUnstakeTransactionDetail.Metadata
	case s.SolanaStakingTransactionDetail != nil:
		return s.SolanaStakingTransactionDetail.Metadata
	}

	return nil
}

func (s *StakingTransactionDetail) GetPartialEthStakingTransactionDetail() *PartialEthStakingTransactionDetail {
	return s.PartialEthStakingTransactionDetail
}

func (s *StakingTransactionDetail) GetPartialEthStakingTransactionMetadata() *PartialEthStakingTransactionMetadata {
	if s.PartialEthStakingTransactionDetail == nil {
		return nil
	}

	return &s.PartialEthStakingTransactionDetail.Metadata
}

func (s *StakingTransactionDetail) GetDedicatedEthStakeTransactionDetail() *DedicatedEthStakeTransactionDetail {
	return s.DedicatedEthStakeTransactionDetail
}

func (s *StakingTransactionDetail) GetDedicatedEthStakeTransactionMetadata() *DedicatedEthStakeTransactionMetadata {
	if s.DedicatedEthStakeTransactionDetail == nil {
		return nil
	}

	return &s.DedicatedEthStakeTransactionDetail.Metadata
}

func (s *StakingTransactionDetail) GetDedicatedEthUnstakeTransactionDetail() *DedicatedEthUnstakeTransactionDetail {
	return s.DedicatedEthUnstakeTransactionDetail
}

func (s *StakingTransactionDetail) GetDedicatedEthUnstakeMetadata() *DedicatedEthUnstakeTransactionMetadata {
	if s.DedicatedEthUnstakeTransactionDetail == nil {
		return nil
	}

	return &s.DedicatedEthUnstakeTransactionDetail.Metadata
}

func (s *StakingTransactionDetail) GetSolanaStakingTransactionMetadata() *SolanaStakingTransactionMetadata {
	if s.SolanaStakingTransactionDetail == nil {
		return nil
	}

	return &s.SolanaStakingTransactionDetail.Metadata
}

func (s *StakingTransactionDetail) GetSolanaStakingTransactionDetail() *SolanaStakingTransactionDetail {
	return s.SolanaStakingTransactionDetail
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
func (c *Client) ReloadStakingOperation(ctx context.Context, currentStakingOperation *StakingOperation) error {
	so, httpResp, err := c.client.StakeAPI.GetExternalStakingOperation(
		ctx,
		currentStakingOperation.NetworkID(),
		currentStakingOperation.AddressID(),
		currentStakingOperation.ID(),
	).Execute()
	if err != nil {
		return errors.MapToUserFacing(err, httpResp)
	}

	currentStakingOperation.model = so

	for _, transactionDetail := range so.TransactionDetails {
		stakingTxInfo, err := newTxDetailFromModel(so.GetTransactionType(), transactionDetail)
		if err != nil {
			return err
		}

		networkSpecificTx, err := newNetworkTransactionFromModel(so.GetTransactionType(), transactionDetail)
		if err != nil {
			return err
		}

		if !currentStakingOperation.hasTransactionByUnsignedPayload(networkSpecificTx.UnsignedPayload()) {
			currentStakingOperation.txDetails = append(currentStakingOperation.txDetails, stakingTxInfo)
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
	for _, stakingTransactionDetails := range s.StakingTransactionDetails() {
		switch s.transactionType {
		case client.STAKINGTRANSACTIONDETAILTYPE_PARTIAL_ETH:
			if stakingTransactionDetails.GetPartialEthStakingTransactionDetail().Transaction.UnsignedPayload() == unsignedPayload {
				return true
			}
		case client.STAKINGTRANSACTIONDETAILTYPE_DEDICATED_ETH_STAKE:
			if stakingTransactionDetails.GetDedicatedEthStakeTransactionDetail().Transaction.UnsignedPayload() == unsignedPayload {
				return true
			}
		case client.STAKINGTRANSACTIONDETAILTYPE_SOLANA:
			if stakingTransactionDetails.GetSolanaStakingTransactionDetail().Transaction.UnsignedPayload() == unsignedPayload {
				return true
			}
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

	transactions := make([]*Transaction, len(m.GetTransactionDetails()))
	stakingTransactionDetails := make([]*StakingTransactionDetail, len(m.GetTransactionDetails()))

	for i, txDetail := range m.GetTransactionDetails() {
		newTx, err := newNetworkTransactionFromModel(m.GetTransactionType(), txDetail)
		if err != nil {
			return nil, err
		}

		stakingTransactionInfo, err := newTxDetailFromModel(m.GetTransactionType(), txDetail)
		if err != nil {
			return nil, err
		}

		stakingTransactionDetails[i] = stakingTransactionInfo
		if newTx != nil {
			transactions[i] = newTx
		}
	}

	return &StakingOperation{
		model:           m,
		transactionType: m.TransactionType,
		txDetails:       stakingTransactionDetails,
		//transactions:    transactions,
	}, nil
}

func newTxDetailFromModel(transactionType client.StakingTransactionDetailType, c client.StakingTransactionDetail) (*StakingTransactionDetail, error) {
	var (
		newTx *Transaction
		err   error
	)

	switch transactionType {
	case client.STAKINGTRANSACTIONDETAILTYPE_PARTIAL_ETH:
		newTx, err = newNetworkTransactionFromModel(transactionType, c)
		if err != nil {
			return nil, err
		}

		return &StakingTransactionDetail{
			PartialEthStakingTransactionDetail: &PartialEthStakingTransactionDetail{
				Transaction: newTx,
				Metadata: PartialEthStakingTransactionMetadata{
					IntegratorContractAddress: c.PartialEthStakingTransactionDetail.IntegratorContractAddress,
				},
			},
		}, nil
	case client.STAKINGTRANSACTIONDETAILTYPE_DEDICATED_ETH_STAKE:
		newTx, err = newNetworkTransactionFromModel(transactionType, c)
		if err != nil {
			return nil, err
		}
		return &StakingTransactionDetail{
			DedicatedEthStakeTransactionDetail: &DedicatedEthStakeTransactionDetail{
				Transaction: newTx,
				Metadata: DedicatedEthStakeTransactionMetadata{
					c.DedicatedEthStakeTransactionDetail.EthereumValidatorDepositData,
				},
			},
		}, nil
	case client.STAKINGTRANSACTIONDETAILTYPE_DEDICATED_ETH_UNSTAKE:
		return &StakingTransactionDetail{
			DedicatedEthUnstakeTransactionDetail: &DedicatedEthUnstakeTransactionDetail{
				Metadata: DedicatedEthUnstakeTransactionMetadata{
					VoluntaryExitMessage: c.DedicatedEthUnstakeTransactionDetail.VoluntaryExitMessage,
				},
			},
		}, nil
	case client.STAKINGTRANSACTIONDETAILTYPE_SOLANA:
		newTx, err = newNetworkTransactionFromModel(transactionType, c)
		if err != nil {
			return nil, err
		}
		return &StakingTransactionDetail{
			SolanaStakingTransactionDetail: &SolanaStakingTransactionDetail{
				Transaction: newTx,
				Metadata: SolanaStakingTransactionMetadata{
					ValidatorAddress: c.SolanaStakingTransactionDetail.ValidatorAddress,
				},
			},
		}, nil
	default:
		return nil, fmt.Errorf("unknown staking transaction type: %s", transactionType)
	}
}

func newNetworkTransactionFromModel(transactionType client.StakingTransactionDetailType, c client.StakingTransactionDetail) (*Transaction, error) {
	var (
		newTx *Transaction
		err   error
	)

	switch transactionType {
	case client.STAKINGTRANSACTIONDETAILTYPE_PARTIAL_ETH:
		partialETHStakingTx := c.PartialEthStakingTransactionDetail
		newTx, err = newTransactionFromModel(&partialETHStakingTx.Transaction)
		if err != nil {
			return nil, err
		}
	case client.STAKINGTRANSACTIONDETAILTYPE_DEDICATED_ETH_STAKE:
		dedicatedETHStakeTx := c.DedicatedEthStakeTransactionDetail
		newTx, err = newTransactionFromModel(&dedicatedETHStakeTx.Transaction)
		if err != nil {
			return nil, err
		}
	case client.STAKINGTRANSACTIONDETAILTYPE_SOLANA:
		solanaStakingTx := c.SolanaStakingTransactionDetail
		newTx, err = newTransactionFromModel(&solanaStakingTx.Transaction)
		if err != nil {
			return nil, err
		}
	default:
		newTx = nil
	}

	return newTx, nil
}
