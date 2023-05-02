package main

import (
	"first_app/pkg/handlers"
	"fmt"
	"net/http"
)

var portNumber = ":8081"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_, _ = fmt.Printf("Starting the application on port %s", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}

}
