package tests

import (
	"context"
	"github.com/artkescha/integration/internal/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient(t *testing.T) {
		tests := []struct {
		name string
		a uint64
		want int
	}{
		{
			name: "added 1",
			a:    1,
			want: 1,
		},
		{
			name: "added 3",
			a:    3,
			want: 3,
		},
	}

	for _, tt := range tests {
		var got int
		client.Run(context.Background(), "127.0.0.1:8081", tt.a, &got)

		assert.Equal(t, tt.want, got)
	}
}
