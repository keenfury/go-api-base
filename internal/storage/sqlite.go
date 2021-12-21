package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/keenfury/api/config"
	_ "github.com/mattn/go-sqlite3"
)

var SqliteDB *sqlx.DB

func init() {
	if config.StorageSQL && config.StorageSqlite {
		var err error
		connectionStr := fmt.Sprintf("%s?cache=shared&mode=wrc", config.SqlitePath)
		SqliteDB, err = sqlx.Open("sqlite3", connectionStr)
		if err != nil {
			log.Panicln("Could not connect with connection string:", connectionStr)
		}
		SqliteDB.SetMaxOpenConns(1)
	}
}
