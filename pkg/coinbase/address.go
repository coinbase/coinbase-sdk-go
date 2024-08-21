package coinbase

// Address represents a blockchain address.
type Address struct {
	networkID string
	id        string
}

// NewExternalAddress creates a new external address.
func NewExternalAddress(networkID string, ID string) *Address {
	return &Address{
		networkID: networkID,
		id:        ID,
	}
}

// NetworkID returns the address network id
func (a *Address) NetworkID() string {
	return a.networkID
}

// ID returns the address id
func (a *Address) ID() string {
	return a.id
}
