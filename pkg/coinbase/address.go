package coinbase

type Address struct {
	NetworkID string `json:"network_id"`
	ID        string `json:"id"`
}

func NewAddress(networkID string, ID string) *Address {
	return &Address{
		NetworkID: networkID,
		ID:        ID,
	}
}
