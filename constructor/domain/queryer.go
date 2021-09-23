package domain

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type queryer interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}
