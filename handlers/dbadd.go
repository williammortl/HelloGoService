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

// telemetry for db get
var telemetryDBAdd *telemetry.Telemetry

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
// @Router /db/{id} [post]
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
