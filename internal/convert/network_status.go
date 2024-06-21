package convert

import (
	"github.com/danesparza/iot-wifi-setup/internal/network/model"
	"github.com/rs/zerolog/log"
)

const (
	NS_STATE = iota
	NS_CONNECTIVITY
	NS_WIFI
)

// ConvertFieldsToNetworkStatus converts a slice of fields to a model.NetworkStatus
// The returned ActiveConnections will be empty
func ConvertFieldsToNetworkStatus(fields []string) model.NetworkStatus {
	retval := model.NetworkStatus{
		ActiveConnections: make([]model.Connection, 0),
	}

	//	If we don't have the right number of fields, just get out
	if len(fields) != 3 {
		log.Warn().Strs("fields", fields).Msg("Wrong number of fields passed.  Returning empty NetworkStatus")
		return retval
	}

	retval.State = fields[NS_STATE]
	retval.Wifi = fields[NS_WIFI]
	retval.Connectivity = fields[NS_CONNECTIVITY]

	return retval
}
