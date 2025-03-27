package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/config"
	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/handlers"
	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateCacheForTemplates()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	fmt.Println("Server starting on port", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	// http.ListenAndServe(port, nil)
}
