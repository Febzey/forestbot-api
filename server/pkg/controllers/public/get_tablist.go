package public_controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/febzey/forestbot-api/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func (f *PublicRoute) GetTablist(w http.ResponseWriter, r *http.Request) {

	var contents webSocketContents

	params := mux.Vars(r)
	mc_server := params["server"]

	var url_ws string

	switch mc_server {
	case "simplyvanilla":
		url_ws = "ws-simply.forestbot.org"
	case "eusurvival":
		url_ws = "ws-eu.forestbot.org"
	case "uneasyvanilla":
		url_ws = "ws-uneasyvanilla.forestbot.org"
	case "localhost":
		url_ws = "localhost:8080"
	}

	u := url.URL{Scheme: "wss", Host: url_ws, Path: "/playerlist"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		w.Write([]byte(`{"error": "An error occured while connecting to the websocket"}`))
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			w.Write([]byte(`{"error": "An error occured while connecting to the websocket"}`))
			return
		}

		err = json.Unmarshal(message, &contents)
		if err != nil {
			w.Write([]byte(`{"error": "An error occured while connecting to the websocket"}`))
			return
		}

		utils.GenerateTablist(contents.PlayerListExtra, w)
		conn.Close()
		return
	}

}
