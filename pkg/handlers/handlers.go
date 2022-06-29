package handlers

import (
	"myapp/pkg/config"
	"myapp/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// holds data sent from handlers to templates
// type TemplateData struct {
// 	StringMap map[string]string
// 	IntMap    map[string]int
// 	FloatMap  map[string]float32
// 	Data      map[string]interface{}
// 	CSRFToken string
// 	Flash     string
// 	Warning   string
// 	Error     string
// }

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
