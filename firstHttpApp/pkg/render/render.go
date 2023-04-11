package render

import (
	"bytes"
	"log"
	"myapp/pkg/config"
	"net/http"
	"path/filepath"
	"text/template"
)

var app *config.AppConfig

// new templates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, templ string) {

	// get the template cache from the app config

	tc := app.TemplateCache
	// get requested template from cahce

	t, ok := tc[templ]

	if !ok {
		log.Fatal("error parsing files: ", ok)
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCahce() (map[string]*template.Template, error) {
	myCahce := map[string]*template.Template{}

	// GET ALL THE FILES THAT END WITH .PAGE.TMPL FROM ./TEMPLATES

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCahce, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCahce, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCahce, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCahce, err
			}
		}

		myCahce[name] = ts

	}

	return myCahce, nil
}
