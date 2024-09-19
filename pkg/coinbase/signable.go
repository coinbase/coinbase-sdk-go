package coinbase

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gagliardetto/solana-go"
)

// This struct implements the Signable interface
type SolanaSignable struct {
	raw *solana.Transaction
}

func (s *SolanaSignable) Sign(k crypto.Signer) (string, error) {
	if s.IsSigned() {
		bytes, err := s.raw.MarshalBinary()
		if err != nil {
			return "", err
		}

		return base58.Encode(bytes), nil
	}

	ed25519Key, ok := k.(*ed25519.PrivateKey)
	if !ok {
		return "", fmt.Errorf("provided key is not an *ed25519.PrivateKey")
	}

	solanaPrivateKey := solana.PrivateKey(*ed25519Key)

	// clear signatures: https://github.com/gagliardetto/solana-go/issues/186
	s.raw.Signatures = nil

	// Sign the transaction
	if _, err := s.raw.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if solanaPrivateKey.PublicKey().Equals(key) {
			return &solanaPrivateKey
		}
		return nil
	}); err != nil {
		return "", fmt.Errorf("error signing transaction: %w", err)
	}

	// Marshal the signed transaction to binary
	marshaledTx, err := s.raw.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("error marshaling transaction: %w", err)
	}

	// Encode the binary transaction using base58
	base58EncodedSignedTx := base58.Encode(marshaledTx)

	return base58EncodedSignedTx, nil
}

func (s *SolanaSignable) IsSigned() bool {
	if s.raw.Signatures == nil {
		return false
	}

	if len(s.raw.Signatures) == 0 {
		return false
	}

	// Check for the placeholder signature.
	if len(s.raw.Signatures) == 1 && s.raw.Signatures[0].Equals(solana.Signature{}) {
		return false
	}

	return true
}

func (s *SolanaSignable) Raw() interface{} {
	return s.raw
}

type EthereumSignable struct {
	raw *types.Transaction
}

func (e *EthereumSignable) Sign(k crypto.Signer) (string, error) {
	if e.IsSigned() {
		bytes, err := e.raw.MarshalBinary()
		if err != nil {
			return "", err
		}

		return hex.EncodeToString(bytes), nil
	}

	ecdsaKey, ok := k.(*ecdsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("provided key is not an *ecdsa.PrivateKey")
	}

	signer := types.LatestSignerForChainID(e.raw.ChainId())
	signedTx, err := types.SignTx(e.raw, signer, ecdsaKey)
	if err != nil {
		return "", err
	}

	bytes, err := signedTx.MarshalBinary()
	if err != nil {
		return "", err
	}

	e.raw = signedTx
	return hex.EncodeToString(bytes), nil
}

func (e *EthereumSignable) IsSigned() bool {
	v, r, s := e.raw.RawSignatureValues()
	if v != nil && v.Cmp(big.NewInt(0)) != 0 {
		return true
	}
	if r != nil && r.Cmp(big.NewInt(0)) != 0 {
		return true
	}
	if s != nil && s.Cmp(big.NewInt(0)) != 0 {
		return true
	}
	return false
}

func (e *EthereumSignable) Raw() interface{} {
	return e.raw
}
