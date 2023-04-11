package main

import (
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCahce()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	// we set UseCache to false in development to use "hot reload", meaning there is no need to
	// restart the server every time we make modifications to the templates. This happens because in
	// the render.go file we recreate the cache every single time when we get a request if useCache is
	// meaning we dont use the cache. This costs performance as the cache is recreated every single time
	// and therefore should be set to true in production
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	_ = http.ListenAndServe(portNumber, nil)
}
