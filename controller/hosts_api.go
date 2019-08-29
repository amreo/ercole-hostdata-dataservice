package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/amreo/ercole-hostdata-dataservice/model"
	"github.com/amreo/ercole-hostdata-dataservice/service"
	"github.com/goji/httpauth"
)

func (this *HostDataController) AuthenticateMiddleware() func(http.Handler) http.Handler {
	return httpauth.SimpleBasicAuth(this.Config.HttpServer.AgentUsername, this.Config.HttpServer.AgentPassword)
}

// UpdateHostInfo update the informations about a host using the HostData in the request
func (this *HostDataController) UpdateHostInfo(w http.ResponseWriter, r *http.Request) {
	var err error

	var hostData model.HostData
	if err = json.NewDecoder(r.Body).Decode(&hostData); err != nil {
		WriteResponsError(w, http.StatusUnprocessableEntity, err)
		return
	}
	log.Println(hostData)

	err := service.SaveHostData(hostData)
	if err != nil {
		writeResponsError(w, http.StatusUnprocessableEntity, err)
		internalServerErrorWithError(err)
		return
	}
}
