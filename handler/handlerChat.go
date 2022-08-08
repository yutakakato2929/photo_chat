package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"photo_chat/infra/mysql"
	"photo_chat/sessions"
	"photo_chat/websocket"
)

func HandlerChat(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "user-basic-info")
	cuser := session.Values["userinfo"].(*(sessions.Cuser))
	username := cuser.Name
	targetname := r.URL.Path[len("/chat/"):]
	db, err := mysql.Openmysql()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	msgs, err := mysql.GetMessageByName(username, targetname, db)
	if err != nil {
		fmt.Println(err)
	}
	value := map[string]interface{}{
		"username":   username,
		"targetname": targetname,
		"messages":   msgs,
	}
	tpl := template.Must(template.ParseFiles("templates/chat.html"))
	if err := tpl.ExecuteTemplate(w, "chat.html", value); err != nil {
		fmt.Println(err)
	}
}
func HandlerWs(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "user-basic-info")
	username := session.Values["userinfo"].(*(sessions.Cuser)).Name
	targetname := r.URL.Path[len("/ws/"):]
	ws, err := websocket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer ws.Close()
	websocket.Clients[ws] = true
	websocket.ClientsInfo[ws] = []string{username, targetname}
	for {
		var msg mysql.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(websocket.Clients, ws)
			break
		}
		InsertMessageDatabase(msg)
		websocket.Broadcast <- msg
	}
}

func InsertMessageDatabase(msg mysql.Message) {
	db, err := mysql.Openmysql()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	err = mysql.InsertMessage(msg, db)
	if err != nil {
		fmt.Println(err)
	}
}
