package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/WilliamMortlMicrosoft/HelloGoService/shared"
	"github.com/WilliamMortlMicrosoft/HelloGoService/telemetry"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// markdown file to display
const markdownFilename string = "README.md"

// telemetry for home
var telemetryHome *telemetry.Telemetry

// HomeHandler is a http handler for "ping" requests
// @Summary home page
// @Description returns the readme file
// @Tags basic services
// @Produce html
// @Success 200 "html"
// @Router / [get]
func HomeHandler() func(w http.ResponseWriter, r *http.Request) {
	if telemetryHome == nil {
		telemetryHome = telemetry.InitializeTelemetryDefault("Home")
	}
	return func(w http.ResponseWriter, r *http.Request) {

		// load readme markdown file
		md, err := ioutil.ReadFile(markdownFilename)
		if err != nil {
			errorTitle := "README Not Found"
			errorMsg := markdownFilename + " was not found!"
			http.Error(w, errorMsg, http.StatusInternalServerError)
			(*telemetryHome).LogError(errorTitle, fmt.Errorf(errorMsg+" IP: ", shared.GetIP(r)))
		}

		// render into html and return
		extensions := parser.CommonExtensions | parser.AutoHeadingIDs
		parser := parser.NewWithExtensions(extensions)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, string(markdown.ToHTML(md, parser, nil)))
		(*telemetryHome).LogInfo("Call", fmt.Sprintf("Home loaded IP: %v", shared.GetIP(r)))
	}
}
