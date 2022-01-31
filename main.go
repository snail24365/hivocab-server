package main

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/snail24365/hivocab-server/api"
	db "github.com/snail24365/hivocab-server/db/sqlc"
	"github.com/snail24365/hivocab-server/util"
)

func main() {
	log.Info("app start", 999)

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
	
	server, err := api.NewServer(config, *store)
	if err != nil {
		log.Fatal("server creation err :", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}