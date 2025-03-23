package handlers

import (
	"net/http"

	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

}
