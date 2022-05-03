package private_controllers

import (
	"fmt"
	"net/http"
	"sync"

	ws "github.com/febzey/forestbot-api/pkg/websocket"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func (wsh *WsHandler) GetTablist(w http.ResponseWriter, r *http.Request) {
	//get websocket connection from
}

func (wsh *WsHandler) WebsocketConnect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	mc_server := vars["server"]
	key := vars["key"]

	// if ws.WebsocketAuth(key) == false {
	// 	fmt.Println("Invalid Websocket Key")
	// 	w.Write([]byte("Invalid Websocket Key"))
	// 	return
	// }

	wsType, isAuthed := ws.WebsocketAuth(key)
	if !isAuthed {
		fmt.Println("Invalid Websocket Key")
		w.Write([]byte("Invalid Websocket Key"))
		return
	}

	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	c := &ws.Connection{
		Send:      make(chan []byte, 256),
		H:         wsh.H,
		Mc_server: mc_server,
		Type:      wsType,
	}
	c.H.AddConnection(c)
	defer c.H.RemoveConnection(c)

	var wg sync.WaitGroup
	wg.Add(2)
	go c.Writer(&wg, wsConn)
	go c.Reader(&wg, wsConn)
	wg.Wait()
	fmt.Println("done")
	wsConn.Close()
}
