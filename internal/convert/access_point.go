package convert

import (
	"github.com/danesparza/iot-wifi-setup/internal/network/model"
	"github.com/rs/zerolog/log"
	"strconv"
)

const (
	AP_INUSE = iota
	AP_BSSID
	AP_SSID
	AP_CHANNEL
	AP_RATE
	AP_SIGNAL
	AP_SECURITY
)

// ConvertFieldsToAccessPoint converts a slice of fields to a model.AccessPoint
func ConvertFieldsToAccessPoint(fields []string) model.AccessPoint {
	retval := model.AccessPoint{}

	//	If we don't have the right number of fields, just get out
	if len(fields) != 7 {
		log.Warn().Strs("fields", fields).Msg("Wrong number of fields passed.  Returning empty AccessPoint")
		return retval
	}

	if fields[AP_INUSE] == "*" {
		retval.InUse = true
	}

	retval.BSSID = fields[AP_BSSID]
	retval.SSID = fields[AP_SSID]

	channel, err := strconv.Atoi(fields[AP_CHANNEL])
	if err == nil {
		retval.Channel = channel
	}

	retval.Rate = fields[AP_RATE]

	signal, err := strconv.Atoi(fields[AP_SIGNAL])
	if err == nil {
		retval.Signal = signal
	}

	retval.Security = fields[AP_SECURITY]

	return retval
}
