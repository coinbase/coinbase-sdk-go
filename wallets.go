package coinbase

import (
	"context"

	"github.com/coinbase/coinbase-sdk-go/gen/client"
)

type Wallet struct {
}

func (c *Client) ListWallets(ctx context.Context) ([]client.Wallet, error) {
	resp, httpRes, err := c.client.WalletsAPI.ListWallets(ctx).Execute()
	if err != nil {
		return nil, err
	}
	if httpRes.StatusCode != 200 {
		return nil, err
	}
	return resp.Data, nil
}
