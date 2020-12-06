package dao

import (
	"context"
	"database/sql"
)

type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
}

type MainDao struct {
	DB *sql.DB
}

func New(db *sql.DB) (dao *MainDao, cf func()) {
	dao = &MainDao{
		DB: db,
	}
	cf = dao.Close
	return
}

func (d *MainDao) Close() {
	d.DB.Close()
}

func (d *MainDao) Ping(ctx context.Context) (err error) {
	return
}
