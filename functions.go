package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	// call func with port. Return port as write only and error as err
	_, err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	// do things
	fmt.Println("Server started on port", port)
	// return port and error
	return port, errors.New("ERROR: ZOMG")
}
