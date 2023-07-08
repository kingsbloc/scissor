package utils

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func HandleSqlError(err error) error {
	var perr *pgconn.PgError
	errors.As(err, &perr)
	switch perr.Code {
	case "1452":
		return errors.New("invalid record entered")
	case "23505":
		return errors.New("record already exists")
	}
	return errors.New("an error occurred")
}
