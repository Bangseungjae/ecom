package main

import (
	"errors"
	"log"
	"os"

	"github.com/Bangseungjae/ecom/config"
	"github.com/Bangseungjae/ecom/db"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/golang-migrate/migrate/v4"

	mysqlDriver "github.com/go-sql-driver/mysql"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

/**
https://pkg.go.dev/github.com/golang-migrate/migrate/v4

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
*/

func main() {
	cfg := mysqlDriver.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	database, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysqlMigrate.WithInstance(database, &mysqlMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}
}
