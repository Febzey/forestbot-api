package private_controllers

import (
	"fmt"
	"net/http"

	"github.com/febzey/forestbot-api/pkg/websocket"
	"github.com/gorilla/mux"
)

func (f *PrivateRoute) WebsocketAuth(w http.ResponseWriter, r *http.Request) {
	ws := f.Ws
	db := f.Db
	//get params
	vars := mux.Vars(r)
	server := vars["server"]
	key := vars["key"]

	//check if key is valid, if not return 401
	if key != "12345" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println(server)

	wss, err := ws.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	websocket.WebSocketMessageHandler(wss, db, server)
}
