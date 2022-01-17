package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/snail24365/hivocab-server/api"
	db "github.com/snail24365/hivocab-server/db/sqlc"
	"github.com/snail24365/hivocab-server/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db :", err)
	}

	store := db.NewStore(conn)
	db.InitializeDatabase(store)
	
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}