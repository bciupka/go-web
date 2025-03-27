package main

import (
	"net/http"

	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/config"
	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(a *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
