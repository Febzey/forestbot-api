package public_controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func (f *PublicRoute) AddFact(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UnixMilli()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Request body has an error"}`))
		return
	}

	defer r.Body.Close()

	var Data struct {
		User   string `json:"user"`
		Fact   string `json:"fact"`
		Server string `json:"mc_server"`
	}

	err = json.Unmarshal(body, &Data)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Bad input"}`))
		return
	}

	db := f.Db

	res, err := db.Exec(
		`INSERT INTO facts (username,fact,mc_server,date) VALUES (?,?,?,?)`,
		Data.User,
		Data.Fact,
		Data.Server,
		now,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	id, _ := res.LastInsertId()

	w.Write([]byte(`{"success": "Fact saved", "id": ` + strconv.Itoa(int(id)) + `}`))

	return

}
