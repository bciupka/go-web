package main

import (
	"log"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, templateName string) {
	renderedTemplate, _ := template.ParseFiles("./templates/" + templateName)
	err := renderedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error there", err)
	}
}
