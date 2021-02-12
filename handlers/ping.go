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

// PingHandler is a http handler for "ping" requests
func PingHandler() func(w http.ResponseWriter, r *http.Request) {
	telemetryPing := telemetry.InitializeTelemetryDefault("Ping")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(pingResponse{
			Message: "Pong!",
			Time:    time.Now().Format(time.RFC850),
		})
		fmt.Fprintf(w, string(response))
		(*telemetryPing).LogInfo("Call", fmt.Sprintf("Ping from: %v", shared.GetIP(r)))
	}
}
