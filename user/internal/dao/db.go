package dao

import (
	"database/sql"

	"github.com/BurntSushi/toml"
	"github.com/go-sql-driver/mysql"
	sdb "github.com/high-quality-sausages/sausage-common-go/db"
)

const (
	_dsn = "%s:%s@tcp(%s)/%s?charset=utf8"
)

func NewDB() (db *sql.DB, err error) {
	var dbConf struct {
		MysqlConf mysql.Config
	}

	if _, err := toml.DecodeFile("../configs/mysql.toml", &dbConf); err != nil {
		panic(err)
	}

	db = sdb.NewMysql(&dbConf.MysqlConf)
	return
}
