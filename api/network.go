package api

import (
	"encoding/json"
	"fmt"
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
