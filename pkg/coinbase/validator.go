package coinbase

import (
	"context"
	"fmt"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

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

func (c *Client) ListValidators(ctx context.Context, networkId string, assetId string) ([]Validator, error) {
	validatorList, _, err := c.client.ValidatorsAPI.ListValidators(
		ctx,
		normalizeNetwork(networkId),
		assetId,
	).Execute()
	if err != nil {
		return nil, err
	}

	validators := make([]Validator, len(validatorList.GetData()))
	for i, validator := range validatorList.GetData() {
		validators[i] = NewValidator(validator)
	}

	return validators, nil
}

func (c *Client) GetValidator(ctx context.Context, networkId string, assetId string, validatorId string) (Validator, error) {
	validator, _, err := c.client.ValidatorsAPI.GetValidator(
		ctx,
		normalizeNetwork(networkId),
		assetId,
		validatorId,
	).Execute()
	if err != nil {
		return Validator{}, err
	}

	return NewValidator(*validator), nil
}
