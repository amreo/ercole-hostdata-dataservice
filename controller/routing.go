// Copyright (c) 2019 Sorint.lab S.p.A.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"github.com/amreo/ercole-hostdata-dataservice/model"
	"github.com/amreo/ercole-hostdata-dataservice/service"
	null "gopkg.in/guregu/null.v3"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})
	router.HandleFunc("/hosts", updateHost).Methods("POST")
}

func updateHost(w http.ResponseWriter, r *http.Request) {
	var hostData model.HostData

	if err := json.NewDecoder(r.Body).Decode(&hostData); err != nil {
		writeResponsError(w, http.StatusUnprocessableEntity, err)
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

//ErrorResponseFE : struct describing errors in response
type ErrorResponseFE struct {
	Error            null.String
	ErrorDescription null.String
}

func writeResponsError(w http.ResponseWriter, statusCode int, err error) {
	writeJSONResponse(w, statusCode, ErrorResponseFE{
		Error:            null.StringFrom(http.StatusText(statusCode)),
		ErrorDescription: null.StringFrom(err.Error()),
	})
}

// writeJSONResponse write the statuscode and the response to w
func writeJSONResponse(w http.ResponseWriter, statusCode int, resp interface{}) {
	//Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// internalServerError log the caller code position and write to w 500 Internal Server Error
func internalServerError() {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Println(file, ":", line)
	}
}

// internalServerErrorWithError log the error with the caller code position and write to w 500 Internal Server Error
func internalServerErrorWithError(err error) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		log.Println(file, ":", line, "err:", err)
	} else {
		log.Println("??? err:", err)
	}
}
