package public_controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (f *PublicRoute) GetLastMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user := params["user"]
	mc_server := params["server"]

	//get lastmessage from the user in database
	db := f.Db
	rows, err := db.Query("SELECT name, message, date FROM messages WHERE name = ? AND mc_server = ? ORDER BY date DESC LIMIT 1", user, mc_server)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		username string
		message  string
		date     string
	)

	for rows.Next() {
		err = rows.Scan(&username, &message, &date)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	if username == "" {
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	w.Write([]byte(`{"username": "` + username + `", "message": "` + message + `", "date": "` + date + `"}`))

}
