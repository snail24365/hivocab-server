package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/snail24365/hivocab-server/util"
)

var testQueries *Queries
var testDB *sql.DB
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db :", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}