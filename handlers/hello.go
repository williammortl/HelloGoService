package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WilliamMortlMicrosoft/HelloGoService/shared"
	"github.com/WilliamMortlMicrosoft/HelloGoService/telemetry"
)

// the response for hello
type helloResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

// telemetry for hello
var telemetryHello *telemetry.Telemetry

// HelloHandler is a http handler for "hello" requests
// @Summary gets a "hello world" message
// @Description get a "hello world" message
// @Tags basic services
// @Produce json
// @Param name query string true "User Name"
// @Success 200 {object} helloResponse
// @Failure 400 "error message"
// @Router /hello [get]
func HelloHandler() func(w http.ResponseWriter, r *http.Request) {
	if telemetryHello == nil {
		telemetryHello = telemetry.InitializeTelemetryDefault("Hello")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" {
			response, _ := json.Marshal(helloResponse{
				Message: "Hello World!",
				Name:    name,
			})
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(response))
			(*telemetryHello).LogInfo("Call", fmt.Sprintf("Hello from: %v IP: %v", name, shared.GetIP(r)))
			return
		}

		// handle no query string
		errorTitle := "Missing QueryString"
		errorMsg := "When calling the service, you forgot to send your \"name\" in the query string!"
		http.Error(w, errorMsg, http.StatusBadRequest)
		(*telemetryHello).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
	}
}
