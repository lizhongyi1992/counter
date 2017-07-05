package main

import "database/sql"

type SQLConn struct {
	db *sql.DB
}

func NewSQLConn(connstr string) (*SQLConn, error) {
	db, e := sql.Open("mysql", connstr)
	if e != nil {
		_err(e)
		return nil, e
	}
	return &SQLConn{db: db}, e
}
