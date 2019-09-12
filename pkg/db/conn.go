package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sirupsen/logrus"
	"github.com/tangxusc/auto_record/pkg/config"
	"os"
	"time"
)

var db *sql.DB

func Conn(ctx context.Context) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s/SQLExpress?database=%s", config.Instance.Db.Username, config.Instance.Db.Password,
		config.Instance.Db.Address, config.Instance.Db.Port, config.Instance.Db.Database)
	var e error
	db, e = sql.Open("mssql", dsn)
	if e != nil {
		logrus.Errorf("[db]连接出现错误,url:%v,错误:%v", dsn, e.Error())
		os.Exit(1)
	}
	db.SetConnMaxLifetime(time.Duration(config.Instance.Db.LifeTime) * time.Second)
	db.SetMaxOpenConns(config.Instance.Db.MaxOpen)
	db.SetMaxIdleConns(config.Instance.Db.MaxIdle)
}

func Disconnection(ctx context.Context) {
	if db != nil {
		db.Close()
	}
}

func Exec(sqlString string, param ...interface{}) error {
	logrus.Debugf("[db]Insert:%s,param:%v", sqlString, param)
	return Tx(func(tx *sql.Tx) error {
		stmt, e := tx.Prepare(sqlString)
		if e != nil {
			return e
		}
		defer stmt.Close()
		_, e = stmt.Exec(param...)
		return e
	})
}

func Tx(f func(tx *sql.Tx) error) error {
	logrus.Debugf("[db]Tx:%v", f)
	tx, e := db.Begin()
	if e != nil {
		return e
	}
	e = f(tx)
	if e != nil {
		defer tx.Rollback()
		return e
	}
	return tx.Commit()
}
