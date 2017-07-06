package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SQLConn struct {
	db *sql.DB
}

func NewSQLConn(connstr string) (*SQLConn, error) {
	db, e := sql.Open("mysql", connstr)
	_dbg(connstr, db, e)
	if e != nil {
		_err(e)
		return nil, e
	}
	e = db.Ping()
	if e != nil {
		_err(e)
		return nil, e
	}
	return &SQLConn{db: db}, e
}
