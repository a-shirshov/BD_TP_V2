package utils

import (
	"github.com/jackc/pgx"
)

func InitPostgresDB() (*pgx.ConnPool, error) {
	config := pgx.ConnConfig{
		User:                 "postgres",
		Database:             "postgres",
		Password:             "password",
		Port: 						5436,	
		PreferSimpleProtocol: false,
	}
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 100,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}
	return pgx.NewConnPool(connPoolConfig)
}

func Prepare(db *pgx.ConnPool) error {
	for _, query := range queries {
		_, err := db.Prepare(query.Name, query.Query)
		if err != nil {
			return err
		}
	}
	return nil
}