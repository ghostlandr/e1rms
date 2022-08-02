package e1rm_model

import (
	"context"
	"e1rms/internal/e1rm"
	e1rm_calc "e1rms/internal/e1rm/calc"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type DB interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
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
		fmt.Printf("Error doing hello world: %v\n", err)
		return err
	}
	fmt.Println(greeting)
	return nil
}

func (m *e1rmModel) ListE1RMs(ctx context.Context) {}

func (m *e1rmModel) ProvisionTables(ctx context.Context) {
	_, err := m.db.Exec(ctx, createTable)
	if err != nil {
		fmt.Printf("Error when creating table: %v\n", err)
		return
	}
	fmt.Printf("Table created successfully?\n")
}
