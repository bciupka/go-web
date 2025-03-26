package handlers

import (
	"net/http"

	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/config"
	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/models"
	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

var Repo *Repository

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["hello"] = "Hello form GO"
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
