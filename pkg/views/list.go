package views

import (
	"html/template"
)

func ListPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/list.html"))
	return temp

}
