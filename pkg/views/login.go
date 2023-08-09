package views

import (
	"html/template"
)

func Login() *template.Template {
	temp := template.Must(template.ParseFiles("templates/login.gohtml"))
	return temp
}
