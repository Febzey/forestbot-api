package private_controllers

import (
	"database/sql"

	ws "github.com/febzey/forestbot-api/pkg/websocket"
)

type PrivateRoute struct {
	Db *sql.DB
}

type WsHandler struct {
	H *ws.WebSocketHub
}
