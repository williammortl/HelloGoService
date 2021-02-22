package main

import (
	"fmt"
	"net/http"

	_ "github.com/WilliamMortlMicrosoft/HelloGoService/docs"
	"github.com/WilliamMortlMicrosoft/HelloGoService/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

// port to listen on
const listenPort int = 8080

// @title Hello GO Service Example API
// @version 1.0
// @description This is a suite of simple service API's.
// @termsOfService http://swagger.io/terms/
// @contact.name William Mortl
// @contact.url https://github.com/williammortl/HelloGoService
// @contact.email will@{insert my full name here}.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// main entry point of the program
func main() {

	router := mux.NewRouter().StrictSlash(true)

	// home page
	router.HandleFunc("/", handlers.HomeHandler())

	// handler for ping
	router.HandleFunc("/Ping", handlers.PingHandler())

	// handler for hello
	router.HandleFunc("/Hello", handlers.HelloHandler())

	// handler for db get
	router.HandleFunc("/Db/{id}", handlers.DBGetHandler()).Methods("GET")

	// handler for db update / add
	router.HandleFunc("/Db/{id}", handlers.DBAddHandler()).Methods("POST")

	// handler for math
	router.HandleFunc("/Math/{operator}", handlers.MathHandler()).Methods("POST")

	// handler for prometheus
	router.Handle("/metrics", promhttp.Handler())

	// handler for swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%v/swagger/doc.json", listenPort)),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	// listen and serve
	http.ListenAndServe(fmt.Sprintf(":%v", listenPort), router)
}
