package coinbase

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/mocks"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type StakingOperationSuite struct {
	suite.Suite
}

func TestStakingOperationSuite(t *testing.T) {
	suite.Run(t, new(StakingOperationSuite))
}

func (s *StakingOperationSuite) TestStakingOperation_Wait_Success() {
	mc := &mockController{
		stakeAPI: mocks.NewStakeAPI(s.T()),
	}
	mockGetExternalStakingOperation(s.T(), mc.stakeAPI, http.StatusOK, "pending")

	c := &Client{
		client: &api.APIClient{
			StakeAPI:  mc.stakeAPI,
			AssetsAPI: mc.assetsAPI,
		},
	}

	so, err := mockStakingOperation(s.T(), "pending")
	s.NoError(err, "failed to create staking operation")
	key, err := crypto.GenerateKey()
	s.NoError(err, "failed to generate ecdsa key")
	err = so.Sign(key)
	s.NoError(err, "failed to sign staking operation")
	signedPayload := so.Transactions()[0].SignedPayload()
	s.NotEmpty(signedPayload, "signed payload should not be empty")
	so, err = c.Wait(context.Background(), so)
	s.NoError(err, "staking operation wait should not error")
	s.Equal("complete", so.Status(), "staking operation status should be complete")
	s.Equal(1, len(so.Transactions()), "staking operation should have 1 transaction")
	s.Equal(signedPayload, so.Transactions()[0].SignedPayload(), "staking operation signed payload should not have changed")
}

func (s *StakingOperationSuite) TestStakingOperation_Wait_Success_CustomOptions() {
	mc := &mockController{
		stakeAPI: mocks.NewStakeAPI(s.T()),
	}
	mockGetExternalStakingOperation(s.T(), mc.stakeAPI, http.StatusOK, "pending")

	c := &Client{
		client: &api.APIClient{
			StakeAPI:  mc.stakeAPI,
			AssetsAPI: mc.assetsAPI,
		},
	}

	so, err := mockStakingOperation(s.T(), "pending")
	s.NoError(err, "staking operation creation should not error")
	so, err = c.Wait(
		context.Background(),
		so,
		WithWaitIntervalSeconds(1),
		WithWaitTimeoutSeconds(3),
	)
	s.NoError(err, "staking operation wait should not error")
	s.Equal("complete", so.Status(), "staking operation status should be complete")
	s.Equal(1, len(so.Transactions()), "staking operation should have 1 transaction")
}

func (s *StakingOperationSuite) TestStakingOperation_Wait_Failure() {
	tests := map[string]struct {
		soStatus      string
		setup         func(*mockController)
		expectedError string
	}{
		"fail to fetch staking operation": {
			setup: func(c *mockController) {
				c.stakeAPI.On("GetExternalStakingOperation", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
					api.ApiGetExternalStakingOperationRequest{ApiService: c.stakeAPI},
				).Once()
				c.stakeAPI.On("GetExternalStakingOperationExecute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
					nil,
					nil,
					fmt.Errorf("failed to fetch staking operation"),
				).Once()
			},
			expectedError: "failed to fetch staking operation",
		},
	}

	for name, tt := range tests {
		s.T().Run(name, func(t *testing.T) {
			mc := &mockController{
				stakeAPI: mocks.NewStakeAPI(t),
			}
			tt.setup(mc)

			c := &Client{
				client: &api.APIClient{
					StakeAPI:  mc.stakeAPI,
					AssetsAPI: mc.assetsAPI,
				},
			}

			so, err := mockStakingOperation(t, tt.soStatus)
			s.NoError(err, "staking operation creation should not error")
			_, err = c.Wait(context.Background(), so)
			s.Error(err, "staking operation wait should error")
		})
	}
}

func (s *StakingOperationSuite) TestStakingOperation_GetSignedVoluntaryExitMessages() {
	stakingOperation, err := mockStakingOperation(s.T(), "pending")
	s.NoError(err, "staking operation creation should not error")

	SignedVoluntaryExitMessages, err := stakingOperation.GetSignedVoluntaryExitMessages()
	s.NoError(err, "get signed voluntary exit messages should not error")

	s.Equal(1, len(SignedVoluntaryExitMessages), "signed voluntary exit messages should have length 1")
	s.Equal("test-data", SignedVoluntaryExitMessages[0], "signed voluntary exit messages should match")
}

func mockStakingOperation(t *testing.T, status string) (*StakingOperation, error) {
	t.Helper()
	return newStakingOperationFromModel(&api.StakingOperation{
		Id:        "staking-operation-id",
		NetworkId: "ethereum-test-network-id",
		AddressId: "test-asset-id",
		Status:    status,
		Transactions: []api.Transaction{
			{
				NetworkId: "ethereum-test-network-id",
				Status:    "pending",
				UnsignedPayload: "7b2274797065223a22307832222c22636861696e4964223a2230783134613334222c226e6f6e63" +
					"65223a22307830222c22746f223a22307834643965346633663464316138623566346637623166" +
					"356235633762386436623262336231623062222c22676173223a22307835323038222c22676173" +
					"5072696365223a6e756c6c2c226d61785072696f72697479466565506572476173223a223078" +
					"3539363832663030222c226d6178466565506572476173223a2230783539363832663030222c22" +
					"76616c7565223a2230783536626337356532643633313030303030222c22696e707574223a22" +
					"3078222c226163636573734c697374223a5b5d2c2276223a22307830222c2272223a2230783022" +
					"2c2273223a22307830222c2279506172697479223a22307830222c2268617368223a2230783664" +
					"633334306534643663323633653363396561396135656438646561346332383966613861363966" +
					"3031653635393462333732386230386138323335333433227d",
			},
		},
		Metadata: &api.StakingOperationMetadata{
			ArrayOfSignedVoluntaryExitMessageMetadata: &[]api.SignedVoluntaryExitMessageMetadata{{
				SignedVoluntaryExit: "dGVzdC1kYXRh",
			}},
		},
	})
}

func mockGetExternalStakingOperation(t *testing.T, stakeAPI *mocks.StakeAPI, statusCode int, soStatus string) {
	t.Helper()
	stakeAPI.On("GetExternalStakingOperation", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
		api.ApiGetExternalStakingOperationRequest{ApiService: stakeAPI},
	).Once()
	stakeAPI.On("GetExternalStakingOperationExecute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
		&api.StakingOperation{
			Id:        "staking-operation-id",
			NetworkId: "ethereum-test-network-id",
			AddressId: "test-asset-id",
			Status:    soStatus,
			Transactions: []api.Transaction{
				{
					NetworkId: "ethereum-test-network-id",
					Status:    "pending",
					UnsignedPayload: "7b2274797065223a22307832222c22636861696e4964223a2230783134613334222c226e6f6e63" +
						"65223a22307830222c22746f223a22307834643965346633663464316138623566346637623166" +
						"356235633762386436623262336231623062222c22676173223a22307835323038222c22676173" +
						"5072696365223a6e756c6c2c226d61785072696f72697479466565506572476173223a223078" +
						"3539363832663030222c226d6178466565506572476173223a2230783539363832663030222c22" +
						"76616c7565223a2230783536626337356532643633313030303030222c22696e707574223a22" +
						"3078222c226163636573734c697374223a5b5d2c2276223a22307830222c2272223a2230783022" +
						"2c2273223a22307830222c2279506172697479223a22307830222c2268617368223a2230783664" +
						"633334306534643663323633653363396561396135656438646561346332383966613861363966" +
						"3031653635393462333732386230386138323335333433227d",
				},
			},
		},
		&http.Response{StatusCode: statusCode},
		nil,
	).Once()
	stakeAPI.On("GetExternalStakingOperation", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
		api.ApiGetExternalStakingOperationRequest{ApiService: stakeAPI},
	).Once()
	stakeAPI.On("GetExternalStakingOperationExecute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
		&api.StakingOperation{
			Id:        "staking-operation-id",
			NetworkId: "ethereum-test-network-id",
			AddressId: "test-asset-id",
			Status:    "complete",
			Transactions: []api.Transaction{
				{
					NetworkId: "ethereum-test-network-id",
					Status:    "complete",
					UnsignedPayload: "7b2274797065223a22307832222c22636861696e4964223a2230783134613334222c226e6f6e63" +
						"65223a22307830222c22746f223a22307834643965346633663464316138623566346637623166" +
						"356235633762386436623262336231623062222c22676173223a22307835323038222c22676173" +
						"5072696365223a6e756c6c2c226d61785072696f72697479466565506572476173223a223078" +
						"3539363832663030222c226d6178466565506572476173223a2230783539363832663030222c22" +
						"76616c7565223a2230783536626337356532643633313030303030222c22696e707574223a22" +
						"3078222c226163636573734c697374223a5b5d2c2276223a22307830222c2272223a2230783022" +
						"2c2273223a22307830222c2279506172697479223a22307830222c2268617368223a2230783664" +
						"633334306534643663323633653363396561396135656438646561346332383966613861363966" +
						"3031653635393462333732386230386138323335333433227d",
				},
			},
		},
		&http.Response{StatusCode: statusCode},
		nil,
	).Once()
}

func (s *StakingOperationSuite) TestSign_AllTransactionsSigned() {
	// Create mock transactions
	signable1 := new(mocks.Signable)
	signable2 := new(mocks.Signable)

	// Set expectations
	signable1.On("IsSigned").Return(true)
	signable2.On("IsSigned").Return(true)

	stakingOp := &StakingOperation{}
	signer, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	s.NoError(err)

	// Set up the staking operation to return these transactions
	stakingOp.transactions = []*Transaction{
		{
			model:    &client.Transaction{},
			signable: signable1,
		},
		{
			model:    &client.Transaction{},
			signable: signable2,
		},
	}

	// Call the Sign method
	err = stakingOp.Sign(signer)

	// Assert no error and that no transactions were signed
	s.NoError(err)
	signable1.AssertNotCalled(s.T(), "Sign", mock.Anything)
	signable2.AssertNotCalled(s.T(), "Sign", mock.Anything)
}

func (s *StakingOperationSuite) TestSign_SomeTransactionsNotSigned() {
	// Create mock transactions
	signable1 := new(mocks.Signable)
	signable2 := new(mocks.Signable)

	signer, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	s.NoError(err)

	// Set expectations
	signable1.On("IsSigned").Return(true)
	signable2.On("IsSigned").Return(false)
	signable2.On("Sign", signer).Return("", nil)

	stakingOp := &StakingOperation{}

	// Set up the staking operation to return these transactions
	stakingOp.transactions = []*Transaction{
		{
			model:    &client.Transaction{},
			signable: signable1,
		},
		{
			model:    &client.Transaction{},
			signable: signable2,
		},
	}

	// Call the Sign method
	err = stakingOp.Sign(signer)

	// Assert no error and that the second transaction was signed
	s.NoError(err)
	signable1.AssertNotCalled(s.T(), "Sign", mock.Anything)
	signable2.AssertCalled(s.T(), "Sign", signer)
}

func (s *StakingOperationSuite) TestSign_SignTransactionFails() {
	// Create mock transactions
	signable1 := new(mocks.Signable)
	signable2 := new(mocks.Signable)

	signer, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	s.NoError(err)

	// Set expectations
	signable1.On("IsSigned").Return(true)
	signable2.On("IsSigned").Return(false)
	signable2.On("Sign", signer).Return("", errors.New("signing failed"))

	stakingOp := &StakingOperation{}

	// Set up the staking operation to return these transactions
	stakingOp.transactions = []*Transaction{
		{
			model:    &client.Transaction{},
			signable: signable1,
		},
		{
			model:    &client.Transaction{},
			signable: signable2,
		},
	}

	// Call the Sign method
	err = stakingOp.Sign(signer)

	// Assert error and that the second transaction was not signed
	s.Error(err)
	s.EqualError(err, "signing failed")
	signable1.AssertNotCalled(s.T(), "Sign", mock.Anything)
	signable2.AssertCalled(s.T(), "Sign", signer)
}
