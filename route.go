package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// A function to handle the request and pass them to its corresponding router
func handleRequests() {
	// Creates a new instance of a mux router
	muxRouter := mux.NewRouter().StrictSlash(true)
	// Replace http.HandleFunc with myRouter.HandleFunc
	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/device", createDevice).Methods("POST")
	muxRouter.HandleFunc("/devices", returnAllDevices)
	muxRouter.HandleFunc("/devices/{id}", returnOneDevice)
	muxRouter.HandleFunc("/delete/{id}", deleteDevice).Methods("DELETE")
	muxRouter.HandleFunc("/update/{id}", updateDevice).Methods("PUT")
	// Passing in the muxRouter as a second argument instead of a nil

	// PostgreSQL Endpoints
	muxRouter.HandleFunc("/psql/device", createDevicePSQL).Methods("POST")
	muxRouter.HandleFunc("/psql/devices", returnAllDevicesPSQL)
	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}
