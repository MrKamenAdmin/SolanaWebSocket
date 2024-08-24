package history_repo

import (
	"GorillaWebSocket/internal/delivery"
	"GorillaWebSocket/pkg/psql"
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

type Repo struct {
	*psql.Repo
}

func New(repo *psql.Repo) *Repo {
	return &Repo{
		repo,
	}
}

func (r *Repo) GetHistory(ctx context.Context) ([]delivery.History, error) {

	data := make([]delivery.History, 0)
	sql := `
		select h.capture_date, h.stake from "history" as h order by capture_date desc;
	`

	query, err := r.Db.Query(ctx, sql)
	if err != nil {
		return []delivery.History{}, err
	}

	for query.Next() {
		var h delivery.History
		err = query.Scan(&h.CaptureDate, &h.Stake)

		if err != nil {
			return []delivery.History{}, err
		}

		data = append(data, h)
	}

	return data, err
}

func (r *Repo) AddStake(ctx context.Context, stake uint64, captureDate time.Time) error {
	args := pgx.NamedArgs{
		"capture_date": captureDate,
		"stake":        stake,
	}

	sql := `
	insert into "history"
		(capture_date, stake)
	values
		(@capture_date, @stake);
	`

	_, err := r.Db.Exec(ctx, sql, args)

	if err != nil {
		return err
	}

	return nil
}
