package main

import (
	"first_app/pkg/handlers"
	"fmt"
	"net/http"
)

var portNumber = ":8090"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_, _ = fmt.Printf("Starting the application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
