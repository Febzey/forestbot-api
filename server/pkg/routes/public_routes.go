package routes

import (
	"database/sql"
	"net/http"

	"github.com/febzey/forestbot-api/pkg/types"

	public_controllers "github.com/febzey/forestbot-api/pkg/controllers/public"
	"github.com/gorilla/mux"
)

func PublicRoutes(router *mux.Router, db *sql.DB) {
	r := public_controllers.PublicRoute{
		Db: db,
	}

	var routes = []types.Route{
		{
			Method:      http.MethodGet,
			Pattern:     "/playtime/{user}/{server}",
			HandlerFunc: r.GetPlaytime,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/kd/{user}/{server}",
			HandlerFunc: r.GetKD,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/joins/{user}/{server}",
			HandlerFunc: r.GetJoins,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/lastseen/{user}/{server}",
			HandlerFunc: r.GetLastSeen,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/joindate/{user}/{server}",
			HandlerFunc: r.GetJoinDate,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/user/{user}/{server}",
			HandlerFunc: r.GetAllUserStats,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/messagecount/{user}/{server}",
			HandlerFunc: r.GetMessageCount,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/quote/{user}/{server}",
			HandlerFunc: r.GetQuote,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/lastdeath/{user}/{server}",
			HandlerFunc: r.GetLastDeath,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/lastmessage/{user}/{server}",
			HandlerFunc: r.GetLastMessage,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/firstmessage/{user}/{server}",
			HandlerFunc: r.GetFirstMessage,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/tab/{server}",
			HandlerFunc: r.GetTablist,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/topstat/{stat}/{server}",
			HandlerFunc: r.GetTopStat,
		},
		{
			Method:      http.MethodPost,
			Pattern:     "/addfact",
			HandlerFunc: r.AddFact,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/getfact/{id}",
			HandlerFunc: r.GetFact,
		},
		{
			Method:      http.MethodGet,
			Pattern:     "/randomfact",
			HandlerFunc: r.GetRandomFact,
		},
	}

	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
}
