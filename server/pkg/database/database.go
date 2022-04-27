package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseOptions struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func CreateConnection() (*sql.DB, error) {
	dbOps := DatabaseOptions{
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Name:     os.Getenv("DATABASE"),
	}

	db, err := sql.Open("mysql", dbOps.User+":"+dbOps.Password+"@tcp("+dbOps.Host+":"+dbOps.Port+")/"+dbOps.Name)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to database")

	return db, nil
}

func EndConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Connection to database closed")
}

func SaveChat(db *sql.DB, username, message, server string) error {
	_, err := db.Exec(
		"INSERT INTO messages (name, message, date, mc_server) VALUES (?, ?, ?,?)",
		username,
		message,
		server,
		fmt.Sprintf("%d", time.Now().UnixMilli()),
	)
	if err != nil {
		return err
	}
	return nil
}
