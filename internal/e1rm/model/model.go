package e1rm_model

import (
	"context"
	"e1rms/internal/e1rm"
	e1rm_calc "e1rms/internal/e1rm/calc"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Internal calculation struct purely for adding the db tag
type e1RMCalculation struct {
	TotalWeight float64
	RPE         float64
	Reps        int16
	E1RM        float64 `db:"estimated_1rm"`
	CreatedAt   time.Time
}

type DB interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type e1rmModel struct {
	db DB
}

func New(db DB) e1rm.E1RMModel {
	return &e1rmModel{db}
}

func (m *e1rmModel) SaveE1RM(ctx context.Context, e1rm e1rm_calc.E1RMCalculation) error {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		fmt.Printf("Error starting transaction: %v\n", err)
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(ctx, `INSERT INTO e1rms VALUES ($1, $2, $3, $4, $5)`, e1rm.TotalWeight, e1rm.RPE, e1rm.Reps, e1rm.E1RM, time.Now().UTC())
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *e1rmModel) ListE1RMs(ctx context.Context) ([]*e1rm_calc.E1RMCalculation, error) {
	var e1rms []*e1RMCalculation
	err := pgxscan.Select(ctx, m.db, &e1rms, `SELECT total_weight, rpe, reps, estimated_1rm, created_at FROM e1rms ORDER BY created_at DESC`)

	if err != nil {
		fmt.Printf("Error selecting with pgxscan: %v\n", err)
		return nil, err
	}

	return modelE1rmsToCalcE1rms(e1rms), nil
}

func modelE1rmsToCalcE1rms(in []*e1RMCalculation) []*e1rm_calc.E1RMCalculation {
	out := make([]*e1rm_calc.E1RMCalculation, len(in))
	for idx, calc := range in {
		out[idx] = &e1rm_calc.E1RMCalculation{
			TotalWeight: calc.TotalWeight,
			RPE:         calc.RPE,
			Reps:        calc.Reps,
			E1RM:        calc.E1RM,
			CreatedAt:   calc.CreatedAt.String(),
		}
	}
	return out
}

func (m *e1rmModel) ProvisionTables(ctx context.Context) {
	_, err := m.db.Exec(ctx, alterTable)
	if err != nil {
		fmt.Printf("Error when creating table: %v\n", err)
		return
	}
	fmt.Printf("Table created successfully?\n")
}
