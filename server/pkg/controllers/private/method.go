package private_controllers

import (
	"database/sql"

	"github.com/gorilla/websocket"
)

type PrivateRoute struct {
	Db *sql.DB
	Ws *websocket.Upgrader
}
