package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

func renderTemplate(w http.ResponseWriter, templ string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templ)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
