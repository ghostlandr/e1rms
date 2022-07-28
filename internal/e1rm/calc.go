package e1rm

// Percentages, as per the chart on page 21 of The Bridge eBook: http://www.barbellmedicine.com/the-bridge/

type percentages map[float64]map[int16]float64

func getPercentages() percentages {
	return percentages{
		10: {
			1:  100,
			2:  96,
			3:  92,
			4:  89,
			5:  86,
			7:  81,
			6:  84,
			8:  79,
			9:  76,
			10: 74,
		},
		9.5: {
			1:  98,
			2:  94,
			3:  91,
			4:  88,
			5:  85,
			6:  82,
			7:  80,
			8:  77,
			9:  75,
			10: 72,
		},
		9: {
			1:  96,
			2:  92,
			3:  89,
			4:  86,
			5:  84,
			6:  81,
			7:  79,
			8:  76,
			9:  74,
			10: 71,
		},
		8.5: {
			1:  94,
			2:  91,
			3:  88,
			4:  85,
			5:  82,
			6:  80,
			7:  77,
			8:  75,
			9:  72,
			10: 69,
		},
		8: {
			1:  92,
			2:  89,
			3:  86,
			4:  84,
			5:  81,
			6:  79,
			7:  76,
			8:  74,
			9:  71,
			10: 68,
		},
		7.5: {
			1:  91,
			2:  88,
			3:  85,
			4:  82,
			5:  80,
			6:  77,
			7:  75,
			8:  72,
			9:  69,
			10: 67,
		},
		7: {
			1:  89,
			2:  86,
			3:  84,
			4:  81,
			5:  79,
			6:  76,
			7:  74,
			8:  71,
			9:  68,
			10: 65,
		},
		6.5: {
			1:  88,
			2:  85,
			3:  82,
			4:  80,
			5:  77,
			6:  75,
			7:  72,
			8:  69,
			9:  67,
			10: 64,
		},
	}
}

func lookupPercentage(rpe float64, reps int16) float64 {
	return getPercentages()[rpe][reps]
}

/**
 * Given a weight, an RPE, and a number of reps, calculate the estimated 1RM for that lift.
 *
 * @param {number} totalWeight The amount of pounds lifted, including the bar, for the set.
 * @param {number} rpe The number representing the RPE (Rate of Perceived Exertion) of the set you just performed.
 * @param {number} reps The number of repetitions performed in this set.
 * @return The estimated 1RM as per calculations from The Bridge eBook
 */
func CalculateE1RM(
	totalWeight float64,
	rpe float64,
	reps int16,
) float64 {
	if totalWeight == 0 || rpe == 1 || reps == 0 {
		return 0
	}
	percentageFromChart := lookupPercentage(rpe, reps)
	return (totalWeight * 100) / percentageFromChart
}
