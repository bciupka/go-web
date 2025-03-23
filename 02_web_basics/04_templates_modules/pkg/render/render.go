package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateOld(w http.ResponseWriter, templateName string) {
	renderedTemplate, _ := template.ParseFiles("./templates/"+templateName, "./templates/base.layout.tmpl")
	err := renderedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error there", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
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
