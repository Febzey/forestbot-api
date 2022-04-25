package private_controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (f *PrivateRoute) GetLiveChatChannels(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := f.Db

	server := params["server"]

	rows, err := db.Query(
		"SELECT channelID FROM livechats WHERE mc_server = ?",
		server,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while communicating with the database"}`))
		return
	}

	defer rows.Close()

	channels := []string{}
	for rows.Next() {
		var channel string
		err := rows.Scan(&channel)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"error": "a bad value has been passed"}`))
			return
		}
		channels = append(channels, channel)
	}

	w.Write([]byte(`{"channels": ` + fmt.Sprintf("%q", channels) + `}`))

	return

}
