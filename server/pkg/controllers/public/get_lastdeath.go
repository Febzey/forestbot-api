package public_controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/febzey/forestbot-api/pkg/utils"
	"github.com/gorilla/mux"
)

func (f *PublicRoute) GetLastDeath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.Db

	user := params["user"]
	mc := params["server"]

	rows, err := db.Query("SELECT username, lastdeathTime, lastdeathString from users WHERE username = ? AND mc_server = ?", user, mc)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		username        string
		lastdeathTime   int
		lastdeathString string
	)
	for rows.Next() {
		err = rows.Scan(&username, &lastdeathTime, &lastdeathString)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	if username == "" {
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	w.Write([]byte(`{"death": "` + lastdeathString + `", "time": ` + strconv.Itoa(lastdeathTime) + `}`))

	return
}

type webSocketContents struct {
	PlayerList      []string     `json:"playerList"`
	PlayerListExtra []utils.User `json:"playerListExtra"`
	UniquePlayers   int          `json:"uniquePlayers"`
}
