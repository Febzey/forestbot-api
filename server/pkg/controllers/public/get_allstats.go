package public_controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (f *PublicRoute) GetAllUserStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.Db

	user := params["user"]
	mc := params["server"]

	rows, err := db.Query("SELECT * FROM users WHERE username = ? AND mc_server = ?", user, mc)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		username        string
		playtime        int
		kills           int
		deaths          int
		joins           int
		leaves          int
		lastseen        string
		joindate        string
		uuid            string
		lastdeathString string
		lastdeathTime   int
		mc_server       string
	)
	for rows.Next() {
		err = rows.Scan(
			&username,
			&kills,
			&deaths,
			&joindate,
			&lastseen,
			&uuid,
			&playtime,
			&joins,
			&leaves,
			&lastdeathTime,
			&lastdeathString,
			&mc_server,
		)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	if username == "" {
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	w.Write([]byte(`{"username": "` + username + `", "kills": ` + strconv.Itoa(kills) + `, "deaths": ` + strconv.Itoa(deaths) + `, "joins": ` + strconv.Itoa(joins) + `, "leaves": ` + strconv.Itoa(leaves) + `, "lastseen": "` + lastseen + `", "joindate": "` + joindate + `", "uuid": "` + uuid + `", "playtime": ` + strconv.Itoa(playtime) + `, "lastdeath": "` + lastdeathString + `", "mc_server": "` + mc_server + `"}`))

	return
}
