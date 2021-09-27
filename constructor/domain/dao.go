package domain

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	DAO struct {
		EPoll           *ePoll
		EPollChannel    *ePollChannel
		EPollDesign     *ePollDesign
		EPollSettings   *ePollSettings
		FScreenMain     *fScreenMain
		FScreenWelcome  *fScreenWelcome
		FScreenComplete *fScreenComplete
	}
	ePoll struct {
		poolRo *pgxpool.Pool
		poolRw *pgxpool.Pool
	}
	ePollChannel struct {
		poolRo *pgxpool.Pool
		poolRw *pgxpool.Pool
	}
	ePollDesign struct {
		poolRo *pgxpool.Pool
		poolRw *pgxpool.Pool
	}
	ePollSettings struct {
		poolRo *pgxpool.Pool
		poolRw *pgxpool.Pool
	}
	fScreenMain struct {
		poolRo *pgxpool.Pool
		poolRw *pgxpool.Pool
	}
	fScreenWelcome struct {
		poolRo *pgxpool.Pool
		poolRw *pgxpool.Pool
	}
	fScreenComplete struct {
		poolRo *pgxpool.Pool
		poolRw *pgxpool.Pool
	}
)

func New(poolRo *pgxpool.Pool, poolRw *pgxpool.Pool) *DAO {
	return &DAO{
		EPoll:           &ePoll{poolRo: poolRo, poolRw: poolRw},
		EPollChannel:    &ePollChannel{poolRo: poolRo, poolRw: poolRw},
		EPollDesign:     &ePollDesign{poolRo: poolRo, poolRw: poolRw},
		EPollSettings:   &ePollSettings{poolRo: poolRo, poolRw: poolRw},
		FScreenMain:     &fScreenMain{poolRo: poolRo, poolRw: poolRw},
		FScreenWelcome:  &fScreenWelcome{poolRo: poolRo, poolRw: poolRw},
		FScreenComplete: &fScreenComplete{poolRo: poolRo, poolRw: poolRw},
	}
}
