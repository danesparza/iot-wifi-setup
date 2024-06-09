package api

import "net/http"

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

}
