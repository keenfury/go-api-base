package storage

import (
	//"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/keenfury/api/config"
)

var MySqlDB *sqlx.DB

func init() {
	if config.StorageSQL && config.StorageMysql {
		var err error
		connectionStr := fmt.Sprintf("%s:%s@%s/%s", config.DBUser, config.DBPass, config.DBHost, config.DBDB)
		if config.DBPass == "" {
			connectionStr = fmt.Sprintf("%s@%s/%s", config.DBUser, config.DBHost, config.DBDB)
		}
		MySqlDB, err = sqlx.Connect("mysql", connectionStr)
		if err != nil {
			log.Panicln("Could not connect with connection string:", connectionStr)
		}
	}
}
