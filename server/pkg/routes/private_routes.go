package routes

import (
	"database/sql"
	"net/http"

	"github.com/febzey/forestbot-api/pkg/types"

	private_controllers "github.com/febzey/forestbot-api/pkg/controllers/private"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func PrivateRoutes(router *mux.Router, db *sql.DB, ws *websocket.Upgrader) {
	r := private_controllers.PrivateRoute{
		Db: db,
		Ws: ws,
	}

	var routes = []types.Route{
		//Post requests
		{
			Method:      http.MethodPost,
			Pattern:     "/saveplaytime",
			HandlerFunc: r.SavePlaytime,
		},
		{
			Method:      http.MethodPost,
			Pattern:     "/savechat",
			HandlerFunc: r.SaveChat,
		},
		{
			Method:      http.MethodPost,
			Pattern:     "/savepvpkill",
			HandlerFunc: r.SavePvpKill,
		},
		{
			Method:      http.MethodPost,
			Pattern:     "/savepvekill",
			HandlerFunc: r.SavePveKill,
		},
		{
			Method:      http.MethodPost,
			Pattern:     "/updateleave",
			HandlerFunc: r.UpdateLeave,
		},
		{
			Method:      http.MethodPost,
			Pattern:     "/updatejoin",
			HandlerFunc: r.UpdateJoin,
		},
		//Get Requests
		{
			Method:      http.MethodGet,
			Pattern:     "/getchannels/{server}",
			HandlerFunc: r.GetLiveChatChannels,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/uniqueplayers/{server}",
			HandlerFunc: r.GetUniquePlayerCount,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/ws-auth/{server}/{key}",
			HandlerFunc: r.WebsocketAuth,
		},
	}

	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}

}
