package db

var createGaugeTableSQL = `
	CREATE TABLE IF NOT EXISTS gauges (
	id varchar PRIMARY KEY,
	value double precision
	)`

var crteateCounterTableSQL = `
	CREATE TABLE IF NOT EXISTS counters (
	id varchar PRIMARY KEY,
	delta bigint
	)`

var insertGaugeSQL = `INSERT INTO gauges VALUES($1, $2) ON CONFLICT (id) DO UPDATE SET value = $2`

var insertCounterSQL = `INSERT INTO counters VALUES($1, $2) ON CONFLICT (id) DO UPDATE SET delta = $2`

var findGaugeSQL = `SELECT * FROM gauges WHERE id = $1 LIMIT 1`

var findCounterSQL = `SELECT * FROM counters WHERE id = $1 LIMIT 1`

var listGaugeSQL = `SELECT * FROM gauges;`

var listCounterSQL = `SELECT * FROM counters;`
