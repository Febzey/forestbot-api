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
 * Checking if a user exists in the database,
 * if they exist, update some stats,
 * if they do not exist, create a new row for them.
 *
 */
func (f *PrivateRoute) UpdateJoin(w http.ResponseWriter, r *http.Request) {
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
		Uuid   string `json:"uuid"`
		Server string `json:"mc_server"`
	}

	er := json.Unmarshal(body, &Data)
	if err != nil {
		fmt.Println(er)
		w.Write([]byte(`{"error": "Bad value has been passed"}`))
		return
	}

	rows, er := db.Query(
		"SELECT username, uuid, mc_server FROM users WHERE uuid = ? AND mc_server = ?",
		Data.Uuid,
		Data.Server,
	)
	if er != nil {
		fmt.Println(er)
		w.Write([]byte(`{"error": "Database error"}`))
		return
	}

	defer rows.Close()

	var (
		username  string
		uuid      string
		mc_server string
	)

	for rows.Next() {
		err := rows.Scan(&username, &uuid, &mc_server)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"error": "Database error"}`))
			return
		}
	}

	if (uuid != "" && uuid == Data.Uuid) && (username != "" && username != Data.User) {
		_, err := db.Exec(
			"UPDATE users SET username = ? WHERE username = ? AND uuid = ? AND mc_server = ?",
			Data.User,
			username,
			Data.Uuid,
			Data.Server,
		)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"error": "Database error"}`))
			return
		}
		w.Write([]byte(`{"success": "user updated", "oldname": "` + username + `"}`))
		return
	}

	if username == "" && uuid == "" {
		_, err := db.Exec(
			"INSERT INTO users(username, joindate, uuid, joins, mc_server) VALUES (?,?,?,?,?)",
			Data.User,
			fmt.Sprintf("%d", now),
			Data.Uuid, 1,
			Data.Server,
		)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"error": "Database error"}`))
			return
		}
		w.Write([]byte(`{"success": "new user added", "newuser": "true"}`))
		return
	}

	db.Exec(
		"UPDATE users SET joins = joins + 1, lastseen = ? WHERE username = ? AND mc_server = ?",
		fmt.Sprintf("%d", now),
		username,
		Data.Server,
	)

	w.Write([]byte(`{"success": "user join updated"}`))

	return

}
