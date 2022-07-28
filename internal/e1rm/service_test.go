package e1rm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateE1RMFromStrings(t *testing.T) {
	tests := []struct {
		name          string
		totalWeight   string
		rpe           string
		reps          string
		expectedError string
		want          float64
	}{
		{
			name:          "Should return error if totalWeight can't become a float",
			totalWeight:   "",
			rpe:           "8",
			reps:          "7",
			expectedError: "totalWeight could not be converted to a float: ",
			want:          0,
		},
		{
			name:          "Should return error if rpe can't become a float",
			totalWeight:   "185",
			rpe:           "",
			reps:          "7",
			expectedError: "rpe could not be converted to a float: ",
			want:          0,
		},
		{
			name:          "Should return error if reps can't become an integer",
			totalWeight:   "185",
			rpe:           "8",
			reps:          "",
			expectedError: "reps could not be converted to an integer: ",
			want:          0,
		},
		{
			name:          "Should return error for reps greater than 10",
			totalWeight:   "185",
			rpe:           "8",
			reps:          "15",
			expectedError: "e1rm can't be calculated for reps above 10 (or less than 1): 15",
			want:          0,
		},
		{
			name:          "Should return error for reps less than 1",
			totalWeight:   "185",
			rpe:           "8",
			reps:          "0",
			expectedError: "e1rm can't be calculated for reps above 10 (or less than 1): 0",
			want:          0,
		},
		{
			name:          "Should return error for rpe less than 6.5",
			totalWeight:   "185",
			rpe:           "5",
			reps:          "7",
			expectedError: "e1rm can't be calculated for rpe outside of this range [6.5 7 7.5 8 8.5 9 9.5 10]: 5",
			want:          0,
		},
		{
			name:          "Should return error for rpe greater than 10",
			totalWeight:   "185",
			rpe:           "11",
			reps:          "7",
			expectedError: "e1rm can't be calculated for rpe outside of this range [6.5 7 7.5 8 8.5 9 9.5 10]: 11",
			want:          0,
		},
		{
			name:          "Should return ~243 for 7 reps of 185 @ rpe 8",
			totalWeight:   "185",
			rpe:           "8",
			reps:          "7",
			expectedError: "",
			want:          243.42105263157896,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := CalculateE1RMFromStrings(test.totalWeight, test.rpe, test.reps)
			if err != nil {
				assert.Equal(t, err.Error(), test.expectedError)
			} else if test.expectedError != "" {
				t.Fatalf("didn't produce expected error: %s", test.expectedError)
			}
			assert.Equal(t, actual, test.want)
		})
	}
}
