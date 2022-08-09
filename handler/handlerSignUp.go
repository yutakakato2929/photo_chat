package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandlerSignUp(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/signup.html"))
	if err := tpl.ExecuteTemplate(w, "signup.html", nil); err != nil {
		fmt.Println(err)
	}
}
