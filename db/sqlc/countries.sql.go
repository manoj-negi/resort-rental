// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: countries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCountry = `-- name: CreateCountry :one
INSERT INTO country (
    iso2,
    short_name,
    long_name,
    numcode,
    calling_code,
    cctld,
    is_deleted
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) RETURNING id, iso2, short_name, long_name, numcode, calling_code, cctld, is_deleted, created_at, updated_at
`

type CreateCountryParams struct {
	Iso2        string      `json:"iso2"`
	ShortName   string      `json:"short_name"`
	LongName    string      `json:"long_name"`
	Numcode     pgtype.Text `json:"numcode"`
	CallingCode string      `json:"calling_code"`
	Cctld       string      `json:"cctld"`
	IsDeleted   pgtype.Bool `json:"is_deleted"`
}

func (q *Queries) CreateCountry(ctx context.Context, arg CreateCountryParams) (Country, error) {
	row := q.db.QueryRow(ctx, createCountry,
		arg.Iso2,
		arg.ShortName,
		arg.LongName,
		arg.Numcode,
		arg.CallingCode,
		arg.Cctld,
		arg.IsDeleted,
	)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.Iso2,
		&i.ShortName,
		&i.LongName,
		&i.Numcode,
		&i.CallingCode,
		&i.Cctld,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCountry = `-- name: DeleteCountry :one
DELETE FROM country WHERE id = $1
RETURNING id, iso2, short_name, long_name, numcode, calling_code, cctld, is_deleted, created_at, updated_at
`

func (q *Queries) DeleteCountry(ctx context.Context, id int32) (Country, error) {
	row := q.db.QueryRow(ctx, deleteCountry, id)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.Iso2,
		&i.ShortName,
		&i.LongName,
		&i.Numcode,
		&i.CallingCode,
		&i.Cctld,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllCountries = `-- name: GetAllCountries :many
SELECT id, iso2, short_name, long_name, numcode, calling_code, cctld, is_deleted, created_at, updated_at FROM country
`

func (q *Queries) GetAllCountries(ctx context.Context) ([]Country, error) {
	rows, err := q.db.Query(ctx, getAllCountries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Country{}
	for rows.Next() {
		var i Country
		if err := rows.Scan(
			&i.ID,
			&i.Iso2,
			&i.ShortName,
			&i.LongName,
			&i.Numcode,
			&i.CallingCode,
			&i.Cctld,
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

const getCountry = `-- name: GetCountry :one
SELECT id, iso2, short_name, long_name, numcode, calling_code, cctld, is_deleted, created_at, updated_at FROM country WHERE id = $1
`

func (q *Queries) GetCountry(ctx context.Context, id int32) (Country, error) {
	row := q.db.QueryRow(ctx, getCountry, id)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.Iso2,
		&i.ShortName,
		&i.LongName,
		&i.Numcode,
		&i.CallingCode,
		&i.Cctld,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCountry = `-- name: UpdateCountry :one
UPDATE country
SET
    iso2 = $2,
    short_name = $3,
    long_name = $4,
    numcode = $5,
    calling_code = $6,
    cctld = $7,
    is_deleted = $8
WHERE id = $1
RETURNING id, iso2, short_name, long_name, numcode, calling_code, cctld, is_deleted, created_at, updated_at
`

type UpdateCountryParams struct {
	ID          int32       `json:"id"`
	Iso2        string      `json:"iso2"`
	ShortName   string      `json:"short_name"`
	LongName    string      `json:"long_name"`
	Numcode     pgtype.Text `json:"numcode"`
	CallingCode string      `json:"calling_code"`
	Cctld       string      `json:"cctld"`
	IsDeleted   pgtype.Bool `json:"is_deleted"`
}

func (q *Queries) UpdateCountry(ctx context.Context, arg UpdateCountryParams) (Country, error) {
	row := q.db.QueryRow(ctx, updateCountry,
		arg.ID,
		arg.Iso2,
		arg.ShortName,
		arg.LongName,
		arg.Numcode,
		arg.CallingCode,
		arg.Cctld,
		arg.IsDeleted,
	)
	var i Country
	err := row.Scan(
		&i.ID,
		&i.Iso2,
		&i.ShortName,
		&i.LongName,
		&i.Numcode,
		&i.CallingCode,
		&i.Cctld,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
