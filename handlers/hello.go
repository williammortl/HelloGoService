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

// HelloHandler is a http handler for "hello" requests
func HelloHandler() func(w http.ResponseWriter, r *http.Request) {
	telemetryHello := telemetry.InitializeTelemetryDefault("Hello")
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" {
			w.Header().Set("Content-Type", "application/json")
			response, _ := json.Marshal(helloResponse{
				Message: "Hello World!",
				Name:    name,
			})
			fmt.Fprintf(w, string(response))
			(*telemetryHello).LogInfo("Call", fmt.Sprintf("Hello from: %v - %v", name, shared.GetIP(r)))
			return
		}

		// handle no query string
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
			<html>
			<title>Missing QueryString</title>
			<body>
				When calling the service, you forgot to send your name in the query string! <br />
				Example: <br />
				&emsp;http://{server name}:{server port}/hello<u><b>?name=YourNameHere</b></u>
			</body>
			</html>`)
		(*telemetryHello).LogError("No params", fmt.Errorf("Hello called, but no name sent! %v", shared.GetIP(r)))
	}
}
