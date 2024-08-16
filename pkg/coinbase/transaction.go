package coinbase

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/ethereum/go-ethereum/core/types"
)

// Transaction represents an onchain transaction
type Transaction struct {
	model *client.Transaction
	raw   *types.Transaction
}

// UnsignedPayload returns the unsigned payload of the transaction
func (t *Transaction) UnsignedPayload() string {
	return t.model.UnsignedPayload
}

// SignedPayload returns the signed payload of the transaction
func (t *Transaction) SignedPayload() string {
	if t.model.SignedPayload == nil {
		return ""
	}

	return *t.model.SignedPayload
}

// TransactionHash returns the hash of the transaction
func (t *Transaction) TransactionHash() string {
	if t.model.TransactionHash == nil {
		return ""
	}

	return *t.model.TransactionHash
}

// Status returns the status of the Transaction
func (t *Transaction) Status() string {
	return t.model.Status
}

// FromAddressID returns the from address for the transaction
func (t *Transaction) FromAddressID() string {
	return t.model.FromAddressId
}

// Raw returns the raw transaction in types.Transaction format
func (t *Transaction) Raw() *types.Transaction {
	return t.raw
}

// IsSigned returns true if the transaction is signed
func (t *Transaction) IsSigned() bool {
	v, r, s := t.Raw().RawSignatureValues()
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

// Sign will sign the transaction using the supplied key
func (t *Transaction) Sign(k *ecdsa.PrivateKey) error {
	if t.IsSigned() {
		return nil
	}

	signer := types.NewEIP155Signer(t.Raw().ChainId())
	signedTx, err := types.SignTx(t.Raw(), signer, k)
	if err != nil {
		return err
	}

	t.raw = signedTx
	return nil
}

func newTransactionFromModel(m *client.Transaction) (*Transaction, error) {
	if m == nil {
		return nil, fmt.Errorf("transaction model cannot be nil")
	}

	rawHex, err := hex.DecodeString(m.UnsignedPayload)
	if err != nil {
		return nil, err
	}

	t := &types.Transaction{}
	err = t.UnmarshalBinary(rawHex)
	if err != nil {
		return nil, err
	}

	return &Transaction{
		model: m,
		raw:   t,
	}, nil
}
