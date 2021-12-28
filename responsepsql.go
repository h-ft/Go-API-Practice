package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JsonResponse struct {
	Type    string   `json:"type"`
	Data    []Device `json:"data"`
	Message string   `json:"message"`
}

// Create a device

// response and request handlers
func createDevicePSQL(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint: create device psql")

	// Limiting the length of incoming requests to 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	// Passing the body of our POST request into a variable, same as createDevice
	reqBody, _ := ioutil.ReadAll(r.Body)

	var newDevice Device
	json.Unmarshal(reqBody, &newDevice)

	id := newDevice.ID
	deviceName := newDevice.Name
	deviceBrand := newDevice.Brand
	deviceYear := newDevice.Year

	var response = JsonResponse{}

	if id == "" || deviceName == "" || deviceBrand == "" || deviceYear == "" {
		response = JsonResponse{Type: "error", Message: "A field is empty."}
	} else {
		db := setupDB()
		fmt.Println("Inserting new device with ID: " + id + ", name: " + deviceName + ", brand: " + deviceBrand + ", and year: " + deviceYear)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO devices(id, deviceName, deviceBrand, deviceYear) VALUES($1, $2, $3, $4) returning id;", id, deviceName, deviceBrand, deviceYear).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The device has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func returnAllDevicesPSQL(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	// Querying and check for errors
	rows, err := db.Query("SELECT * FROM devices")
	checkErr(err)

	// var response []JsonResponse
	var devices []Device

	// Foreach device
	for rows.Next() {
		var id string
		var deviceName string
		var deviceBrand string
		var deviceYear string

		err = rows.Scan(&id, &deviceName, &deviceBrand, &deviceYear)

		// check errors
		checkErr(err)

		devices = append(devices, Device{ID: id, Name: deviceName, Brand: deviceBrand, Year: deviceYear})
	}

	var response = JsonResponse{Type: "success", Data: devices}

	json.NewEncoder(w).Encode(response)

	fmt.Println("Endpoint: return all devices psql")
}

// Function to update a device in the PostgreSQL Database
func updateDevicePSQL(w http.ResponseWriter, r *http.Request) {
	// db := setupDB()
	fmt.Println("Endpoint: update device psql")

	// Limiting the length of incoming requests to 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	// Passing the body of our POST request into a variable, same as createDevice
	reqBody, _ := ioutil.ReadAll(r.Body)

	var newDevice Device
	json.Unmarshal(reqBody, &newDevice)
}
