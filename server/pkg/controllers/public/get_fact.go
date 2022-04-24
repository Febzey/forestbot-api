package public_controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (f *PublicRoute) GetFact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := f.Db
	id := params["id"]

	rows, err := db.Query(
		"SELECT username, fact, date, id FROM facts WHERE id = ?",
		id,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Error while performing lookup"}`))
		return
	}

	defer rows.Close()

	var (
		username string
		fact     string
		date     int
		ID       int
	)

	for rows.Next() {
		err = rows.Scan(&username, &fact, &date, &ID)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(`{"error": "Error while performing lookup"}`))
			return
		}
	}

	if username == "" {
		w.Write([]byte(`{"error": "Fact not found"}`))
		return
	}

	w.Write([]byte(`{"username": "` + username + `", "fact": "` + fact + `", "date": ` + strconv.Itoa(date) + `}`))

}
