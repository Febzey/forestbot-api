package utils

import (
	"log"
	"net/http"
)

func StartServer(server *http.Server) {

	go func() {
		log.Println("Server is starting...")
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

}
