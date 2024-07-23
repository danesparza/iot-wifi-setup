package network

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "Generate random string",
			n:    4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateRandomString(tt.n)
			t.Logf("result: %s", result)
		})
	}
}
