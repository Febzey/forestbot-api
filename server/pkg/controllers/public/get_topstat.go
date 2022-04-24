package public_controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/febzey/forestbot-api/pkg/utils"
	"github.com/gorilla/mux"
)

func (f *PublicRoute) GetTopStat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := f.Db

	stat := params["stat"]
	server := params["server"]

	allowedStats := []string{
		"kills", "deaths", "joins", "playtime",
	}

	if !utils.ArrayContains(allowedStats, stat) {
		w.Write([]byte(`{"error": "Cannot get top stat for the stat you specified"}`))
		return
	}

	rows, err := db.Query(
		`SELECT username, `+stat+` FROM users WHERE mc_server = ? ORDER BY `+stat+` DESC LIMIT 5`,
		server,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	userStatArr := []string{}

	for rows.Next() {
		var (
			username   string
			stat_value int
		)

		err = rows.Scan(&username, &stat_value)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}

		userStatArr = append(userStatArr, `{"username": "`+username+`", "stat": `+strconv.Itoa(stat_value)+`}`)

	}

	w.Write([]byte(`{"top_stat": [` + strings.Join(userStatArr, ",") + `]}`))

	return

}
