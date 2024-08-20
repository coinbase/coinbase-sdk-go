package coinbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddress_NetworkID(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		{
			name: "should eq test_network",
			a:    NewExternalAddress("test_network", "test_id"),
			want: "test_network",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.a.NetworkID())
		})
	}
}

func TestAddress_ID(t *testing.T) {
	tests := []struct {
		name string
		a    *Address
		want string
	}{
		{
			name: "should eq test_id",
			a:    NewExternalAddress("test_network", "test_id"),
			want: "test_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.a.ID())
		})
	}
}
