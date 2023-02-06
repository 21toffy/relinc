package main

import (
	"database/sql"
	"log"
	"relinc/api"
	db "relinc/db/sqlc"
	"relinc/util"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://tofunmi:toffy123@172.17.0.2:5432/relinc_db?sslmode=disable"
// 	serverAddress = "0.0.0.0:8081"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can  not load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(*store)

	err = server.Start(config.ServerAddres)
	if err != nil {
		log.Fatal("cannot start server")
	}

}
