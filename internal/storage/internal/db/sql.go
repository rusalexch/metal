package db

// createGaugeTableSQL - запрос на создание таблицы метрик типа gauge.
var createGaugeTableSQL = `
	CREATE TABLE IF NOT EXISTS gauges (
	id varchar PRIMARY KEY,
	value double precision
	)`

// crteateCounterTableSQL - запрос на создание таблицы метрик типа counter.
var crteateCounterTableSQL = `
	CREATE TABLE IF NOT EXISTS counters (
	id varchar PRIMARY KEY,
	delta bigint
	)`

// insertGaugeSQL - запрос на добавление новой метрики типа gauge или обновление ее значения в случае её отсутствия.
var insertGaugeSQL = `INSERT INTO gauges VALUES($1, $2) ON CONFLICT (id) DO UPDATE SET value = $2`

// insertCounterSQL - запрос на добавление новой метрики типа counter или обновление ее значения в случае её отсутствия.
var insertCounterSQL = `INSERT INTO counters AS c VALUES($1, $2) ON CONFLICT (id) DO UPDATE SET delta = c.delta + $2`

// findGaugeSQL - запрос на поиск метрики типа guage по идентификатору (имени).
var findGaugeSQL = `SELECT * FROM gauges WHERE id = $1 LIMIT 1`

// findCounterSQL - запрос на поиск метрики типа counter по идентификатору (имени).
var findCounterSQL = `SELECT * FROM counters WHERE id = $1 LIMIT 1`

// listGaugeSQL - запрос на получение всех метрик типа gauge.
var listGaugeSQL = `SELECT * FROM gauges;`

// listCounterSQL - запрос на получение всех метрик типа counter.
var listCounterSQL = `SELECT * FROM counters;`
