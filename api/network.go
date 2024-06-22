package api

import (
	"encoding/json"
	"fmt"
	"github.com/danesparza/iot-wifi-setup/internal/network/model"
	"net/http"
)

// GetNetworkStatus godoc
// @Summary Gets current network status
// @Description Gets current network status
// @Tags network
// @Accept  json
// @Produce  json
// @Success 200 {object} api.SystemResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /status [get]
func (service Service) GetNetworkStatus(rw http.ResponseWriter, req *http.Request) {
	//	Call the network manager to get the status
	status, err := service.NM.NetworkStatus(req.Context())
	if err != nil {
		err = fmt.Errorf("error getting network status: %v", err)
		sendErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	//	Show network status
	response := SystemResponse{
		Message: "Network status",
		Data:    status,
	}

	//	Serialize to JSON & return the response:
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(response)
}

// ListAccessPoints godoc
// @Summary List all nearby wifi access points
// @Description List all nearby wifi access points
// @Tags network
// @Accept  json
// @Produce  json
// @Success 200 {object} api.SystemResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /aps [get]
func (service Service) ListAccessPoints(rw http.ResponseWriter, req *http.Request) {
	//	Call the network manager to list the APs
	aps, err := service.NM.ListAccessPoints(req.Context())
	if err != nil {
		err = fmt.Errorf("error getting access points: %v", err)
		sendErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	//	Show wifi access points
	response := SystemResponse{
		Message: "Nearby wifi access points",
		Data:    aps,
	}

	//	Serialize to JSON & return the response:
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(response)
}

// StartAPMode godoc
// @Summary Start Access Point mode
// @Description Start Access Point mode
// @Tags network
// @Accept  json
// @Produce  json
// @Param request body model.APModeRequest true "The SSID (required) and passphrase (optional) to use with the AP"
// @Success 200 {object} api.SystemResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /aps [put]
func (service Service) StartAPMode(rw http.ResponseWriter, req *http.Request) {

	//	Parse the body to get the request info
	request := model.APModeRequest{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		err = fmt.Errorf("problem decoding start AP mode request: %w", err)
		sendErrorResponse(rw, err, http.StatusBadRequest)
		return
	}

	//	Call the network manager to list the APs
	err = service.NM.StartAPMode(req.Context(), request.SSID, request.Passphrase)
	if err != nil {
		err = fmt.Errorf("error starting ap mode: %w", err)
		sendErrorResponse(rw, err, http.StatusInternalServerError)
		return
	}

	//	Show wifi access points
	response := SystemResponse{
		Message: "AP mode started",
		Data:    request,
	}

	//	Serialize to JSON & return the response:
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(response)
}
