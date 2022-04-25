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
 * Updating when a user leaves the game.
 *
 * @param user
 * @param user uuid
 * @param mc_server
 */

func (f *PrivateRoute) UpdateLeave(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UnixMilli()
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
		w.Write([]byte(`{"error": "Bad value has been passed"}`))
		return
	}

	db.Exec(
		"UPDATE users SET leaves = leaves + 1, lastseen = ? WHERE username = ? AND mc_server = ?",
		fmt.Sprintf("%d", now),
		Data.User,
		Data.Server,
	)

	w.Write([]byte(`{"success": "user leave updated"}`))

	return

}
