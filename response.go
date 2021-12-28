package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Homepage default function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok!")
	fmt.Println("Endpoint: homePage")
}

// A function to return all devices from the database
func returnAllDevices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: returnAlldevices")
	json.NewEncoder(w).Encode(Devices)
}

// A function to return only one device from the database, queried by the ID
func returnOneDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the device.ID equals the key we pass in
	// return the article encoded as JSON
	for _, device := range Devices {
		if device.ID == key {
			json.NewEncoder(w).Encode(device)
		}
	}
}

// A function to insert a device to the database
func createDevice(w http.ResponseWriter, r *http.Request) {
	// Limiting the length of incoming requests to 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	// Passing the body of our POST request into a variable
	reqBody, _ := ioutil.ReadAll(r.Body)

	var device Device
	json.Unmarshal(reqBody, &device)

	// Appending our newly created device into the global devices array
	Devices = append(Devices, device)

	json.NewEncoder(w).Encode(device)
}

// A function to delete a device from the database by ID
func deleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the device.ID equals the key we pass in
	// return the article encoded as JSON
	for index, device := range Devices {
		if device.ID == key {
			Devices = append(Devices[:index], Devices[index+1:]...)
		}
	}
}

// A function to update a device from the database
func updateDevice(w http.ResponseWriter, r *http.Request) {
	// Limiting the length of incoming requests to 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	vars := mux.Vars(r)
	key := vars["id"]

	// Passing the body of our POST request into a variable, same as createDevice
	reqBody, _ := ioutil.ReadAll(r.Body)

	var updateDevice Device
	json.Unmarshal(reqBody, &updateDevice)

	// Instead of appending, this time the array will be traversed and the matching entry will be updated
	for index, device := range Devices {
		if device.ID == key {
			device.ID = updateDevice.ID
			device.Brand = updateDevice.Brand
			device.Name = updateDevice.Name
			device.Year = updateDevice.Year
			Devices = append(Devices[:index], device)
			json.NewEncoder(w).Encode(device)
		}
	}
}
