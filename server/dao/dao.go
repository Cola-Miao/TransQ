package dao

import (
	"database/sql"
	"fmt"
	"github.com/Cola-Miao/TransQ/server/format"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitSqlite(dbPath string) error {
	format.FuncStart("InitSqlite")
	defer format.FuncEnd("InitSqlite")

	slt, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	err = slt.Ping()
	if err != nil {
		return fmt.Errorf("slt.Ping: %w", err)
	}

	db = slt

	return nil
}
