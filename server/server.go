package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func ConnectionUrlBuilder(n string) (string, error) {
	var url string
	switch n {
	case "server":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)

	default:
		return "", fmt.Errorf("Invalid connection name: %s", n)
	}

	return url, nil
}

func ServerConfig(router *mux.Router) *http.Server {
	serverConnectionUrl, _ := ConnectionUrlBuilder("server")
	readTImeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return &http.Server{
		Handler:     router,
		Addr:        serverConnectionUrl,
		ReadTimeout: time.Duration(readTImeoutSecondsCount) * time.Second,
	}
}

func StartServer(server *http.Server) {
	go func() {
		log.Println("Server is starting...")
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Error starting server: ", err)
		}
	}()

}
