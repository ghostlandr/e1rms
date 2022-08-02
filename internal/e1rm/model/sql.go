package e1rm_model

var createTable = `CREATE TABLE e1rms IF NOT EXISTS (
	total_weight NUMERIC(5, 1)
	rpe NUMERIC(5, 1)
	reps integer
	estimated_1rm NUMERIC(5, 1)
)`
