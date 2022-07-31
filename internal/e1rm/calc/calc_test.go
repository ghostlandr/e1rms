package e1rm_calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_e1rms(t *testing.T) {
	tests := []struct {
		name        string
		totalWeight float64
		rpe         float64
		reps        int16
		want        float64
	}{
		{
			name:        "CalculateE1RM should return zero if passed all zeroes",
			totalWeight: 0,
			rpe:         0,
			reps:        0,
			want:        0,
		},
		{
			name:        "CalculateE1RM should return ~243 for 7 reps of 185 @ rpe 8",
			totalWeight: 185,
			rpe:         8,
			reps:        7,
			want:        243.42105263157896,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CalculateE1RM(test.totalWeight, test.rpe, test.reps)
			assert.Equal(t, test.want, got)
		})
	}
}
