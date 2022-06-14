package websocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	MainKey     string = "12345"
	LiveChatKey string = "livechat"
)

// hub for all websocket connections
type WebSocketHub struct {
	ConnectionsMx sync.RWMutex

	// a map that takes the minecraft server as key,
	// and an array of connected websocket clients as the value
	ConnectedClients map[string][]*Connection

	// all messages recieved from websockets will go here to be broadcasted
	Broadcast chan []byte
}

// a single websocket connection
type Connection struct {
	// Buffered channel of outbound messages.
	// we should change this to livechat message output.
	Send chan []byte

	// The hub.
	H *WebSocketHub

	//the minecraft server this connection is for
	Mc_server string
	Type      string
}

type LiveChatMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Server   string `json:"mc_server"`
}

//create a map with an array is the value

func NewHub() *WebSocketHub {
	h := &WebSocketHub{
		ConnectionsMx:    sync.RWMutex{},
		Broadcast:        make(chan []byte),
		ConnectedClients: make(map[string][]*Connection),
	}

	go func() {
		for {
			msg := <-h.Broadcast

			h.ConnectionsMx.RLock()
			for _, v := range h.ConnectedClients {
				for _, conn := range v {
					fmt.Println(conn.Mc_server)
					conn.Send <- msg
				}
			}

			//loop through ConnectedClients if key is eusurvival'
			// for _, v := range h.ConnectedClients["eusurvival"] {

			// }

			// for c := range h.Connections {
			// 	select {
			// 	case c.Send <- msg:
			// 	case <-time.After(1 * time.Second):
			// 		log.Printf("Shutting down connection %v", c)
			// 		h.RemoveConnection(c)
			// 	}
			// }
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
	h.ConnectedClients[conn.Mc_server] = append(h.ConnectedClients[conn.Mc_server], conn)
	fmt.Println("adding connection to map")
	// for k, v := range h.ConnectedClients {
	// 	fmt.Println("key:", k, "value:", v)
	// }
	return
}
func (h *WebSocketHub) RemoveConnection(conn *Connection) {
	h.ConnectionsMx.Lock()
	defer h.ConnectionsMx.Unlock()
	mc_server := conn.Mc_server
	if _, ok := h.ConnectedClients[mc_server]; ok {
		fmt.Println("Removing connection for mc server")
		for i, v := range h.ConnectedClients[mc_server] {
			if v == conn {

				h.ConnectedClients[mc_server] = append(h.ConnectedClients[mc_server][:i], h.ConnectedClients[mc_server][i+1:]...)
				return
			}
		}
	}
}

// read messages from websocket connection
func (c *Connection) Reader(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	defer fmt.Println("done reader")
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		// fmt.Println(string(message))

		//broadcasting message from websockets to the channel
		c.H.Broadcast <- message
	}
}

// write messages to websocket connection
func (c *Connection) Writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	defer fmt.Println("done write")

	for message := range c.Send {
		fmt.Println(string(message), " <-- from writer")
	}

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

func WebsocketAuth(key string) (string, bool) {
	if key == MainKey {
		return "main", true
	}
	if key == LiveChatKey {
		return "livechat", true
	}
	return "Not authorized.", false
}
