package websocket

import (
	// "database/sql"
	// "fmt"

	"fmt"
	"log"
	"sync"
	"time"

	//"log"

	//"github.com/febzey/forestbot-api/pkg/database"
	"github.com/gorilla/websocket"
)

// type IClients struct {
// 	Mu     sync.Mutex
// 	Active map[string]*websocket.Conn
// }

type UserMsg struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Server   string `json:"mc_server"`
}

type WebSocketHub struct {
	// the mutext to protect connections
	ConnectionsMx sync.RWMutex

	// Registered Connections
	Connections map[*Connection]struct{}

	ConnectionsTest map[string][]*Connection

	//we want to be able to have multiple websocket connections for the same server
	//maybe a map of maps
	//TODO: change key name to mc_server, put connection in struct, find a way to iterate through each connection under a certain mc_server.
	//TODO: create an array inside Connection, which holds all connections for the given minecraft server
	//check where mc_Server = mc_server

	//we need to filter between which messages are livechats,
	//and which are get_tablists, maybe stream msg to different channels
	//depending on the connection type

	// Inbound messages from the connections.
	Broadcast chan []byte

	LogMx sync.RWMutex
	Log   [][]byte
}

type Connection struct {
	// Buffered channel of outbound messages.
	Send chan []byte

	// The hub.
	H         *WebSocketHub
	Mc_server string
	WsType    string
}

var (
	APIkey string = "12345"
)

//create a map with an array is the value

func NewHub() *WebSocketHub {
	h := &WebSocketHub{
		ConnectionsMx:   sync.RWMutex{},
		Broadcast:       make(chan []byte),
		Connections:     make(map[*Connection]struct{}),
		ConnectionsTest: make(map[string][]*Connection),
	}

	go func() {
		for {
			msg := <-h.Broadcast

			fmt.Println(string(msg), "msg inside of newhub go func")
			h.ConnectionsMx.RLock()
			for c := range h.Connections {
				select {
				case c.Send <- msg:
				case <-time.After(1 * time.Second):
					log.Printf("Shutting down connection %v", c)
					h.RemoveConnection(c)
				}
			}
			h.ConnectionsMx.RUnlock()
		}
	}()

	return h
}

//
// Add or remove websocket connections to Connections map
// These methods will be accesible from the WebSocketHub struct
//
func (h *WebSocketHub) AddConnection(conn *Connection) {
	h.ConnectionsMx.Lock()
	defer h.ConnectionsMx.Unlock()
	mc_server := conn.Mc_server

	//add a connection to the ConnectionsTest map for the given mc_server if it doesnt exist
	if _, ok := h.ConnectionsTest[mc_server]; !ok {
		fmt.Println("Adding new connection for mc server")
		h.ConnectionsTest[mc_server] = append(h.ConnectionsTest[mc_server], conn)
	} else {
		fmt.Println("Adding existing connection for mc server")
		h.ConnectionsTest[mc_server] = append(h.ConnectionsTest[mc_server], conn)
	}

	for k, v := range h.ConnectionsTest {
		fmt.Println("key:", k, "value:", v)
	}

	return

}
func (h *WebSocketHub) RemoveConnection(conn *Connection) {
	h.ConnectionsMx.Lock()
	defer h.ConnectionsMx.Unlock()
	mc_server := conn.Mc_server

	//do the opposite of AddConnection
	//remove the connection from the ConnectionsTest map for the given mc_server if it exists
	if _, ok := h.ConnectionsTest[mc_server]; ok {
		fmt.Println("Removing connection for mc server")
		for i, v := range h.ConnectionsTest[mc_server] {
			if v == conn {
				h.ConnectionsTest[mc_server] = append(h.ConnectionsTest[mc_server][:i], h.ConnectionsTest[mc_server][i+1:]...)
				return
			} else {
				fmt.Println("no connection found, so cant remove it")
				return
			}
		}
	}
}

//
// Write and Read message handlers.
// These methods will be accesible from the Connections struct
//
func (c *Connection) Reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	defer fmt.Println("done reader")
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println(string(message))

		//will broadcast message to all connected websockets
		c.H.Broadcast <- message
	}
}

func (c *Connection) Writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	defer fmt.Println("done write")
	//print c.Send

	return
	// for message := range c.Send {
	// 	fmt.Println("message from writer", message)
	// }

	// for {
	// 	fmt.Println("supposed to be writing to all clients here...")
	// }
	// for message := range c.Send {
	// 	fmt.Println("Sending message")
	// 	err := wsConn.WriteMessage(websocket.TextMessage, message)
	// 	if err != nil {
	// 		break
	// 	}
	// }
}

//create a websocket server that listens on port 8080 and returns the websocket server when done
// func Create_websocket_server() *websocket.Upgrader {
// 	upgrader := websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 		CheckOrigin: func(r *http.Request) bool {
// 			return true
// 		},
// 	}
// 	return &upgrader
// }

func WebsocketAuth(key string) bool {
	if key != APIkey {
		return false
	}
	return true
}
