package coinbase_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"log"
	"testing"

	"github.com/coinbase/coinbase-sdk-go/pkg/coinbase"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/stretchr/testify/suite"
)

type SolanaSignableSuite struct {
	suite.Suite

	privateKey ed25519.PrivateKey
	signable   *coinbase.SolanaSignable
}

func TestSolanaSignableSuite(t *testing.T) {
	suite.Run(t, new(SolanaSignableSuite))
}

func (s *SolanaSignableSuite) SetupTest() {
	var err error
	var pubKey ed25519.PublicKey
	pubKey, s.privateKey, err = ed25519.GenerateKey(rand.Reader)
	s.NoError(err)

	// Use a fake blockhash
	fakeBlockhash := solana.MustHashFromBase58("11111111111111111111111111111111")

	// Generate a random recipient public key
	recipientPubKey := solana.NewWallet().PublicKey()

	// Create a transfer instruction to send 1 SOL to the recipient
	transferInstruction := system.NewTransferInstruction(
		1_000_000_000,
		solana.PublicKeyFromBytes(pubKey[:]),
		recipientPubKey,
	).Build()

	// Create a new transaction
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			transferInstruction,
		},
		fakeBlockhash,
	)
	s.NoError(err)

	// Create a SolanaSignable instance
	s.signable = coinbase.NewSolanaSignable(tx)
}

func (s *SolanaSignableSuite) TestSign_InvalidKeyType_RSA() {
	// Generate RSA key
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	s.NoError(err)

	// Test with RSA key
	_, err = s.signable.Sign(rsaKey)
	s.Error(err)
	s.Equal("provided key is not an *ed25519.PrivateKey", err.Error())
}

func (s *SolanaSignableSuite) TestSign_InvalidKeyType_ECDSA() {
	// Generate ECDSA key
	ecdsaKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	s.NoError(err)

	// Test with ECDSA key
	_, err = s.signable.Sign(ecdsaKey)
	s.Error(err)
	s.Equal("provided key is not an *ed25519.PrivateKey", err.Error())
}

func (s *SolanaSignableSuite) TestSign_CorrectPublicKey() {
	s.False(s.signable.IsSigned())

	// Call the Sign method with the new private key
	signedTx, err := s.signable.Sign(&s.privateKey)

	// Assert no error and check the signed transaction
	s.NoError(err)
	s.NotEmpty(signedTx)

	// Verify that the signable is signed
	s.True(s.signable.IsSigned())
}

func (s *SolanaSignableSuite) TestSign_IncorrectPublicKey() {
	// Generate a new ed25519 key pair
	_, newPrivKey, err := ed25519.GenerateKey(rand.Reader)
	s.NoError(err)

	s.False(s.signable.IsSigned())

	// Call the Sign method with the new private key
	signedTx, err := s.signable.Sign(&newPrivKey)

	// Assert no error and check the signed transaction
	s.Error(err)
	s.Empty(signedTx)

	// Verify that the signable is still not signed
	s.False(s.signable.IsSigned())
}

func (s *SolanaSignableSuite) TestRaw_Parsable() {
	s.False(s.signable.IsSigned())

	// Call the Sign method with the new private key
	raw := s.signable.Raw()

	s.NotEmpty(raw)

	rawTx := s.signable.Raw()
	_, ok := rawTx.(*solana.Transaction)
	if !ok {
		log.Fatal("failed to cast raw transaction to solana.Transaction")
	}
}
