package main

import (
	"fmt"
	"log"

	"os"
	"os/signal"

	"github.com/febzey/forestbot-api/pkg/database"
	"github.com/febzey/forestbot-api/pkg/routes"
	"github.com/febzey/forestbot-api/pkg/websocket"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	db, err := database.CreateConnection()
	if err != nil {
		fmt.Println("Error creating connection to database: ", err)
	}

	//websocket := websocket.Create_websocket_server()

	//watch for new connections to websocket

	wsHub := websocket.NewHub()

	routes.PublicRoutes(router, db)
	routes.PrivateRoutes(router, db, wsHub)

	router.Use(mux.CORSMethodMiddleware(router))

	server := ServerConfig(router)
	StartServer(server)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Server is shutting down...")
	database.EndConnection(db)
	server.Close()
	log.Println("Server has shut down.")

}
