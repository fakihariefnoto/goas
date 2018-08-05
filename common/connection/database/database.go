package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var UserDB *sql.DB

func Connect() (err error) {
	UserDB, _ := sql.Open("sqlite3", "./nraboy.db")
}
