package main

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

func GetDbConnection(typedb string, dsn string) (*sql.DB, error) {
	Debug("Open DB connection")
	db, err := sql.Open(typedb, dsn)
	HandleError(err, "DB Open")
	Info("DB Opened ...")
	return db, err
}

func GetDbConnectionFromCfg(cfg *viper.Viper) (*sql.DB, error) {
	return GetDbConnection(cfg.GetString("database.typedb"), cfg.GetString("database.dsn"))
}

func GenerateDsnForDb(dbuser string, dbpass string, dbhost string, dbport int, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbuser, dbpass, dbhost, dbport, dbname)
}
