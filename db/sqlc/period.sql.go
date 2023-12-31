// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: period.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPeriod = `-- name: CreatePeriod :one
INSERT INTO period (
    period,
    is_deleted
) VALUES (
    $1,
    $2
) RETURNING id, period, is_deleted, created_at, updated_at
`

type CreatePeriodParams struct {
	Period    string      `json:"period"`
	IsDeleted pgtype.Bool `json:"is_deleted"`
}

func (q *Queries) CreatePeriod(ctx context.Context, arg CreatePeriodParams) (Period, error) {
	row := q.db.QueryRow(ctx, createPeriod, arg.Period, arg.IsDeleted)
	var i Period
	err := row.Scan(
		&i.ID,
		&i.Period,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllPeriod = `-- name: GetAllPeriod :many
SELECT id, period, is_deleted, created_at, updated_at FROM period
`

func (q *Queries) GetAllPeriod(ctx context.Context) ([]Period, error) {
	rows, err := q.db.Query(ctx, getAllPeriod)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Period{}
	for rows.Next() {
		var i Period
		if err := rows.Scan(
			&i.ID,
			&i.Period,
			&i.IsDeleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
