package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/vpaklatzis/go-simple-bank/api"
	db "github.com/vpaklatzis/go-simple-bank/db/sqlc"
	"github.com/vpaklatzis/go-simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config file:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Could not connect to postgres:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Could not start server:", err)
	}
}
