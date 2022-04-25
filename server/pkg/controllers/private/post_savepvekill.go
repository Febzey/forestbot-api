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
 * Updating when a user dies or kills themselves.
 * pve messages
 *
 * @param victim
 * @param deathMsg
 * @param mc_server
 */
func (f *PrivateRoute) SavePveKill(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UnixMilli()
	db := f.Db

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while reading message"}`))
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
		w.Write([]byte(`{"error": "bad value"}`))
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
