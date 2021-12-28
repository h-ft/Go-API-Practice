package main

import (
	"fmt"
)

func main() {
	go fmt.Println("Server started")
	handleRequests()
}
