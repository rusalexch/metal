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

var insertGuageSQL = `INSERT INTO guages VALUES($1, $2)`

var insertCounterSQL = `INSERT INTO counters VALUES($1, $2)`

var updateGuageSQL = `UPDATE guages SET value = $1 WHERE id = $2`

var updateCounterSQL = `UPDATE counters SET delta = $1 WHERE id = $2`

var findGaugeSQL = `SELECT * FROM guages WHERE id = $1 LIMIT 1`

var findCounterSQL = `SELECT * FROM counters WHERE id = $1 LIMIT 1`

var listGaugeSQL = `SELECT * FROM gauges;`

var listCounterSQL = `SELECT * FROM counters;`
