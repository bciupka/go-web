package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/config"
	"github.com/bciupka/go-mod/02_web_basics/04_templates_modules/pkg/models"
)

// no cache

func RenderTemplateOld(w http.ResponseWriter, templateName string) {
	renderedTemplate, _ := template.ParseFiles("./templates/"+templateName, "./templates/base.layout.tmpl")
	err := renderedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error there", err)
	}
}

// basic caching

var tc = make(map[string]*template.Template)

func RenderTemplateBasic(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]

	if !inMap {
		err = createCachedTemplate(t)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Created cached template")
	} else {
		log.Println("Using cached template")
	}
	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func createCachedTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	renderedTemplate, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = renderedTemplate
	return nil
}

// advanced caching

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateCacheForTemplates()
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template not found")
		return
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err = t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateCacheForTemplates() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
