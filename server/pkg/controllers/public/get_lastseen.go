package public_controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (f *PublicRoute) GetLastSeen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.Db

	user := params["user"]
	mc_server := params["server"]

	rows, err := db.Query("SELECT username, lastseen FROM users WHERE username = ? AND mc_server = ?", user, mc_server)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		username string
		lastseen string
	)
	for rows.Next() {
		err = rows.Scan(&username, &lastseen)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	if username == "" {
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	w.Write([]byte(`{"lastseen": "` + lastseen + `"}`))

	return
}
