package private_controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 *
 * Saving user chat messages
 *
 * @param user
 * @params message
 * @param mc_server
 */
func (f *PrivateRoute) SaveChat(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Request body has an error"}`))
		return
	}

	defer r.Body.Close()

	var Data struct {
		Message string `json:"message"`
		User    string `json:"user"`
		Server  string `json:"mc_server"`
	}

	err = json.Unmarshal(body, &Data)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Bad value has been passed"}`))
		return
	}

	db := f.Db
	now := time.Now().UnixMilli()

	rows, err := db.Query(
		"INSERT INTO messages (name, message, date, mc_server) VALUES(?,?,?,?)",
		Data.User,
		Data.Message,
		fmt.Sprintf("%d", now),
		Data.Server,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while communicating with the database"}`))
		return
	}

	defer rows.Close()

	w.Write([]byte(`{"success": "chat message saved"}`))

	return
}
