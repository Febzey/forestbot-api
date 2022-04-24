package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	DbErr    string = `{"error": "Error while communicating with the database"}`
	BadValue string = `{"error": "a bad value has been passed"}`
	BadBody  string = `{"error": "Request body is incorrect."}`
)

/**
 *
 * Get an array of live chat channel id's, for users who
 * have set up live chat's within their discord servers,
 * specified by mc_server
 *
 * @param mc_server string
 * @returns channels[]
 */
func (f *Routes) GetLiveChatChannels(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := f.DB

	server := params["server"]

	rows, err := db.Query(
		"SELECT channelID FROM livechats WHERE mc_server = ?",
		server,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(DbErr))
		return
	}

	defer rows.Close()

	channels := []string{}
	for rows.Next() {
		var channel string
		err := rows.Scan(&channel)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(BadValue))
			return
		}
		channels = append(channels, channel)
	}

	w.Write([]byte(`{"channels": ` + fmt.Sprintf("%q", channels) + `}`))

	return

}

/**
 *
 * Saving user chat messages
 *
 * @param user
 * @params message
 * @param mc_server
 */
func (f *Routes) SaveChat(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(BadValue))
		return
	}

	db := f.DB
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
		w.Write([]byte(DbErr))
		return
	}

	defer rows.Close()

	w.Write([]byte(`{"success": "chat message saved"}`))

	return
}

/**
 *
 * Incrementing users playtime by
 * 60000 milliseconds. (60 seconds)
 *
 * @param user
 * @param mc_server
 */
func (f *Routes) SavePlaytime(w http.ResponseWriter, r *http.Request) {
	db := f.DB

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
		w.Write([]byte(BadValue))
		return
	}

	rows, err := db.Query(
		"UPDATE users SET playtime = playtime + 60000 WHERE username = ? and mc_server = ?",
		Data.User,
		Data.Server,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(DbErr))
		return
	}

	defer rows.Close()

	w.Write([]byte(`{"success": "Playtime was saved"}`))

	return
}

/**
 *
 * Updating when a user kills another user.
 * pvp messages
 *
 * @param victim
 * @param murderer
 * @param deathMsg
 * @param mc_server
 */
func (f *Routes) SavePvpKill(w http.ResponseWriter, r *http.Request) {
	db := f.DB
	now := time.Now().UnixMilli()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while reading message"}`))
		return
	}

	defer r.Body.Close()

	var Data struct {
		Victim   string `json:"victim"`
		Murderer string `json:"murderer"`
		DeathMsg string `json:"deathmsg"`
		Server   string `json:"mc_server"`
	}

	err = json.Unmarshal(body, &Data)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(BadValue))
		return
	}

	db.Exec(
		"UPDATE users SET deaths = deaths + 1, lastdeathString = ?, lastdeathTime = ? WHERE username = ? AND mc_server = ?",
		Data.DeathMsg,
		fmt.Sprintf("%d", now),
		Data.Victim,
		Data.Server,
	)

	db.Exec(
		"UPDATE users SET kills = kills + 1 WHERE username = ? AND mc_server = ?",
		Data.Murderer, Data.Server,
	)

	w.Write([]byte(`{"success": "pvp death saved"}`))

	return

}

/**
 *
 * Updating when a user dies or kills themselves.
 * pve messages
 *
 * @param victim
 * @param deathMsg
 * @param mc_server
 */
func (f *Routes) SavePveKill(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UnixMilli()
	db := f.DB

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(BadBody))
		return
	}

	defer r.Body.Close()

	var Data struct {
		Victim   string `json:"victim"`
		DeathMsg string `json:"deathmsg"`
		Server   string `json:"mc_server"`
	}

	err = json.Unmarshal(body, &Data)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(BadValue))
		return
	}

	db.Exec(
		"UPDATE users SET deaths = deaths + 1, lastdeathString = ?, lastdeathTime = ? WHERE username = ? AND mc_server = ?",
		Data.DeathMsg,
		fmt.Sprintf("%d", now),
		Data.Victim,
		Data.Server,
	)

	w.Write([]byte(`{"success": "pve death saved"}`))

	return

}

/**
 *
 * Updating when a user leaves the game.
 *
 * @param user
 * @param user uuid
 * @param mc_server
 */

func (f *Routes) UpdateLeave(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UnixMilli()
	db := f.DB

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(BadBody))
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
		w.Write([]byte(BadValue))
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

/**
 *
 * Get unique playercount for
 * specified server
 *
 * @param mc_server
 */
func (f *Routes) GetUniquePlayerCount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := f.DB

	server := params["server"]

	rows, err := db.Query(
		"SELECT COUNT(*) as cnt FROM users WHERE mc_server = ?",
		server,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(DbErr))
		return
	}

	defer rows.Close()

	var (
		cnt int
	)

	for rows.Next() {
		err := rows.Scan(&cnt)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(BadValue))
			return
		}
	}

	w.Write([]byte(`{"cnt": ` + fmt.Sprintf("%d", cnt) + `}`))

	return
}

/**
 *
 * Checking if a user exists in the database,
 * if they exist, update some stats,
 * if they do not exist, create a new row for them.
 *
 */
func (f *Routes) UpdateJoin(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UnixMilli()
	db := f.DB

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(BadBody))
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
		w.Write([]byte(BadValue))
		return
	}

	rows, er := db.Query(
		"SELECT username, uuid, mc_server FROM users WHERE uuid = ? AND mc_server = ?",
		Data.Uuid,
		Data.Server,
	)
	if er != nil {
		fmt.Println(er)
		w.Write([]byte(DbErr))
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
			w.Write([]byte(BadValue))
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
			w.Write([]byte(DbErr))
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
			w.Write([]byte(DbErr))
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
