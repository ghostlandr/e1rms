package e1rm_model

var createTable = `CREATE TABLE IF NOT EXISTS e1rms (
	total_weight NUMERIC(5, 1) NOT NULL,
	rpe NUMERIC(5, 1) NOT NULL,
	reps integer NOT NULL,
	estimated_1rm NUMERIC(5, 1) NOT NULL
)`

var alterTableAddCreated = `ALTER TABLE e1rms ADD COLUMN created_at TIMESTAMP`
var alterTableAddLift = `ALTER TABLE e1rms ADD COLUMN lift varchar(50)`
