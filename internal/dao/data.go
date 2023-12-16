package dao

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mangohow/mangokit-template/internal/conf"
)

func NewMysqlClient(conf *conf.Data) (*sqlx.DB, func(), error) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	var err error
	db, err := sqlx.ConnectContext(timeoutCtx, "mysql", conf.Mysql.DataSourceName)
	if err != nil {
		return nil, nil, err
	}

	db.SetMaxOpenConns(int(conf.Mysql.MaxOpenConns))
	db.SetMaxIdleConns(int(conf.Mysql.MaxIdleConns))

	cleanup := func() {
		_ = db.Close()
	}

	return db, cleanup, nil
}

func NewFakeMysqlClient(conf *conf.Data) (*sqlx.DB, func(), error) {
	return &sqlx.DB{}, func() {}, nil
}
