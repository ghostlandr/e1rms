package e1rm

import (
	"context"
	"net/http"

	e1rm_calc "e1rms/internal/e1rm/calc"
)

type E1RMService interface {
	CalculateE1RM(context.Context, string, string, string) (float64, error)
}

type E1RMHandler interface {
	ServeE1rmRequest(w http.ResponseWriter, r *http.Request)
}

type E1RMModel interface {
	SaveE1RM(context.Context, e1rm_calc.E1RMCalculation) error
	ListE1RMs(context.Context)
	ProvisionTables(context.Context)
}
