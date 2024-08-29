package coinbase

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	api "github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/coinbase/coinbase-sdk-go/pkg/mocks"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestStakingOperation_Wait_Success(t *testing.T) {
	mc := &mockController{
		stakeAPI: mocks.NewStakeAPI(t),
	}
	mockGetExternalStakingOperation(t, mc.stakeAPI, http.StatusOK, "pending")

	c := &Client{
		client: &api.APIClient{
			StakeAPI:  mc.stakeAPI,
			AssetsAPI: mc.assetsAPI,
		},
	}

	so, err := mockStakingOperation(t, "pending")
	require.NoError(t, err, "failed to create staking operation")
	key, err := crypto.GenerateKey()
	require.NoError(t, err, "failed to generate ecdsa key")
	err = so.Sign(key)
	require.NoError(t, err, "failed to sign staking operation")
	signedPayload := so.Transactions()[0].SignedPayload()
	require.NotEmpty(t, signedPayload, "signed payload should not be empty")
	so, err = c.Wait(context.Background(), so)
	assert.NoError(t, err, "staking operation wait should not error")
	assert.Equal(t, "complete", so.Status(), "staking operation status should be complete")
	assert.Equal(t, 1, len(so.Transactions()), "staking operation should have 1 transaction")
	assert.Equal(t, signedPayload, so.Transactions()[0].SignedPayload(), "staking operation signed payload should not have changed")
}

func TestStakingOperation_Wait_Success_CustomOptions(t *testing.T) {
	mc := &mockController{
		stakeAPI: mocks.NewStakeAPI(t),
	}
	mockGetExternalStakingOperation(t, mc.stakeAPI, http.StatusOK, "pending")

	c := &Client{
		client: &api.APIClient{
			StakeAPI:  mc.stakeAPI,
			AssetsAPI: mc.assetsAPI,
		},
	}

	so, err := mockStakingOperation(t, "pending")
	assert.NoError(t, err, "staking operation creation should not error")
	so, err = c.Wait(
		context.Background(),
		so,
		WithWaitIntervalSeconds(1),
		WithWaitTimeoutSeconds(3),
	)
	assert.NoError(t, err, "staking operation wait should not error")
	assert.Equal(t, "complete", so.Status(), "staking operation status should be complete")
	assert.Equal(t, 1, len(so.Transactions()), "staking operation should have 1 transaction")
}

func TestStakingOperation_Wait_Failure(t *testing.T) {
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
		t.Run(name, func(t *testing.T) {
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
			assert.NoError(t, err, "staking operation creation should not error")
			_, err = c.Wait(context.Background(), so)
			assert.Error(t, err, "staking operation wait should error")
		})
	}
}

func TestStakingOperation_GetSignedVoluntaryExitMessages(t *testing.T) {
	stakingOperation, err := mockStakingOperation(t, "pending")
	assert.NoError(t, err, "staking operation creation should not error")

	SignedVoluntaryExitMessages, err := stakingOperation.GetSignedVoluntaryExitMessages()
	assert.NoError(t, err, "get signed voluntary exit messages should not error")

	assert.Equal(t, 1, len(SignedVoluntaryExitMessages), "signed voluntary exit messages should have length 1")
	assert.Equal(t, "test-data", SignedVoluntaryExitMessages[0], "signed voluntary exit messages should match")
}

func mockStakingOperation(t *testing.T, status string) (*StakingOperation, error) {
	t.Helper()
	return newStakingOperationFromModel(&api.StakingOperation{
		Id:        "staking-operation-id",
		NetworkId: "test-network-id",
		AddressId: "test-asset-id",
		Status:    status,
		Transactions: []api.Transaction{
			{
				NetworkId: "test-network-id",
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
			NetworkId: "test-network-id",
			AddressId: "test-asset-id",
			Status:    soStatus,
			Transactions: []api.Transaction{
				{
					NetworkId: "test-network-id",
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
			NetworkId: "test-network-id",
			AddressId: "test-asset-id",
			Status:    "complete",
			Transactions: []api.Transaction{
				{
					NetworkId: "test-network-id",
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
