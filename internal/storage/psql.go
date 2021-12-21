package storage

import (
	//"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	log "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	"github.com/keenfury/api/config"
)

var PsqlDB *sqlx.DB

func init() {
	if config.StorageSQL {
		var err error
		connectionStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", config.DBUser, config.DBPass, config.DBDB, config.DBHost)
		if config.DBPass == "" {
			connectionStr = fmt.Sprintf("user=%s dbname=%s host=%s sslmode=disable", config.DBUser, config.DBDB, config.DBHost)
		}

		PsqlDB, err = sqlx.Connect("DB", connectionStr)
		if err != nil {
			log.Panicln("Could not connect with connection string:", connectionStr)
		}
	}
}
