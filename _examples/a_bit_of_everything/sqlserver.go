package main

import (
	"context"
	"database/sql"

	models "github.com/xo/xo/_examples/a_bit_of_everything/sqlserver"
)

func runSqlserver(ctx context.Context, db *sql.DB) error {
	return models.Run()
}