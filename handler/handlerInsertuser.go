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
	user, err := mysql.GetByAccount(r.FormValue("account"), db)
	if user.Account == "" {
		hashedPasswd := utility.HashStr(r.FormValue("passwd"), "sha256")
		err = mysql.InsertUser(r.FormValue("account"), r.FormValue("name"), hashedPasswd, db)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}

}
