package e1rm_service

import (
	"context"
	"fmt"
	"strconv"

	"e1rms/internal/e1rm"
	e1rm_calc "e1rms/internal/e1rm/calc"
)

type e1rmService struct {
	model e1rm.E1RMModel
}

func New(model e1rm.E1RMModel) e1rm.E1RMService {
	return &e1rmService{model}
}

var acceptableRpe = []float64{6.5, 7, 7.5, 8, 8.5, 9, 9.5, 10}

func isRpeInRange(rpe float64) bool {
	for _, goodRpe := range acceptableRpe {
		if rpe == goodRpe {
			return true
		}
	}
	return false
}

// CalculateE1RM takes the total weight lifted, the rpe, and the reps performed as strings
// and does the work to convert them to the values that the calculator expects.
func (s *e1rmService) CalculateE1RM(ctx context.Context, totalWeight, rpe, reps, lift string) (float64, error) {
	totalWeightF, err := strconv.ParseFloat(totalWeight, 64)
	if err != nil {
		return 0, fmt.Errorf("totalWeight could not be converted to a float: %s", totalWeight)
	}
	rpeF, err := strconv.ParseFloat(rpe, 64)
	if err != nil {
		return 0, fmt.Errorf("rpe could not be converted to a float: %s", rpe)
	}
	repsI, err := strconv.ParseInt(reps, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("reps could not be converted to an integer: %s", reps)
	}

	if repsI > 10 || repsI < 1 {
		// Can't calculate an e1rm for this amount of reps
		return 0, fmt.Errorf("e1rm can't be calculated for reps above 10 (or less than 1): %s", reps)
	}

	if !isRpeInRange(rpeF) {
		// Can't calculate an e1rm for these rpes
		return 0, fmt.Errorf("e1rm can't be calculated for rpe outside of this range %v: %s", acceptableRpe, rpe)
	}

	result := e1rm_calc.CalculateE1RM(totalWeightF, rpeF, int16(repsI))
	result.Lift = lift

	s.model.SaveE1RM(ctx, result)

	return result.E1RM, nil
}

func (s *e1rmService) ListE1RMs(ctx context.Context) ([]*e1rm_calc.E1RMCalculation, error) {
	return s.model.ListE1RMs(ctx)
}
