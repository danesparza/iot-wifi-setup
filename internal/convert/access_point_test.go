package convert

import (
	"github.com/danesparza/iot-wifi-setup/internal/network/model"
	"reflect"
	"testing"
)

func TestConvertFieldsToAccessPoint(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		fields []string
		want   model.AccessPoint
	}{
		{
			name: "Success - Too few fields = blank model",
			fields: []string{
				"*",
				"66:22:32:1D:24:E3",
				"NotEnough",
			},
			want: model.AccessPoint{},
		},
		{
			name: "Success - Too many fields = blank model",
			fields: []string{
				"*",
				"66:22:32:1D:24:E3",
				"Nahual",
				"153",
				"540 Mbit/s",
				"74",
				"WPA2",
				"TooMany",
			},
			want: model.AccessPoint{},
		},
		{
			name: "Success - Valid fields = valid model",
			fields: []string{
				"*",
				"66:22:32:1D:24:E3",
				"Nahual",
				"153",
				"540 Mbit/s",
				"74",
				"WPA2",
			},
			want: model.AccessPoint{
				InUse:    true,
				BSSID:    "66:22:32:1D:24:E3",
				SSID:     "Nahual",
				Channel:  153,
				Rate:     "540 Mbit/s",
				Signal:   74,
				Security: "WPA2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertFieldsToAccessPoint(tt.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertFieldsToAccessPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
