package public_controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (f *PublicRoute) GetKD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := f.Db

	user := params["user"]
	mc_server := params["server"]

	rows, err := db.Query("SELECT username, kills, deaths FROM users WHERE username = ? AND mc_server = ?", user, mc_server)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		username string
		kills    int
		deaths   int
	)
	for rows.Next() {
		err = rows.Scan(&username, &kills, &deaths)
		if err != nil {
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	if username == "" {
		w.Write([]byte(`{"error": "User not found"}`))
		return
	}

	w.Write([]byte(`{"kills": ` + strconv.Itoa(kills) + `, "deaths": ` + strconv.Itoa(deaths) + `}`))

	return
}
