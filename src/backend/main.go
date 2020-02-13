/*
	initialise the API
*/

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ddo/go-vue-handler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gitlab.com/bobymcbobs/go-mux-vuejs-buefy-typescript-cloud-template/src/backend/common"
	"gitlab.com/bobymcbobs/go-mux-vuejs-buefy-typescript-cloud-template/src/backend/routes"
)

func handleWebserver() {
	// bring up the API
	port := common.GetAppPort()
	router := mux.NewRouter().StrictSlash(true)
	apiEndpointPrefix := "/api"

	for _, endpoint := range routes.GetEndpoints(apiEndpointPrefix) {
		router.HandleFunc(endpoint.EndpointPath, endpoint.HandlerFunc).Methods(endpoint.HttpMethod, http.MethodOptions)
	}

	router.HandleFunc(apiEndpointPrefix+"/{.*}", routes.APIUnknownEndpoint)
	router.HandleFunc(apiEndpointPrefix, routes.APIroot)
	router.PathPrefix("/").Handler(vue.Handler(common.GetAppDistFolder())).Methods("GET")

	router.Use(common.Logging)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Handler:      c.Handler(router),
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on", port)
	log.Fatal(srv.ListenAndServe())
}

func main() {
	// initialise the app
	handleWebserver()
}
