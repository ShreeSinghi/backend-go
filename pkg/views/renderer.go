package views

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, fname string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + fname + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
