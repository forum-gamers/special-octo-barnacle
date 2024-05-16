package database

import (
	"database/sql"
	"log"
	"os"

	h "github.com/forum-gamers/special-octo-barnacle/helpers"
	_ "github.com/lib/pq"
)

func Conn() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "user=apple password=password dbname=forum-gamers-tournament sslmode=disable"
	}
	db, err := sql.Open("postgres", url)
	h.PanicIfError(err)
	h.PanicIfError(db.Ping())

	DB = db
	log.Println("connected to the database")
}
