package repository

import (
	"context"
	"database/sql"
	"time"
)

type TimeRepository interface {
	CurrentTime(ctx context.Context) (time.Time, error)
}

type PostgresTimeRepository struct {
	db *sql.DB
}

func NewPostgresTimeRepository(db *sql.DB) *PostgresTimeRepository {
	return &PostgresTimeRepository{db: db}
}

func (r *PostgresTimeRepository) CurrentTime(ctx context.Context) (time.Time, error) {
	var currentTime time.Time
	err := r.db.QueryRowContext(ctx, "SELECT NOW()").Scan(&currentTime)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime, nil
}
