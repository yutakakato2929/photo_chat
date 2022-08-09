package main

import (
	"net/http"
	"os"
	"photo_chat/handler"
	"photo_chat/websocket"
)

func main() {
	http.HandleFunc("/signin", handler.HandlerSignIn)
	http.HandleFunc("/login", handler.HandlerLogin)
	http.HandleFunc("/signup", handler.HandlerSignUp)
	http.HandleFunc("/insertuser", handler.HandlerInsertuser)
	http.HandleFunc("/initial", handler.HandlerInitial)
	http.HandleFunc("/chat/", handler.HandlerChat)
	http.HandleFunc("/ws/", handler.HandlerWs)
	http.HandleFunc("/logout", handler.HandlerLogout)
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset/"))))
	go websocket.HandleMessages()
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
