package coinbase_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"math/big"
	"testing"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
)

type EthereumSignableSuite struct {
	suite.Suite
	privKey  *ecdsa.PrivateKey
	pubKey   ecdsa.PublicKey
	signable *coinbase.EthereumSignable
}

func TestEthereumSignableSuite(t *testing.T) {
	suite.Run(t, new(EthereumSignableSuite))
}

func (s *EthereumSignableSuite) SetupTest() {
	// Generate a new ECDSA key pair
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	s.NoError(err)

	s.privKey = privKey
	s.pubKey = privKey.PublicKey

	// Create a fake transaction (for example purposes)
	tx := types.NewTransaction(
		0,                                // Nonce
		crypto.PubkeyToAddress(s.pubKey), // To address
		big.NewInt(1000000000000000000),  // Value (1 ETH)
		21000,                            // Gas limit
		big.NewInt(1),                    // Gas price
		nil,                              // Data
	)

	// Create an EthereumSignable instance
	s.signable = coinbase.NewEthereumSignable(tx)
}

func (s *EthereumSignableSuite) TestSign_InvalidKeyType_RSA() {
	// Generate RSA key
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	s.NoError(err)

	// Test with RSA key
	_, err = s.signable.Sign(rsaKey)
	s.Error(err)
	s.Equal("provided key is not an *ecdsa.PrivateKey", err.Error())
}

func (s *EthereumSignableSuite) TestSign_InvalidKeyType_ED25519() {
	// Generate ECDSA key with a different curve
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)

	s.NoError(err)

	// Test with ECDSA key
	_, err = s.signable.Sign(privateKey)
	s.Error(err)
	s.Equal("provided key is not an *ecdsa.PrivateKey", err.Error())
}

func (s *EthereumSignableSuite) TestSign_SuccessfulSigning() {
	// Call the Sign method
	signedTx, err := s.signable.Sign(s.privKey)

	// Assert no error and check the signed transaction
	s.NoError(err)
	s.NotEmpty(signedTx)

	// Verify that the signable is signed
	s.True(s.signable.IsSigned())
}

func (s *EthereumSignableSuite) TestSign_NewPublicKey() {
	// Generate a new ECDSA key pair
	newPrivKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	s.NoError(err)

	// Call the Sign method with the new private key
	signedTx, err := s.signable.Sign(newPrivKey)

	// Assert no error and check the signed transaction
	s.NoError(err)
	s.NotEmpty(signedTx)

	// Verify that the signable is signed
	s.True(s.signable.IsSigned())
}
