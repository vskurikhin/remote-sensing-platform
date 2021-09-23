package domain

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	DAO struct {
		EPollSettings *ePollSettings
	}
	ePollSettings struct {
		poolRo *pgxpool.Pool
		poolRw *pgxpool.Pool
	}
)

func New(poolRo *pgxpool.Pool, poolRw *pgxpool.Pool) *DAO {
	return &DAO{
		EPollSettings: &ePollSettings{poolRo: poolRo, poolRw: poolRw},
	}
}
