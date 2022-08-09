package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func renderTemplateHTML(file string, w http.ResponseWriter, data interface{}) {
	files := []string{
		"views/" + file + ".html",
		"views/base.html",
	}
	tmpt, err := template.ParseFiles(files...)

	if err != nil {
		panic("Error Template: " + err.Error())
	}

	errExec := tmpt.ExecuteTemplate(w, "base", data)

	if errExec != nil {
		panic("Error Execute: " + errExec.Error())
	}

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	renderTemplateHTML("index", w, nil)
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {
		nameEmploy := r.FormValue("name")
		log.Println(nameEmploy)
	} else {
		renderTemplateHTML("create", w, nil)
	}
}
