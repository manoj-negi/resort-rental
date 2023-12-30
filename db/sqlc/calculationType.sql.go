// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: calculationType.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCalculationType = `-- name: CreateCalculationType :one
INSERT INTO calculation_type (
    description,
    is_deleted
) VALUES (
    $1,
    $2
) RETURNING id, description, is_deleted, created_at, updated_at
`

type CreateCalculationTypeParams struct {
	Description string      `json:"description"`
	IsDeleted   pgtype.Bool `json:"is_deleted"`
}

func (q *Queries) CreateCalculationType(ctx context.Context, arg CreateCalculationTypeParams) (CalculationType, error) {
	row := q.db.QueryRow(ctx, createCalculationType, arg.Description, arg.IsDeleted)
	var i CalculationType
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllCalculationType = `-- name: GetAllCalculationType :many
SELECT id, description, is_deleted, created_at, updated_at FROM calculation_type
`

func (q *Queries) GetAllCalculationType(ctx context.Context) ([]CalculationType, error) {
	rows, err := q.db.Query(ctx, getAllCalculationType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CalculationType{}
	for rows.Next() {
		var i CalculationType
		if err := rows.Scan(
			&i.ID,
			&i.Description,
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