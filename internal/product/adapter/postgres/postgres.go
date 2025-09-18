package postgres

import (
	"github.com/amagkn/my-go-clean-architecture-template/pkg/postgres"
)

type Postgres struct {
	pool *postgres.Pool
}

func New(p *postgres.Pool) *Postgres {
	return &Postgres{
		pool: p,
	}
}
