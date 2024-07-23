package main

import (
	"log"

	"github.com/asfandyarjalil/golang-practice-project/config"
	"github.com/asfandyarjalil/golang-practice-project/server"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("******* Starting Golang Project ******")
	log.Println("****** Initializing configuration ******")
	config := config.InitConfig(".env")
	log.Println("****** Initializing DATABASE ****** ")
	dbHandler := server.InitDatabse(config)
	log.Println("**** Initializing HTTP server *****")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
