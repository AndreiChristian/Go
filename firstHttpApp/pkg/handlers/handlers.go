package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/render"
	"net/http"
)

// The repository type
type Repository struct {
	App *config.AppConfig
}

// The repository used by the handlers
var Repo *Repository

// New Repo creates a new repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New Repo sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
