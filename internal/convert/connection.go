package convert

import (
	"github.com/danesparza/iot-wifi-setup/internal/network/model"
	"github.com/rs/zerolog/log"
)

const (
	CON_NAME = iota
	CON_UUID
	CON_TYPE
	CON_DEVICE
)

// ConvertFieldsToActiveConnections converts a slice of fields to a model.Connection
func ConvertFieldsToActiveConnections(fields []string) model.Connection {
	retval := model.Connection{}

	//	If we don't have the right number of fields, just get out
	if len(fields) != 4 {
		log.Warn().Strs("fields", fields).Msg("Wrong number of fields passed.  Returning empty Connection")
		return retval
	}

	retval.Device = fields[CON_DEVICE]
	retval.Type = fields[CON_TYPE]
	retval.Name = fields[CON_NAME]
	retval.UUID = fields[CON_UUID]

	return retval
}
