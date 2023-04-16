package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rusalexch/metal/internal/app"
)

type dbStorage struct {
	pool *pgxpool.Pool
}

type dbCounter struct {
	ID    string
	Delta int64
}

type dbGauge struct {
	ID    string
	Value float64
}

// New конструктор хранилища БД
func New(URL string) *dbStorage {
	pool, err := pgxpool.New(context.Background(), URL)
	if err != nil {
		log.Panic(err)
	}
	db := &dbStorage{
		pool: pool,
	}
	err = db.init()
	if err != nil {
		log.Panic(err)
	}

	return db
}

// Add добавление новой метрики
func (db *dbStorage) Add(m app.Metrics) error {
	switch m.Type {
	case app.Counter:
		return db.saveCounter(m.ID, *m.Delta)
	case app.Gauge:
		return db.saveGauge(m.ID, *m.Value)
	default:
		return app.ErrNotFound
	}
}

// Get получение метрики name с типом mType
func (db *dbStorage) Get(name string, mType app.MetricType) (*app.Metrics, error) {
	switch mType {
	case app.Counter:
		{
			counter, err := db.findCounter(name)
			if err != nil {
				return nil, err
			}
			m := &app.Metrics{
				ID:    counter.ID,
				Type:  app.Counter,
				Delta: &counter.Delta,
			}
			return m, nil
		}
	case app.Gauge:
		{
			gauge, err := db.findGauge(name)
			if err != nil {
				return nil, err
			}
			m := &app.Metrics{
				ID:    gauge.ID,
				Type:  app.Gauge,
				Value: &gauge.Value,
			}
			return m, nil
		}
	default:
		return nil, app.ErrNotFound
	}
}

// List получение списка всех метрик
func (db *dbStorage) List() ([]app.Metrics, error) {
	counters, err := db.listCounter()
	if err != nil {
		return nil, err
	}
	gauges, err := db.listGauge()
	if err != nil {
		return nil, err
	}
	metrics := make([]app.Metrics, 0, len(counters)+len(gauges))
	for _, c := range counters {
		m := app.Metrics{
			ID:    c.ID,
			Type:  app.Counter,
			Delta: &c.Delta,
		}
		metrics = append(metrics, m)
	}
	for _, g := range gauges {
		m := app.Metrics{
			ID:    g.ID,
			Type:  app.Gauge,
			Value: &g.Value,
		}
		metrics = append(metrics, m)
	}

	return metrics, nil
}

// Ping проверка работы БД
func (db *dbStorage) Ping() error {
	return db.pool.Ping(context.Background())
}

// Close закрыть подключение к БД
func (db *dbStorage) Close() {
	db.pool.Close()
}

// init инициализация БД, создание таблиц если их нет
func (db *dbStorage) init() error {
	ctx := context.Background()
	_, err := db.pool.Exec(ctx, createGaugeTableSQL)
	if err != nil {
		return err
	}
	_, err = db.pool.Exec(ctx, crteateCounterTableSQL)
	if err != nil {
		return err
	}

	return nil
}

// saveCounter сохранение метрики типа counter
func (db *dbStorage) saveCounter(name string, delta int64) error {
	counter, _ := db.findCounter(name)
	if counter == nil {
		_, err := db.pool.Exec(context.Background(), insertCounterSQL, name, delta)
		return err
	}
	_, err := db.pool.Exec(context.Background(), updateCounterSQL, delta+counter.Delta, name)
	return err
}

// saveGauge сохранение метрики типа gauge
func (db *dbStorage) saveGauge(name string, value float64) error {
	gauge, _ := db.findGauge(name)
	if gauge == nil {
		_, err := db.pool.Exec(context.Background(), insertGuageSQL, name, value)
		return err
	}
	_, err := db.pool.Exec(context.Background(), updateGuageSQL, value, name)
	return err
}

// findCounter поиск метрики типа counter по идентификатору name
func (db *dbStorage) findCounter(name string) (*dbCounter, error) {
	var counter dbCounter
	row := db.pool.QueryRow(context.Background(), findCounterSQL, name)
	err := row.Scan(&counter.ID, &counter.Delta)
	if err != nil {
		return nil, err
	}

	return &counter, nil
}

// findGauge поиск метрики типа gauge по идентификатору name
func (db *dbStorage) findGauge(name string) (*dbGauge, error) {
	var gauge dbGauge
	row := db.pool.QueryRow(context.Background(), findGaugeSQL, name)
	err := row.Scan(&gauge.ID, &gauge.Value)
	if err != nil {
		return nil, err
	}

	return &gauge, nil
}

// listCounter получение всех метрик типа counter
func (db *dbStorage) listCounter() ([]dbCounter, error) {
	rows, err := db.pool.Query(context.Background(), listCounterSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	counters := make([]dbCounter, 0)

	for rows.Next() {
		var counter dbCounter
		err = rows.Scan(&counter.ID, &counter.Delta)
		if err != nil {
			return nil, err
		}
		counters = append(counters, counter)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return counters, nil
}

// listGauge получение всех метрик типа gauge
func (db *dbStorage) listGauge() ([]dbGauge, error) {
	rows, err := db.pool.Query(context.Background(), listGaugeSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	gauges := make([]dbGauge, 0)

	for rows.Next() {
		var gauge dbGauge
		err = rows.Scan(&gauge.ID, &gauge.Value)
		if err != nil {
			return nil, err
		}
		gauges = append(gauges, gauge)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return gauges, nil
}
