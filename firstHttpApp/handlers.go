package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, templ string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templ)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
