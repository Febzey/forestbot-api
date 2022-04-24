package controllers

import (
	"database/sql"
)

type Routes struct {
	DB *sql.DB
}

//
// Playtime endpoint handler
//
// func (f *Routes) GetPlaytime(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	db := f.DB

// 	user := params["user"]
// 	mc_server := params["server"]

// 	rows, err := db.Query("SELECT username, playtime FROM users WHERE username = ? AND mc_server = ?", user, mc_server)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		playtime int
// 		username string
// 	)
// 	for rows.Next() {
// 		err = rows.Scan(&username, &playtime)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"playtime": ` + strconv.Itoa(playtime) + `}`))

// 	return
// }

//
// kills / deaths endpoint handler
//
// func (f *Routes) GetKD(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	db := f.DB

// 	user := params["user"]
// 	mc_server := params["server"]

// 	rows, err := db.Query("SELECT username, kills, deaths FROM users WHERE username = ? AND mc_server = ?", user, mc_server)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username string
// 		kills    int
// 		deaths   int
// 	)
// 	for rows.Next() {
// 		err = rows.Scan(&username, &kills, &deaths)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"kills": ` + strconv.Itoa(kills) + `, "deaths": ` + strconv.Itoa(deaths) + `}`))

// 	return
// }

//
// Join count endpoint handler
//
// func (f *Routes) GetJoins(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	db := f.DB

// 	user := params["user"]
// 	mc_server := params["server"]

// 	rows, err := db.Query("SELECT username, joins FROM users WHERE username = ? AND mc_server = ?", user, mc_server)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username string
// 		joins    int
// 	)
// 	for rows.Next() {
// 		err = rows.Scan(&username, &joins)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"joins": ` + strconv.Itoa(joins) + `}`))

// 	return
// }

//
// Lastseen endpoint handler
//
// func (f *Routes) GetLastSeen(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	db := f.DB

// 	user := params["user"]
// 	mc_server := params["server"]

// 	rows, err := db.Query("SELECT username, lastseen FROM users WHERE username = ? AND mc_server = ?", user, mc_server)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username string
// 		lastseen string
// 	)
// 	for rows.Next() {
// 		err = rows.Scan(&username, &lastseen)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"lastseen": "` + lastseen + `"}`))

// 	return
// }

//
// Joindate endpoint handler
//
// func (f *Routes) GetJoinDate(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	db := f.DB

// 	user := params["user"]
// 	mc_server := params["server"]

// 	rows, err := db.Query("SELECT username, joindate FROM users WHERE username = ? AND mc_server = ?", user, mc_server)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username string
// 		joindate string
// 	)
// 	for rows.Next() {
// 		err = rows.Scan(&username, &joindate)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"joindate": "` + joindate + `"}`))

// 	return
// }

//
// Get all user stats endpoint
//
// func (f *Routes) GetAllUserStats(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	db := f.DB

// 	user := params["user"]
// 	mc := params["server"]

// 	rows, err := db.Query("SELECT * FROM users WHERE username = ? AND mc_server = ?", user, mc)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username        string
// 		playtime        int
// 		kills           int
// 		deaths          int
// 		joins           int
// 		leaves          int
// 		lastseen        string
// 		joindate        string
// 		uuid            string
// 		lastdeathString string
// 		lastdeathTime   int
// 		mc_server       string
// 	)
// 	for rows.Next() {
// 		err = rows.Scan(
// 			&username,
// 			&kills,
// 			&deaths,
// 			&joindate,
// 			&lastseen,
// 			&uuid,
// 			&playtime,
// 			&joins,
// 			&leaves,
// 			&lastdeathTime,
// 			&lastdeathString,
// 			&mc_server,
// 		)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"username": "` + username + `", "kills": ` + strconv.Itoa(kills) + `, "deaths": ` + strconv.Itoa(deaths) + `, "joins": ` + strconv.Itoa(joins) + `, "leaves": ` + strconv.Itoa(leaves) + `, "lastseen": "` + lastseen + `", "joindate": "` + joindate + `", "uuid": "` + uuid + `", "playtime": ` + strconv.Itoa(playtime) + `, "lastdeath": "` + lastdeathString + `", "mc_server": "` + mc_server + `"}`))

// 	return
// }

//
// Total message count endpoint handler
//
// func (f *Routes) GetMessageCount(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	db := f.DB

// 	user := params["user"]
// 	mc := params["server"]

// 	rows, err := db.Query("SELECT name,COUNT(name) AS cnt FROM messages WHERE name = ? AND mc_server = ?", user, mc)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		name string
// 		cnt  int
// 	)
// 	for rows.Next() {
// 		err = rows.Scan(&name, &cnt)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if name == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"messagecount": ` + strconv.Itoa(cnt) + `}`))

// 	return
// }

//
// Random quote endpoint handler
//
// func (f *Routes) GetQuote(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	db := f.DB

// 	user := params["user"]
// 	mc := params["server"]

// 	rows, err := db.Query("SELECT name,message,date FROM messages WHERE name=? AND mc_server = ? AND LENGTH(message) > 30 ORDER BY RAND() LIMIT 1", user, mc)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		name    string
// 		message string
// 		date    string
// 	)
// 	for rows.Next() {
// 		err = rows.Scan(&name, &message, &date)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if name == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"name": "` + name + `", "message": "` + message + `", "date": "` + date + `"}`))

// 	return
// }

//
// lastdeath endpoint handler
//
// func (f *Routes) GetLastDeath(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	db := f.DB

// 	user := params["user"]
// 	mc := params["server"]

// 	rows, err := db.Query("SELECT username, lastdeathTime, lastdeathString from users WHERE username = ? AND mc_server = ?", user, mc)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username        string
// 		lastdeathTime   int
// 		lastdeathString string
// 	)
// 	for rows.Next() {
// 		err = rows.Scan(&username, &lastdeathTime, &lastdeathString)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"death": "` + lastdeathString + `", "time": ` + strconv.Itoa(lastdeathTime) + `}`))

// 	return
// }

// type webSocketContents struct {
// 	PlayerList      []string     `json:"playerList"`
// 	PlayerListExtra []utils.User `json:"playerListExtra"`
// 	UniquePlayers   int          `json:"uniquePlayers"`
// }

//
//lastmessage endpoint handler
//
// func (f *Routes) GetLastMessage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	user := params["user"]
// 	mc_server := params["server"]

// 	//get lastmessage from the user in database
// 	db := f.DB
// 	rows, err := db.Query("SELECT name, message, date FROM messages WHERE name = ? AND mc_server = ? ORDER BY date DESC LIMIT 1", user, mc_server)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username string
// 		message  string
// 		date     string
// 	)

// 	for rows.Next() {
// 		err = rows.Scan(&username, &message, &date)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"username": "` + username + `", "message": "` + message + `", "date": "` + date + `"}`))

// }

//
// First message endpoint handler
//
// func (f *Routes) GetFirstMessage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	user := params["user"]
// 	mc_server := params["server"]

// 	db := f.DB
// 	rows, err := db.Query("SELECT name, message, date FROM messages WHERE name = ? AND mc_server = ? ORDER BY date ASC LIMIT 1", user, mc_server)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username string
// 		message  string
// 		date     string
// 	)

// 	for rows.Next() {
// 		err = rows.Scan(&username, &message, &date)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "User not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"username": "` + username + `", "message": "` + message + `", "date": "` + date + `"}`))

// }

//
// Tablist endpoint handler
//
// func (f *Routes) GetTablist(w http.ResponseWriter, r *http.Request) {

// 	var contents webSocketContents

// 	params := mux.Vars(r)
// 	mc_server := params["server"]

// 	var url_ws string

// 	switch mc_server {
// 	case "simplyvanilla":
// 		url_ws = "ws-simply.forestbot.org"
// 	case "eusurvival":
// 		url_ws = "ws-eu.forestbot.org"
// 	case "uneasyvanilla":
// 		url_ws = "ws-uneasyvanilla.forestbot.org"
// 	case "localhost":
// 		url_ws = "localhost:8080"
// 	}

// 	u := url.URL{Scheme: "wss", Host: url_ws, Path: "/playerlist"}

// 	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		w.Write([]byte(`{"error": "An error occured while connecting to the websocket"}`))
// 		return
// 	}

// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			w.Write([]byte(`{"error": "An error occured while connecting to the websocket"}`))
// 			return
// 		}

// 		err = json.Unmarshal(message, &contents)
// 		if err != nil {
// 			w.Write([]byte(`{"error": "An error occured while connecting to the websocket"}`))
// 			return
// 		}

// 		utils.GenerateTablist(contents.PlayerListExtra, w)
// 		conn.Close()
// 		return
// 	}

// }

// func (f *Routes) GetTopStat(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	db := f.DB

// 	stat := params["stat"]
// 	server := params["server"]

// 	allowedStats := []string{
// 		"kills", "deaths", "joins", "playtime",
// 	}

// 	if !utils.ArrayContains(allowedStats, stat) {
// 		w.Write([]byte(`{"error": "Cannot get top stat for the stat you specified"}`))
// 		return
// 	}

// 	rows, err := db.Query(
// 		`SELECT username, `+stat+` FROM users WHERE mc_server = ? ORDER BY `+stat+` DESC LIMIT 5`,
// 		server,
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	userStatArr := []string{}

// 	for rows.Next() {
// 		var (
// 			username   string
// 			stat_value int
// 		)

// 		err = rows.Scan(&username, &stat_value)
// 		if err != nil {
// 			fmt.Println(err)
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}

// 		userStatArr = append(userStatArr, `{"username": "`+username+`", "stat": `+strconv.Itoa(stat_value)+`}`)

// 	}

// 	w.Write([]byte(`{"top_stat": [` + strings.Join(userStatArr, ",") + `]}`))

// 	return

// }

// func (f *Routes) AddFact(w http.ResponseWriter, r *http.Request) {
// 	now := time.Now().UnixMilli()

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Request body has an error"}`))
// 		return
// 	}

// 	defer r.Body.Close()

// 	var Data struct {
// 		User   string `json:"user"`
// 		Fact   string `json:"fact"`
// 		Server string `json:"mc_server"`
// 	}

// 	err = json.Unmarshal(body, &Data)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(badValue))
// 		return
// 	}

// 	db := f.DB

// 	res, err := db.Exec(
// 		`INSERT INTO facts (username,fact,mc_server,date) VALUES (?,?,?,?)`,
// 		Data.User,
// 		Data.Fact,
// 		Data.Server,
// 		now,
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	id, _ := res.LastInsertId()

// 	w.Write([]byte(`{"success": "Fact saved", "id": ` + strconv.Itoa(int(id)) + `}`))

// 	return

// }

// func (f *Routes) GetFact(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	db := f.DB
// 	id := params["id"]

// 	rows, err := db.Query(
// 		"SELECT username, fact, date, id FROM facts WHERE id = ?",
// 		id,
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username string
// 		fact     string
// 		date     int
// 		ID       int
// 	)

// 	for rows.Next() {
// 		err = rows.Scan(&username, &fact, &date, &ID)
// 		if err != nil {
// 			fmt.Println(err)
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "Fact not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"username": "` + username + `", "fact": "` + fact + `", "date": ` + strconv.Itoa(date) + `}`))

// }

// func (f *Routes) GetRandomFact(w http.ResponseWriter, r *http.Request) {

// 	db := f.DB

// 	rows, err := db.Query("SELECT username, fact, date, id FROM facts ORDER BY RAND() LIMIT 1")
// 	if err != nil {
// 		fmt.Println(err)
// 		w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 		return
// 	}

// 	defer rows.Close()

// 	var (
// 		username string
// 		fact     string
// 		date     int
// 		ID       int
// 	)

// 	for rows.Next() {
// 		err = rows.Scan(&username, &fact, &date, &ID)
// 		if err != nil {
// 			fmt.Println(err)
// 			w.Write([]byte(`{"error": "Error while performing lookup"}`))
// 			return
// 		}
// 	}

// 	if username == "" {
// 		w.Write([]byte(`{"error": "Fact not found"}`))
// 		return
// 	}

// 	w.Write([]byte(`{"username": "` + username + `", "fact": "` + fact + `", "date": ` + strconv.Itoa(date) + `, "id": ` + strconv.Itoa(ID) + `}`))

// 	return

// }
