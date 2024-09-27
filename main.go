package main

import (
	"database/sql"
	"log"

	"github.com/hoangthienclub/simple-bank/api"
	db "github.com/hoangthienclub/simple-bank/db/sqlc"
	"github.com/hoangthienclub/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start the server:", err)
	}
}
