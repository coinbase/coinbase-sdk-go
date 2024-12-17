package coinbase

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ValidatorSuite struct {
	suite.Suite

	mockStakeAPI *mocks.StakeAPI
	client       *Client
}

func TestValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorSuite))
}

func (s *ValidatorSuite) SetupTest() {
	s.mockStakeAPI = mocks.NewStakeAPI(s.T())
	s.client = &Client{client: &api.APIClient{
		StakeAPI: s.mockStakeAPI,
	}}
}

func (s *ValidatorSuite) TearDownTest() {
	s.mockStakeAPI.AssertExpectations(s.T())
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
	s.EqualError(err, "APIError{HttpStatusCode: 500, Code: unknown, Message: some error calling api}")
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
	s.EqualError(err, "APIError{HttpStatusCode: 500, Code: unknown, Message: some error calling api}")
}

func (s *ValidatorSuite) TestGetters() {
	validator := NewValidator(api.Validator{
		ValidatorId: "validator-1",
		NetworkId:   EthereumHolesky,
		AssetId:     Eth,
		Status:      api.VALIDATORSTATUS_ACTIVE,
		Details: &api.ValidatorDetails{
			EthereumValidatorMetadata: &api.EthereumValidatorMetadata{
				Index:             "0",
				PublicKey:         "public-key-1",
				WithdrawalAddress: "withdrawal-address-1",
				Slashed:           false,
				ActivationEpoch:   "epoch-1",
				ExitEpoch:         "exit-epoch-2",
				WithdrawableEpoch: "withdrawable-epoch-3",
				Balance: api.Balance{
					Amount: "100",
					Asset: api.Asset{
						NetworkId: EthereumHolesky,
						AssetId:   Eth,
					},
				},
				EffectiveBalance: api.Balance{
					Amount: "200",
					Asset: api.Asset{
						NetworkId: EthereumHolesky,
						AssetId:   Eth,
					},
				},
			},
		},
	})

	s.Assert().Equal("validator-1", validator.ID())
	s.Assert().Equal(EthereumHolesky, validator.NetworkID())
	s.Assert().Equal(Eth, validator.AssetID())
	s.Assert().Equal(api.VALIDATORSTATUS_ACTIVE, validator.Status())
	s.Assert().Equal("0", validator.Index())
	s.Assert().Equal("public-key-1", validator.PublicKey())
	s.Assert().Equal("withdrawal-address-1", validator.WithdrawalAddress())
	s.Assert().Equal(false, validator.Slashed())
	s.Assert().Equal("epoch-1", validator.ActivationEpoch())
	s.Assert().Equal("exit-epoch-2", validator.ExitEpoch())
	s.Assert().Equal("withdrawable-epoch-3", validator.WithdrawableEpoch())
	s.Assert().Equal("100", validator.Balance().Amount)
	s.Assert().Equal(EthereumHolesky, validator.Balance().Asset.NetworkId)
	s.Assert().Equal(Eth, validator.Balance().Asset.AssetId)
	s.Assert().Equal("200", validator.EffectiveBalance().Amount)
	s.Assert().Equal(EthereumHolesky, validator.EffectiveBalance().Asset.NetworkId)
	s.Assert().Equal(Eth, validator.EffectiveBalance().Asset.AssetId)
}

func (s *ValidatorSuite) mockSuccessfulListValidators(ctx context.Context, networkId string, assetId string, mockValidators *api.ValidatorList) {
	s.T().Helper()

	s.mockStakeAPI.On("ListValidators", ctx, networkId, assetId).Return(api.ApiListValidatorsRequest{
		ApiService: s.mockStakeAPI,
	}, nil, nil).Once()

	s.mockStakeAPI.On("ListValidatorsExecute", mock.Anything).Return(mockValidators, nil, nil).Once()
}

func (s *ValidatorSuite) mockFailedListValidators(ctx context.Context, networkId string, assetId string, err error) {
	s.T().Helper()

	s.mockStakeAPI.On("ListValidators", ctx, networkId, assetId).Return(api.ApiListValidatorsRequest{
		ApiService: s.mockStakeAPI,
	}, nil, nil).Once()

	s.mockStakeAPI.On("ListValidatorsExecute", mock.Anything).Return(nil, internalFailureHttpResponse(), err).Once()
}

func (s *ValidatorSuite) mockSuccessfulGetValidator(ctx context.Context, networkId string, assetId string, validatorId string, validator *api.Validator) {
	s.T().Helper()

	s.mockStakeAPI.On("GetValidator", ctx, networkId, assetId, validatorId).Return(api.ApiGetValidatorRequest{
		ApiService: s.mockStakeAPI,
	}, nil, nil)

	s.mockStakeAPI.On("GetValidatorExecute", mock.Anything).Return(validator, nil, nil)
}

func (s *ValidatorSuite) mockFailedGetValidator(ctx context.Context, networkId string, assetId string, validatorId string, err error) {
	s.T().Helper()

	s.mockStakeAPI.On("GetValidator", ctx, networkId, assetId, validatorId).Return(api.ApiGetValidatorRequest{
		ApiService: s.mockStakeAPI,
	}, nil, nil)

	s.mockStakeAPI.On("GetValidatorExecute", mock.Anything).Return(nil, internalFailureHttpResponse(), err)
}

func internalFailureHttpResponse() *http.Response {
	return &http.Response{
		StatusCode: http.StatusInternalServerError,
	}
}
