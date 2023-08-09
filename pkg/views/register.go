package views

import (
	"html/template"
)

func Register() *template.Template {
	temp := template.Must(template.ParseFiles("templates/register.gohtml"))
	return temp
}
