// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: costCalculation.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCostCalculation = `-- name: CreateCostCalculation :one
INSERT INTO cost_calculation (
    parkcost,
    assetcost,
    loancost,
    tax_cost,
    asset_value_cost,
    revenue,
    total_cost,
    result,
    result_perc,
    saving_revenue_amount,
    saving_revenue_perc,
    yield_diff,
    is_deleted
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13
) RETURNING id, parkcost, assetcost, loancost, tax_cost, asset_value_cost, revenue, total_cost, result, result_perc, saving_revenue_amount, saving_revenue_perc, yield_diff, is_deleted, created_at, updated_at
`

type CreateCostCalculationParams struct {
	Parkcost            pgtype.Int4 `json:"parkcost"`
	Assetcost           pgtype.Int4 `json:"assetcost"`
	Loancost            pgtype.Int4 `json:"loancost"`
	TaxCost             pgtype.Int4 `json:"tax_cost"`
	AssetValueCost      pgtype.Int4 `json:"asset_value_cost"`
	Revenue             pgtype.Int4 `json:"revenue"`
	TotalCost           pgtype.Int4 `json:"total_cost"`
	Result              pgtype.Int4 `json:"result"`
	ResultPerc          pgtype.Int4 `json:"result_perc"`
	SavingRevenueAmount pgtype.Int4 `json:"saving_revenue_amount"`
	SavingRevenuePerc   pgtype.Int4 `json:"saving_revenue_perc"`
	YieldDiff           pgtype.Int4 `json:"yield_diff"`
	IsDeleted           pgtype.Bool `json:"is_deleted"`
}

func (q *Queries) CreateCostCalculation(ctx context.Context, arg CreateCostCalculationParams) (CostCalculation, error) {
	row := q.db.QueryRow(ctx, createCostCalculation,
		arg.Parkcost,
		arg.Assetcost,
		arg.Loancost,
		arg.TaxCost,
		arg.AssetValueCost,
		arg.Revenue,
		arg.TotalCost,
		arg.Result,
		arg.ResultPerc,
		arg.SavingRevenueAmount,
		arg.SavingRevenuePerc,
		arg.YieldDiff,
		arg.IsDeleted,
	)
	var i CostCalculation
	err := row.Scan(
		&i.ID,
		&i.Parkcost,
		&i.Assetcost,
		&i.Loancost,
		&i.TaxCost,
		&i.AssetValueCost,
		&i.Revenue,
		&i.TotalCost,
		&i.Result,
		&i.ResultPerc,
		&i.SavingRevenueAmount,
		&i.SavingRevenuePerc,
		&i.YieldDiff,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllCostCalculation = `-- name: GetAllCostCalculation :many
SELECT id, parkcost, assetcost, loancost, tax_cost, asset_value_cost, revenue, total_cost, result, result_perc, saving_revenue_amount, saving_revenue_perc, yield_diff, is_deleted, created_at, updated_at FROM cost_calculation
`

func (q *Queries) GetAllCostCalculation(ctx context.Context) ([]CostCalculation, error) {
	rows, err := q.db.Query(ctx, getAllCostCalculation)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CostCalculation{}
	for rows.Next() {
		var i CostCalculation
		if err := rows.Scan(
			&i.ID,
			&i.Parkcost,
			&i.Assetcost,
			&i.Loancost,
			&i.TaxCost,
			&i.AssetValueCost,
			&i.Revenue,
			&i.TotalCost,
			&i.Result,
			&i.ResultPerc,
			&i.SavingRevenueAmount,
			&i.SavingRevenuePerc,
			&i.YieldDiff,
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
