package e1rm

import (
	"fmt"
	"strconv"
)

var acceptableRpe = []float64{6.5, 7, 7.5, 8, 8.5, 9, 9.5, 10}

func isRpeInRange(rpe float64) bool {
	for _, goodRpe := range acceptableRpe {
		if rpe == goodRpe {
			return true
		}
	}
	return false
}

func CalculateE1RMFromStrings(totalWeight, rpe, reps string) (float64, error) {
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

	return CalculateE1RM(totalWeightF, rpeF, int16(repsI)), nil
}
