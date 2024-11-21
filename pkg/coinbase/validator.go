package coinbase

import (
	"context"
	"fmt"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/errors"
)

type ValidatorStatus string

const (
	ValidatorStatusUnknown             ValidatorStatus = "unknown"
	ValidatorStatusProvisioning        ValidatorStatus = "provisioning"
	ValidatorStatusProvisioned         ValidatorStatus = "provisioned"
	ValidatorStatusDeposited           ValidatorStatus = "deposited"
	ValidatorStatusPendingActivation   ValidatorStatus = "pending_activation"
	ValidatorStatusActive              ValidatorStatus = "active"
	ValidatorStatusExiting             ValidatorStatus = "exiting"
	ValidatorStatusExited              ValidatorStatus = "exited"
	ValidatorStatusWithdrawalAvailable ValidatorStatus = "withdrawal_available"
	ValidatorStatusWithdrawalComplete  ValidatorStatus = "withdrawal_complete"
	ValidatorStatusActiveSlashed       ValidatorStatus = "active_slashed"
	ValidatorStatusExitedSlashed       ValidatorStatus = "exited_slashed"
	ValidatorStatusReaped              ValidatorStatus = "reaped"
)

// getAPIValidatorStatus maps the user facing ValidatorStatus to api facing ValidatorStatus
func getAPIValidatorStatus(status *ValidatorStatus) client.ValidatorStatus {
	if status == nil {
		return client.VALIDATORSTATUS_UNKNOWN
	}

	switch *status {
	case ValidatorStatusUnknown:
		return client.VALIDATORSTATUS_UNKNOWN
	case ValidatorStatusProvisioning:
		return client.VALIDATORSTATUS_PROVISIONING
	case ValidatorStatusProvisioned:
		return client.VALIDATORSTATUS_PROVISIONED
	case ValidatorStatusDeposited:
		return client.VALIDATORSTATUS_DEPOSITED
	case ValidatorStatusPendingActivation:
		return client.VALIDATORSTATUS_PENDING_ACTIVATION
	case ValidatorStatusActive:
		return client.VALIDATORSTATUS_ACTIVE
	case ValidatorStatusExiting:
		return client.VALIDATORSTATUS_EXITING
	case ValidatorStatusExited:
		return client.VALIDATORSTATUS_EXITED
	case ValidatorStatusWithdrawalAvailable:
		return client.VALIDATORSTATUS_WITHDRAWAL_AVAILABLE
	case ValidatorStatusWithdrawalComplete:
		return client.VALIDATORSTATUS_WITHDRAWAL_COMPLETE
	case ValidatorStatusActiveSlashed:
		return client.VALIDATORSTATUS_ACTIVE_SLASHED
	case ValidatorStatusExitedSlashed:
		return client.VALIDATORSTATUS_EXITED_SLASHED
	case ValidatorStatusReaped:
		return client.VALIDATORSTATUS_REAPED
	default:
		return client.VALIDATORSTATUS_UNKNOWN
	}
}

// ListValidatorsOption is a function that modifies the ListValidators request.
type ListValidatorsOption func(client.ApiListValidatorsRequest) client.ApiListValidatorsRequest

// WithListValidatorsStatusOption filters the list of validators by status.
// Note: The Go generated code doesn't allow updating the request by reference.
// As a result we are forced to return back a copy of the request.
func WithListValidatorsStatusOption(status ValidatorStatus) ListValidatorsOption {
	return func(request client.ApiListValidatorsRequest) client.ApiListValidatorsRequest {
		return request.Status(getAPIValidatorStatus(&status))
	}
}

type Validator struct {
	model client.Validator
}

func NewValidator(validator client.Validator) Validator {
	return Validator{
		model: validator,
	}
}

func (v Validator) ID() string {
	return v.model.ValidatorId
}

func (v Validator) Status() client.ValidatorStatus {
	return v.model.Status
}

func (v Validator) ToString() string {
	return fmt.Sprintf(
		"Validator { Id: '%s' Status: '%s' }",
		v.ID(),
		v.Status(),
	)
}

func (c *Client) ListValidators(
	ctx context.Context,
	networkId string,
	assetId string,
	o ...ListValidatorsOption,
) ([]Validator, error) {
	listValidatorReq := c.client.StakeAPI.ListValidators(ctx, normalizeNetwork(networkId), assetId)

	for _, f := range o {
		if f == nil {
			continue
		}

		listValidatorReq = f(listValidatorReq)
	}

	validatorList, httpResp, err := listValidatorReq.Execute()
	if err != nil {
		return nil, errors.MapToUserFacing(err, httpResp)
	}

	validators := make([]Validator, len(validatorList.GetData()))
	for i, validator := range validatorList.GetData() {
		validators[i] = NewValidator(validator)
	}

	return validators, nil
}

func (c *Client) GetValidator(ctx context.Context, networkId string, assetId string, validatorId string) (Validator, error) {
	validator, httpResp, err := c.client.StakeAPI.GetValidator(
		ctx,
		normalizeNetwork(networkId),
		assetId,
		validatorId,
	).Execute()
	if err != nil {
		return Validator{}, errors.MapToUserFacing(err, httpResp)
	}

	return NewValidator(*validator), nil
}
