package e1rm_model

import (
	"context"
	"e1rms/internal/e1rm"
	e1rm_calc "e1rms/internal/e1rm/calc"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type DB interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type e1rmModel struct {
	db DB
}

func New(db DB) e1rm.E1RMModel {
	return &e1rmModel{db}
}

func (m *e1rmModel) SaveE1RM(ctx context.Context, e1rm e1rm_calc.E1RMCalculation) error {
	var greeting string
	err := m.db.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return err
	}
	fmt.Println(greeting)
	return nil
}

func (m *e1rmModel) ListE1RMs(ctx context.Context) {}
