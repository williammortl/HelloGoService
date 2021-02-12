package main

import (
	"fmt"
	"net/http"

	"github.com/WilliamMortlMicrosoft/HelloGoService/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// port to listen on
const listenPort int = 8080

// main entry point of the program
func main() {

	// handler for ping
	http.HandleFunc("/ping", handlers.PingHandler())

	// handler for hello
	http.HandleFunc("/hello", handlers.HelloHandler())

	// handler for prometheus
	http.Handle("/metrics", promhttp.Handler())

	// listen and serve
	http.ListenAndServe(fmt.Sprintf(":%v", listenPort), nil)
}
