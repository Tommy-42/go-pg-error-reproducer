package main

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/tommy-42/go-pg-reproducer/user"
)

var (
	DB *pg.DB
)

func init() {
	// Options to connect to PostgreSQL
	option := &pg.Options{
		User:     "gaming",
		Password: "game42",
		Database: "game",
		Addr:     "localhost:5432",
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * time.Millisecond,
	}
	// Connect to PostgresQL
	DB = pg.Connect(option)

	err := DB.CreateTable((*user.User)(nil), &orm.CreateTableOptions{
		Temp:          true, // create temp table
		FKConstraints: true,
	})
	if err != nil {
		log.Printf("Err CreateTable: %+v", err)
		panic(err)
	}

	err = DB.Insert(&user.User{Id: "1", Username: "test"})
	if err != nil {
		log.Printf("Err Inserting user")
		panic(err)
	}
}

func main() {
	var user *user.User

	err := DB.Model(user).Where("username = ?", "test").Select()
	if err != nil {
		if err == pg.ErrNoRows {
			log.Printf("ErrNoRows")
			panic(err)
		}

		log.Printf("Err: %+v", err)
		panic(err)
	}
}
