package coinbase

import (
	"context"
	"fmt"
	"testing"

	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ValidatorSuite struct {
	suite.Suite

	mockValidatorsAPI *mocks.ValidatorsAPI
	client            *Client
}

func TestValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorSuite))
}

func (s *ValidatorSuite) SetupTest() {
	s.mockValidatorsAPI = mocks.NewValidatorsAPI(s.T())
	s.client = &Client{client: &api.APIClient{
		ValidatorsAPI: s.mockValidatorsAPI,
	}}
}

func (s *ValidatorSuite) TearDownTest() {
	s.mockValidatorsAPI.AssertExpectations(s.T())
}

func (s *ValidatorSuite) TestListValidators_Success() {
	ctx := context.Background()
	networkId := "test-network"
	assetId := "test-asset"

	mockValidators := &api.ValidatorList{
		Data: []api.Validator{
			{
				ValidatorId: "validator-1",
				Status:      api.VALIDATORSTATUS_ACTIVE,
			},
		},
	}

	s.mockSuccessfulListValidators(ctx, networkId, assetId, mockValidators)

	validators, err := s.client.ListValidators(ctx, networkId, assetId)

	s.Assert().NoError(err)
	s.Len(validators, 1)
	s.Equal("validator-1", validators[0].ID())
	s.Equal(api.VALIDATORSTATUS_ACTIVE, validators[0].Status())
}

// Since options are a slice a user can pass a nil which we should try and handle.
func (s *ValidatorSuite) TestListValidators_Success_WithNilOptions() {
	ctx := context.Background()
	networkId := "test-network"
	assetId := "test-asset"

	mockValidators := &api.ValidatorList{
		Data: []api.Validator{
			{
				ValidatorId: "validator-1",
				Status:      api.VALIDATORSTATUS_ACTIVE,
			},
		},
	}

	s.mockSuccessfulListValidators(ctx, networkId, assetId, mockValidators)

	validators, err := s.client.ListValidators(ctx, networkId, assetId, nil)

	s.Assert().NoError(err)
	s.Len(validators, 1)
	s.Equal("validator-1", validators[0].ID())
	s.Equal(api.VALIDATORSTATUS_ACTIVE, validators[0].Status())
}

func (s *ValidatorSuite) TestListValidators_Failure() {
	ctx := context.Background()
	networkId := "test-network"
	assetId := "test-asset"
	errorMessage := "some error calling api"

	s.mockFailedListValidators(ctx, networkId, assetId, fmt.Errorf(errorMessage))

	validators, err := s.client.ListValidators(ctx, networkId, assetId)

	s.Assert().Nil(validators)
	s.EqualError(err, errorMessage)
}

func (s *ValidatorSuite) TestGetValidator_Success() {
	ctx := context.Background()
	networkId := "test-network"
	assetId := "test-asset"
	validatorId := "validator-1"

	s.mockSuccessfulGetValidator(ctx, networkId, assetId, validatorId, &api.Validator{
		ValidatorId: validatorId,
		Status:      api.VALIDATORSTATUS_ACTIVE,
	})

	validator, err := s.client.GetValidator(ctx, networkId, assetId, validatorId)

	s.Assert().NoError(err)
	s.Equal(validatorId, validator.ID())
	s.Equal(api.VALIDATORSTATUS_ACTIVE, validator.Status())
	s.Equal("Validator { Id: 'validator-1' Status: 'active' }", validator.ToString())
}

func (s *ValidatorSuite) TestGetValidator_Failure() {
	ctx := context.Background()
	networkId := "test-network"
	assetId := "test-asset"
	validatorId := "validator-1"
	errorMessage := "some error calling api"

	s.mockFailedGetValidator(ctx, networkId, assetId, validatorId, fmt.Errorf(errorMessage))

	validator, err := s.client.GetValidator(ctx, networkId, assetId, validatorId)

	s.Assert().Empty(validator)
	s.EqualError(err, errorMessage)
}

func (s *ValidatorSuite) mockSuccessfulListValidators(ctx context.Context, networkId string, assetId string, mockValidators *api.ValidatorList) {
	s.T().Helper()

	s.mockValidatorsAPI.On("ListValidators", ctx, networkId, assetId).Return(api.ApiListValidatorsRequest{
		ApiService: s.mockValidatorsAPI,
	}, nil, nil).Once()

	s.mockValidatorsAPI.On("ListValidatorsExecute", mock.Anything).Return(mockValidators, nil, nil).Once()
}

func (s *ValidatorSuite) mockFailedListValidators(ctx context.Context, networkId string, assetId string, err error) {
	s.T().Helper()

	s.mockValidatorsAPI.On("ListValidators", ctx, networkId, assetId).Return(api.ApiListValidatorsRequest{
		ApiService: s.mockValidatorsAPI,
	}, nil, nil).Once()

	s.mockValidatorsAPI.On("ListValidatorsExecute", mock.Anything).Return(nil, nil, err).Once()
}

func (s *ValidatorSuite) mockSuccessfulGetValidator(ctx context.Context, networkId string, assetId string, validatorId string, validator *api.Validator) {
	s.T().Helper()

	s.mockValidatorsAPI.On("GetValidator", ctx, networkId, assetId, validatorId).Return(api.ApiGetValidatorRequest{
		ApiService: s.mockValidatorsAPI,
	}, nil, nil)

	s.mockValidatorsAPI.On("GetValidatorExecute", mock.Anything).Return(validator, nil, nil)
}

func (s *ValidatorSuite) mockFailedGetValidator(ctx context.Context, networkId string, assetId string, validatorId string, err error) {
	s.T().Helper()

	s.mockValidatorsAPI.On("GetValidator", ctx, networkId, assetId, validatorId).Return(api.ApiGetValidatorRequest{
		ApiService: s.mockValidatorsAPI,
	}, nil, nil)

	s.mockValidatorsAPI.On("GetValidatorExecute", mock.Anything).Return(nil, nil, err)
}
