package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"photo_chat/infra/mysql"
	"photo_chat/sessions"
)

func HandlerInitial(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "user-basic-info")
	cuser := session.Values["userinfo"].(*(sessions.Cuser))
	if cuser == nil {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	db, err := mysql.Openmysql()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	users, err := mysql.GetUserAll(db)
	if err != nil {
		fmt.Println(err)
	}
	value := map[string]interface{}{
		"name":  cuser.Name,
		"users": users,
	}
	tpl := template.Must(template.ParseFiles("templates/initial.html"))
	if err := tpl.ExecuteTemplate(w, "initial.html", value); err != nil {
		fmt.Println(err)
	}
}
