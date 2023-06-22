package handlers

import (
	"net/http"

	"github.com/Denis-Andrei/goapp/pkg/config"
	"github.com/Denis-Andrei/goapp/pkg/models"
	"github.com/Denis-Andrei/goapp/pkg/render"
)


var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.Appconfig
}

// NewRepo creates a new repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlres
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
// m is a receiver and now all the handlers have access to the Repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
