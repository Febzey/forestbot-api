package routes

import (
	"database/sql"
	"net/http"

	"github.com/febzey/forestbot-api/pkg/types"

	private_controllers "github.com/febzey/forestbot-api/pkg/controllers/private"
	"github.com/gorilla/mux"
)

func PrivateRoutes(router *mux.Router, db *sql.DB) {
	r := private_controllers.PrivateRoute{
		Db: db,
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
	}

	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}

}
