package public_controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (f *PublicRoute) GetQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.Db

	user := params["user"]
	mc := params["server"]

	rows, err := db.Query("SELECT name,message,date FROM messages WHERE name=? AND mc_server = ? AND LENGTH(message) > 30 ORDER BY RAND() LIMIT 1", user, mc)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		name    string
		message string
		date    string
	)
	for rows.Next() {
		err = rows.Scan(&name, &message, &date)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	if name == "" {
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	w.Write([]byte(`{"name": "` + name + `", "message": "` + message + `", "date": "` + date + `"}`))

	return
}
