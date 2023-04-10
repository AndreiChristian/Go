package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, templ string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templ, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// chech to see if we already have the template in our cahce
	_, inMap := tc[t]

	if !inMap {
		// need to create a template
		log.Println("creating template and adding it to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}
	tc[t] = tmpl

	return nil
}
