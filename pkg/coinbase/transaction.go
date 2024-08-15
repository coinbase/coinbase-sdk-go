package coinbase

import "github.com/coinbase/coinbase-sdk-go/gen/client"

// Transaction represents an onchain transaction
type Transaction struct {
	model *client.Transaction
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

func newTransactionFromModel(m *client.Transaction) *Transaction {
	if m == nil {
		return nil
	}
	return &Transaction{
		model: m,
	}
}
