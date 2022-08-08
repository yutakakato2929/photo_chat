package handler

import (
	"net/http"
	"photo_chat/sessions"
)

func HandlerLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "user-basic-info")
	delete(session.Values, "userinfo")
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
}
