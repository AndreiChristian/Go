package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8090"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_, _ = fmt.Printf("Starting the application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
