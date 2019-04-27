package config

import (
	"database/sql"
	"github.com/mmuflih/rest-api-doc/app"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mmuflih/envgo/conf"
)

func NewMysqlConn(cfg conf.Config) (error, *sql.DB) {
	dbUser := cfg.GetString(`mysql.user`)
	dbPass := cfg.GetString(`mysql.pass`)
	dbName := cfg.GetString(`mysql.database`)
	dbHost := cfg.GetString(`mysql.address`)
	dbPort := cfg.GetString(`mysql.port`)

	uri := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	app.Logger(uri)
	db, err := sql.Open("mysql", uri)

	err = db.Ping()

	return err, db
}
