package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/WilliamMortlMicrosoft/HelloGoService/shared"
	"github.com/WilliamMortlMicrosoft/HelloGoService/telemetry"
)

// the response for ping
type pingResponse struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

// telemetry for ping
var telemetryPing *telemetry.Telemetry

// PingHandler is a http handler for "ping" requests
func PingHandler() func(w http.ResponseWriter, r *http.Request) {
	if telemetryPing == nil {
		telemetryPing = telemetry.InitializeTelemetryDefault("Ping")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		response, _ := json.Marshal(pingResponse{
			Message: "Pong!",
			Time:    time.Now().Format(time.RFC850),
		})
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(response))
		(*telemetryPing).LogInfo("Call", fmt.Sprintf("Ping! IP: %v", shared.GetIP(r)))
	}
}
