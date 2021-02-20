package main

import (
	"fmt"
	"net/http"

	"github.com/WilliamMortlMicrosoft/HelloGoService/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// port to listen on
const listenPort int = 8080

// main entry point of the program
func main() {

	router := mux.NewRouter().StrictSlash(true)

	// home page
	router.HandleFunc("/", handlers.HomeHandler())

	// handler for ping
	router.HandleFunc("/ping", handlers.PingHandler())

	// handler for hello
	router.HandleFunc("/hello", handlers.HelloHandler())

	// handler for db get
	router.HandleFunc("/db/{id}", handlers.DBGetHandler()).Methods("GET")

	// handler for db update / add
	router.HandleFunc("/db/{id}", handlers.DBAddHandler()).Methods("POST")

	// handler for math
	router.HandleFunc("/math/{operator}", handlers.MathHandler()).Methods("POST")

	// handler for prometheus
	router.Handle("/metrics", promhttp.Handler())

	// listen and serve
	http.ListenAndServe(fmt.Sprintf(":%v", listenPort), router)
}
