package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseCliOutput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		src  string
		want []string
	}{
		{
			name: "Success - valid fields",
			src:  `*:66\:22\:32\:1D\:24\:E3:Nahual:153:540 Mbit/s:74:WPA2`,
			want: []string{
				"*",
				"66:22:32:1D:24:E3",
				"Nahual",
				"153",
				"540 Mbit/s",
				"74",
				"WPA2",
			},
		},
		{
			name: "Success - empty string",
			src:  ``,
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseCliOutput(tt.src)
			assert.Equal(t, tt.want, parseCliOutput(tt.src), "parseCliOutput() = %v, want %v", got, tt.want)
		})
	}
}
