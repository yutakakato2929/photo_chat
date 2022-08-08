package websocket

import (
	"fmt"
	"photo_chat/infra/mysql"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)
var ClientsInfo = make(map[*websocket.Conn][]string)
var Broadcast = make(chan mysql.Message)
var Upgrader = websocket.Upgrader{}

func HandleMessages() {
	for {
		msg := <-Broadcast
		for client := range Clients {
			if (ClientsInfo[client][0] == msg.Targetname && ClientsInfo[client][1] == msg.Username) ||
				(ClientsInfo[client][0] == msg.Username && ClientsInfo[client][1] == msg.Targetname) {
				err := client.WriteJSON(msg)
				if err != nil {
					fmt.Printf("error: %v", err)
					client.Close()
					delete(Clients, client)
				}
			}
		}
	}

}
