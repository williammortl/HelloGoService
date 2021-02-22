package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/WilliamMortlMicrosoft/HelloGoService/db"
	"github.com/WilliamMortlMicrosoft/HelloGoService/shared"
	"github.com/WilliamMortlMicrosoft/HelloGoService/telemetry"
	"github.com/gorilla/mux"
)

// telemetry
var telemetryDBGet *telemetry.Telemetry
var telemetryDBAdd *telemetry.Telemetry

// DBGetHandler is a http handler for db get requests
// @Summary db service - retrieve
// @Description gets a record
// @Tags advanced services
// @Produce json
// @Param id path int true "database id"
// @Success 200 {object} db.Person
// @Failure 400 "error message"
// @Router /Db/{id} [get]
func DBGetHandler() func(w http.ResponseWriter, r *http.Request) {

	// initialize telemetry only on the first call
	if telemetryDBGet == nil {
		telemetryDBGet = telemetry.InitializeTelemetryDefault("DBGet")
	}

	return func(w http.ResponseWriter, r *http.Request) {

		// get {id} from the REST path
		idAsString := mux.Vars(r)["id"]
		idAsInt, err := strconv.Atoi(idAsString)
		if err != nil {
			errorTitle := "Invalid ID"
			errorMsg := fmt.Sprintf("The id: %v is not a valid int!", idAsString)
			http.Error(w, errorMsg, http.StatusBadRequest)
			(*telemetryDBGet).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
			return
		}

		// get the record from the db
		person := db.GetPersonByID(idAsInt)
		if person != nil {
			response, _ := json.Marshal(*person)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(response))
			(*telemetryDBGet).LogInfo("Call",
				fmt.Sprintf("Retrieved id: %v name: %v IP: %v",
					idAsInt,
					(*person).Name,
					shared.GetIP(r)))
			return
		}

		// handle no matching id
		errorTitle := "Missing ID"
		errorMsg := fmt.Sprintf("The id: %v does not exist in the db!", idAsInt)
		http.Error(w, errorMsg, http.StatusBadRequest)
		(*telemetryDBGet).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
	}
}

// DBAddHandler is a http handler for adding / updating records
// @Summary db service - add / update
// @Description adds or updates a record
// @Tags advanced services
// @Accept json
// @Produce json
// @Param id path int true "database id"
// @Param message body db.Person true "data"
// @Success 200 "ok message"
// @Failure 400 "error message"
// @Router /Db/{id} [post]
func DBAddHandler() func(w http.ResponseWriter, r *http.Request) {

	// initialize telemetry only on the first call
	if telemetryDBAdd == nil {
		telemetryDBAdd = telemetry.InitializeTelemetryDefault("DBAdd")
	}

	return func(w http.ResponseWriter, r *http.Request) {

		// read JSON body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errorTitle := "Missing JSON"
			errorMsg := "JSON was not posted posted to this endpoint!"
			http.Error(w, errorMsg, http.StatusBadRequest)
			(*telemetryDBAdd).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
			return
		}

		// get JSON from the body
		var person db.Person
		err = json.Unmarshal(body, &person)
		if (err != nil) || (person.Name == "") {
			errorTitle := "Malformed JSON"
			errorMsg := "Malformed JSON was posted posted to this endpoint!"
			http.Error(w, errorMsg, http.StatusBadRequest)
			(*telemetryDBAdd).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
			return
		}

		// get {id} from the REST path
		idAsString := mux.Vars(r)["id"]
		idAsInt, err := strconv.Atoi(idAsString)
		if err != nil {
			errorTitle := "Invalid ID"
			errorMsg := fmt.Sprintf("The id: %v is not a valid int!", idAsInt)
			http.Error(w, errorMsg, http.StatusBadRequest)
			(*telemetryDBAdd).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
			return
		}

		// add the record to the db
		db.AddPerson(idAsInt, person)
		http.Error(w, "OK", http.StatusOK)
		(*telemetryDBAdd).LogInfo("Call",
			fmt.Sprintf("Successfully updated id: %v name: %v IP: %v",
				idAsInt,
				person.Name,
				shared.GetIP(r)))
	}
}
