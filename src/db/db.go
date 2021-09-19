package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

//Db postgres connection
func Db() *pgx.Conn {

	POSTGRESCONNECTIONSTRING := "postgres://deepakraj:kannan@localhost:5432/deepakdb"

	conn, _ := pgx.Connect(context.Background(), POSTGRESCONNECTIONSTRING)

	return conn
}
