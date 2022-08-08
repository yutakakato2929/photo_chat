package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandlerRegister(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/register.html"))
	if err := tpl.ExecuteTemplate(w, "register.html", nil); err != nil {
		fmt.Println(err)
	}
}
