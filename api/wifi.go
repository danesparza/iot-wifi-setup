package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

	//	If we've gotten this far, indicate a successful upload
	response := SystemResponse{
		Message: "Nearby wifi access points",
		Data:    aps,
	}

	//	Serialize to JSON & return the response:
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(rw).Encode(response)
}
