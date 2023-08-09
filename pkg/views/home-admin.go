package views

import (
	"html/template"
)

func HomeAdmin() *template.Template {
	temp := template.Must(template.ParseFiles("templates/home-admin.gohtml"))
	return temp
}
