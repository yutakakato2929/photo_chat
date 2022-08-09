package handler

import (
	"fmt"
	"net/http"
	"photo_chat/infra/mysql"
	"photo_chat/sessions"
	"photo_chat/utility"
	"time"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	db, err := mysql.Openmysql()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	hashedPasswd := utility.HashStr(r.FormValue("passwd"), "sha256")
	user, err := mysql.GetUserByFlag(r.FormValue("account"), "ACCOUNT", db)
	if user.Account != "" {
		if hashedPasswd == user.Passwd {
			session, _ := sessions.Store.Get(r, "user-basic-info")
			session.Values["userinfo"] = &sessions.Cuser{Name: user.Name, LoginTime: time.Now(), LogoutTime: time.Now()}
			err := session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/initial", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/signin?pa=passwd", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/signin?pa=account", http.StatusSeeOther)
	}
}
