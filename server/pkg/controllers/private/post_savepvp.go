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
 * Updating when a user kills another user.
 * pvp messages
 *
 * @param victim
 * @param murderer
 * @param deathMsg
 * @param mc_server
 */
func (f *PrivateRoute) SavePvpKill(w http.ResponseWriter, r *http.Request) {
	db := f.Db
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

	db.Exec(
		"UPDATE users SET kills = kills + 1 WHERE username = ? AND mc_server = ?",
		Data.Murderer, Data.Server,
	)

	w.Write([]byte(`{"success": "pvp death saved"}`))

	return

}
