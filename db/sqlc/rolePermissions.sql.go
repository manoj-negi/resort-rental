// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: rolePermissions.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRolePermission = `-- name: CreateRolePermission :one
INSERT INTO roles_permission (
    role_id,
    permission_id,
    is_deleted
) VALUES (
    $1,
    $2,
    $3
) RETURNING id, role_id, permission_id, is_deleted, created_at, updated_at
`

type CreateRolePermissionParams struct {
	RoleID       int32       `json:"role_id"`
	PermissionID int32       `json:"permission_id"`
	IsDeleted    pgtype.Bool `json:"is_deleted"`
}

func (q *Queries) CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) (RolesPermission, error) {
	row := q.db.QueryRow(ctx, createRolePermission, arg.RoleID, arg.PermissionID, arg.IsDeleted)
	var i RolesPermission
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.PermissionID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRolePermission = `-- name: DeleteRolePermission :one
DELETE FROM roles_permission WHERE id = $1
RETURNING id, role_id, permission_id, is_deleted, created_at, updated_at
`

func (q *Queries) DeleteRolePermission(ctx context.Context, id int32) (RolesPermission, error) {
	row := q.db.QueryRow(ctx, deleteRolePermission, id)
	var i RolesPermission
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.PermissionID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllRolePermissions = `-- name: GetAllRolePermissions :many
SELECT id, role_id, permission_id, is_deleted, created_at, updated_at FROM roles_permission
`

func (q *Queries) GetAllRolePermissions(ctx context.Context) ([]RolesPermission, error) {
	rows, err := q.db.Query(ctx, getAllRolePermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RolesPermission{}
	for rows.Next() {
		var i RolesPermission
		if err := rows.Scan(
			&i.ID,
			&i.RoleID,
			&i.PermissionID,
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

const getRolePermission = `-- name: GetRolePermission :one
SELECT id, role_id, permission_id, is_deleted, created_at, updated_at FROM roles_permission WHERE id = $1
`

func (q *Queries) GetRolePermission(ctx context.Context, id int32) (RolesPermission, error) {
	row := q.db.QueryRow(ctx, getRolePermission, id)
	var i RolesPermission
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.PermissionID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateRolePermission = `-- name: UpdateRolePermission :one
UPDATE roles_permission
SET
    role_id = $2,
    permission_id = $3,
    is_deleted = $4
WHERE id = $1
RETURNING id, role_id, permission_id, is_deleted, created_at, updated_at
`

type UpdateRolePermissionParams struct {
	ID           int32       `json:"id"`
	RoleID       int32       `json:"role_id"`
	PermissionID int32       `json:"permission_id"`
	IsDeleted    pgtype.Bool `json:"is_deleted"`
}

func (q *Queries) UpdateRolePermission(ctx context.Context, arg UpdateRolePermissionParams) (RolesPermission, error) {
	row := q.db.QueryRow(ctx, updateRolePermission,
		arg.ID,
		arg.RoleID,
		arg.PermissionID,
		arg.IsDeleted,
	)
	var i RolesPermission
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.PermissionID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}