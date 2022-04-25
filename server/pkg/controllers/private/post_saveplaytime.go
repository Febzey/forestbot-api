package private_controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 *
 * Incrementing users playtime by
 * 60000 milliseconds. (60 seconds)
 *
 * @param user
 * @param mc_server
 */
func (f *PrivateRoute) SavePlaytime(w http.ResponseWriter, r *http.Request) {
	db := f.Db

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Request body has an error"}`))
		return
	}

	defer r.Body.Close()

	var Data struct {
		User   string `json:"user"`
		Server string `json:"mc_server"`
	}

	err = json.Unmarshal(body, &Data)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Bad value"}`))
		return
	}

	rows, err := db.Query(
		"UPDATE users SET playtime = playtime + 60000 WHERE username = ? and mc_server = ?",
		Data.User,
		Data.Server,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Database error"}`))
		return
	}

	defer rows.Close()

	w.Write([]byte(`{"success": "Playtime was saved"}`))

	return
}
