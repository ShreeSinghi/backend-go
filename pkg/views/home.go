package views

import (
	"html/template"
)

func Home() *template.Template {
	temp := template.Must(template.ParseFiles("templates/home.gohtml"))
	return temp
}
