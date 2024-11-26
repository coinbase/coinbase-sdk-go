package coinbase

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/coinbase/coinbase-sdk-go/gen/client"
	"github.com/ethereum/go-ethereum/core/types"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

// Transaction represents an onchain transaction
type Transaction struct {
	model    *client.Transaction
	signable Signable
}

type Signable interface {
	Sign(crypto.Signer) (string, error)
	IsSigned() bool
	Raw() interface{}
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

// TransactionLink returns the link to the transaction on the blockchain explorer.
func (t *Transaction) TransactionLink() string {
	if t.model.TransactionLink == nil {
		return ""
	}

	return *t.model.TransactionLink
}

// Status returns the status of the Transaction
func (t *Transaction) Status() string {
	return t.model.Status
}

// FromAddressID returns the from address for the transaction
func (t *Transaction) FromAddressID() string {
	return t.model.FromAddressId
}

// ToAddressID returns the to address for the transaction if it exists
func (t *Transaction) ToAddressID() string {
	if t.model.ToAddressId == nil {
		return ""
	}

	return *t.model.ToAddressId
}

// Raw returns the raw transaction in the underlying blockchain's format
func (t *Transaction) Raw() interface{} {
	return t.signable.Raw()
}

// IsSigned returns true if the transaction is signed
func (t *Transaction) IsSigned() bool {
	return t.signable.IsSigned()
}

func (t *Transaction) Content() *client.TransactionContent {
	return t.model.Content
}

// String returns a string representation of the transaction
func (t *Transaction) String() string {
	return fmt.Sprintf("Transaction{ TransactionHash: %s TransactionLink: %s}", t.TransactionHash(), t.TransactionLink())
}

// Sign will sign the transaction using the supplied key
func (t *Transaction) Sign(k crypto.Signer) error {
	if t.IsSigned() {
		return nil
	}

	signedTx, err := t.signable.Sign(k)
	if err != nil {
		return err
	}

	t.model.SignedPayload = &signedTx

	return nil
}

func newTransactionFromModel(m *client.Transaction) (*Transaction, error) {
	if m == nil {
		return nil, fmt.Errorf("transaction model cannot be nil")
	}

	resp := &Transaction{}

	if strings.HasPrefix(m.NetworkId, "ethereum") {
		rawHex, err := hex.DecodeString(m.UnsignedPayload)
		if err != nil {
			return nil, err
		}

		t := &types.Transaction{}
		if err := t.UnmarshalJSON(rawHex); err != nil {
			return nil, err
		}

		resp.signable = &EthereumSignable{raw: t}
	} else if strings.HasPrefix(m.NetworkId, "solana") {
		data := base58.Decode(m.UnsignedPayload)

		tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(data))
		if err != nil {
			return nil, err
		}

		resp.signable = &SolanaSignable{raw: tx}
	} else {
		return nil, fmt.Errorf("unsupported network id: %s", m.NetworkId)
	}

	resp.model = m

	return resp, nil
}
