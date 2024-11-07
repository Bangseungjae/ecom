package main

import (
	"database/sql"
	"github.com/Bangseungjae/ecom/cmd/api"
	"github.com/Bangseungjae/ecom/config"
	"github.com/Bangseungjae/ecom/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	database, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(database)

	server := api.NewAPIServer(":8080", database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
