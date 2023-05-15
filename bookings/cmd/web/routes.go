package main

import (
	"net/http"

	"github.com/andreichristian/bookings-app/pkg/config"
	"github.com/andreichristian/bookings-app/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// `http.FileServer(http.Dir("./static/"))`
	// This line is creating a new file server that serves static files.
	// The `http.Dir("./static/")` part is telling the file server to use the "static" directory
	// in your current directory (that's what the "./" part means) as the source of your static files.

	fileserver := http.FileServer(http.Dir("./static/"))

	// `mux.Handle("/static/*", http.StripPrefix("/static", fileserver))`
	// This line is telling the `mux` (which is an HTTP request multiplexer, or router)
	// to handle all requests where the URL path starts with "/static/".

	// `http.StripPrefix("/static", fileserver)`
	// This is a handler that serves HTTP requests by removing the given prefix ("/static")
	// from the request URL's path and then invoking the handler `fileserver`.

	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
