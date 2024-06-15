package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseCliOutputLine(t *testing.T) {
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
			got := ParseCliOutputLine(tt.src)
			assert.Equal(t, tt.want, ParseCliOutputLine(tt.src), "ParseCliOutputLine() = %v, want %v", got, tt.want)
		})
	}
}

func TestParseCliOutput(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		output string
		want   []string
	}{
		{
			name:   "Success - no text = empty slice",
			output: "",
			want:   []string{},
		},
		{
			name: "Success - parsed into lines",
			output: ` :6A\:22\:32\:1D\:24\:E3:Princess_Wifi:153:540 Mbit/s:92:
*:66\:22\:32\:1D\:24\:E3:Nahual:153:540 Mbit/s:92:WPA2
 :60\:22\:32\:1D\:24\:E3:Chupacabra:153:540 Mbit/s:92:WPA2`,
			want: []string{
				` :6A\:22\:32\:1D\:24\:E3:Princess_Wifi:153:540 Mbit/s:92:`,
				`*:66\:22\:32\:1D\:24\:E3:Nahual:153:540 Mbit/s:92:WPA2`,
				` :60\:22\:32\:1D\:24\:E3:Chupacabra:153:540 Mbit/s:92:WPA2`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ParseCliOutput(tt.output), "ParseCliOutput(%v)", tt.output)
		})
	}
}
