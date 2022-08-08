package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/home.html"))
	if err := tpl.ExecuteTemplate(w, "home.html", nil); err != nil {
		fmt.Println(err)
	}
}
