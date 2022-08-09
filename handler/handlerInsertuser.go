package handler

import (
	"fmt"
	"net/http"
	"photo_chat/infra/mysql"
	"photo_chat/utility"
)

func HandlerInsertuser(w http.ResponseWriter, r *http.Request) {
	db, err := mysql.Openmysql()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	user_ac, err := mysql.GetUserByFlag(r.FormValue("account"), "ACCOUNT", db)
	user_nm, err := mysql.GetUserByFlag(r.FormValue("name"), "NAME", db)
	if user_ac.Account == "" && user_nm.Name == "" {
		hashedPasswd := utility.HashStr(r.FormValue("passwd"), "sha256")
		err = mysql.InsertUser(r.FormValue("account"), r.FormValue("name"), hashedPasswd, db)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/signin?pa=success", http.StatusSeeOther)
	} else if user_ac.Account == "" {
		http.Redirect(w, r, "/signup?pa=name", http.StatusAccepted)
	} else {
		http.Redirect(w, r, "/signup?pa=account", http.StatusAccepted)
	}
}
