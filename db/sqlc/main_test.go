package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/21toffy/relinc/util"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://tofunmi:toffy123@172.17.0.2:5432/relinc_db?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")

	if err != nil {
		log.Fatal("can not load config file: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())

}
