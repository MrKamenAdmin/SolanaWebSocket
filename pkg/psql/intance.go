package psql

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	Db *pgxpool.Pool
}

var (
	pgInstance *Repo
	pgOnce     sync.Once
)

func NewPool(ctx context.Context, connString string) *Repo {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)

		if err != nil {
			panic(err)
		}

		pgInstance = &Repo{db}
	})

	return pgInstance
}

func (pg *Repo) Ping(ctx context.Context) error {
	return pg.Db.Ping(ctx)
}

func (pg *Repo) Close() {
	pg.Db.Close()
}
