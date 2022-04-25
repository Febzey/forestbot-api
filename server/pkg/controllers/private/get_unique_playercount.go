package private_controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/**
 *
 * Get unique playercount for
 * specified server
 *
 * @param mc_server
 */
func (f *PrivateRoute) GetUniquePlayerCount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db := f.Db

	server := params["server"]

	rows, err := db.Query(
		"SELECT COUNT(*) as cnt FROM users WHERE mc_server = ?",
		server,
	)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`{"error": "Database error"}`))
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
			w.Write([]byte(`{"error": "Bad value"}`))
			return
		}
	}

	w.Write([]byte(`{"cnt": ` + fmt.Sprintf("%d", cnt) + `}`))

	return
}
