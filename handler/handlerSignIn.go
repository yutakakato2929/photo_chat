package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/signin.html"))
	if err := tpl.ExecuteTemplate(w, "signin.html", nil); err != nil {
		fmt.Println(err)
	}
}
