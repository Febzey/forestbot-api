package websocket

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	//"log"
	"net/http"

	//"github.com/febzey/forestbot-api/pkg/database"
	"github.com/gorilla/websocket"
)

type IClients struct {
	mu     sync.Mutex
	active map[string]*websocket.Conn
}

type UserMsg struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Server   string `json:"mc_server"`
}

//create a websocket server that listens on port 8080 and returns the websocket server when done
func Create_websocket_server() *websocket.Upgrader {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &upgrader
}

func ChatMsgHandler() {

}

var (
	Clients = IClients{
		mu:     sync.Mutex{},
		active: make(map[string]*websocket.Conn),
	}
)

func WebSocketMessageHandler(conn *websocket.Conn, db *sql.DB, mc_server string) {

	fmt.Println("Client added")

	Clients.mu.Lock()
	Clients.active[mc_server] = conn
	Clients.mu.Unlock()

	//remove conn from map where key is mc_server
	defer func() {
		Clients.mu.Lock()
		delete(Clients.active, mc_server)
		Clients.mu.Unlock()
	}()

	defer fmt.Println("Client removed")

	for {
		var userMsg UserMsg

		err := conn.ReadJSON(&userMsg)
		if err != nil {
			log.Println(err)
		}
		log.Println(userMsg.Username + ": " + userMsg.Message + " on " + userMsg.Server)

		// err = database.SaveChat(db, userMsg.Username, userMsg.Message, userMsg.Server)
		// if err != nil {
		// 	log.Println("Error saving chat message", err)
		// }

	}
}
