package main

import (
	"fmt"
	"log"

	"os"
	"os/signal"

	"github.com/febzey/forestbot-api/pkg/configs"
	"github.com/febzey/forestbot-api/pkg/database"
	"github.com/febzey/forestbot-api/pkg/routes"
	"github.com/febzey/forestbot-api/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	db, err := database.CreateConnection()
	if err != nil {
		fmt.Println("Error creating connection to database: ", err)
	}

	router := mux.NewRouter()

	routes.PublicRoutes(router, db)
	routes.PrivateRoutes(router, db)

	router.Use(mux.CORSMethodMiddleware(router))

	server := configs.ServerConfig(router)

	utils.StartServer(server)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	database.EndConnection(db)

	log.Println("Server is shutting down...")
}
