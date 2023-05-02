package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	_, _ = fmt.Fprintf(w, "Home page")
}

func About(w http.ResponseWriter, r *http.Request) {

	_, _ = fmt.Fprintf(w, "About page")
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8090", nil)

}
