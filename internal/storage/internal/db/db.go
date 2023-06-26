package db

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rusalexch/metal/internal/app"
)

// dbStorage - структура модуля хранилища PostgreSQL
type dbStorage struct {
	pool *pgxpool.Pool
	sync.Mutex
}

// dbCounter - структура метрики типа counter
type dbCounter struct {
	ID    string `db:"id"`
	Delta int64  `db:"delta"`
}

// dbGauge - структура метрики типа gauge
type dbGauge struct {
	ID    string  `db:"id"`
	Value float64 `db:"value"`
}

// New - конструктор хранилища БД PostgreSQL
func New(URL string) *dbStorage {
	pool, err := pgxpool.New(context.Background(), URL)
	if err != nil {
		log.Panic(err)
	}
	db := &dbStorage{
		pool: pool,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	err = db.init(ctx)
	if err != nil {
		log.Panic(err)
	}

	return db
}

// Add - метод добавления/обновления метрики
func (db *dbStorage) Add(ctx context.Context, m app.Metrics) error {
	db.Lock()
	defer db.Unlock()
	switch m.Type {
	case app.Counter:
		return db.saveCounter(ctx, m.ID, *m.Delta)
	case app.Gauge:
		return db.saveGauge(ctx, m.ID, *m.Value)
	default:
		return app.ErrIncorrectType
	}
}

// AddList - метода добавления/обновления списка метрик
func (db *dbStorage) AddList(ctx context.Context, m []app.Metrics) error {
	db.Lock()
	defer db.Unlock()
	tx, err := db.pool.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return err
	}
	stmtGauge, err := tx.Prepare(ctx, "gaugeInsert", insertGaugeSQL)
	if err != nil {
		return err
	}
	stmtCounter, err := tx.Prepare(ctx, "counterInsert", insertCounterSQL)
	if err != nil {
		return err
	}
	for _, v := range m {
		switch v.Type {
		case app.Counter:
			{
				if _, err = tx.Exec(ctx, stmtCounter.Name, v.ID, *v.Delta); err != nil {
					if err = tx.Rollback(ctx); err != nil {
						return err
					}
					return err
				}
			}
		case app.Gauge:
			{
				if _, err = tx.Exec(ctx, stmtGauge.Name, v.ID, *v.Value); err != nil {
					if err = tx.Rollback(ctx); err != nil {
						return err
					}
					return err
				}
			}
		default:
			log.Println("invalid metric type")
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// Get - метод получение метрики name с типом mType
func (db *dbStorage) Get(ctx context.Context, name string, mType app.MetricType) (app.Metrics, error) {
	db.Lock()
	defer db.Unlock()
	switch mType {
	case app.Counter:
		{
			counter, err := db.findCounter(ctx, name)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return app.Metrics{}, app.ErrNotFound
				}
				return app.Metrics{}, err
			}
			return app.AsCounter(counter.Delta, counter.ID), nil
		}
	case app.Gauge:
		{
			gauge, err := db.findGauge(ctx, name)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return app.Metrics{}, app.ErrNotFound
				}
				return app.Metrics{}, err
			}
			return app.AsGauge(gauge.Value, gauge.ID), nil
		}
	default:
		return app.Metrics{}, app.ErrIncorrectType
	}
}

// List - метод получение списка всех метрик
func (db *dbStorage) List(ctx context.Context) ([]app.Metrics, error) {
	db.Lock()
	defer db.Unlock()
	counters, err := db.listCounter(ctx)
	if err != nil {
		return nil, err
	}
	gauges, err := db.listGauge(ctx)
	if err != nil {
		return nil, err
	}
	metrics := make([]app.Metrics, 0, len(counters)+len(gauges))
	for _, c := range counters {
		m := app.AsCounter(c.Delta, c.ID)
		metrics = append(metrics, m)
	}
	for _, g := range gauges {
		m := app.AsGauge(g.Value, g.ID)
		metrics = append(metrics, m)
	}

	return metrics, nil
}

// Ping - метод проверка работы БД
func (db *dbStorage) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

// Close - метод окончания сессии БД
func (db *dbStorage) Close() {
	db.pool.Close()
}

// init - инициализация БД, создание таблиц если их нет
func (db *dbStorage) init(ctx context.Context) error {
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

// saveCounter - сохранение метрики типа counter
func (db *dbStorage) saveCounter(ctx context.Context, name string, delta int64) error {
	_, err := db.pool.Exec(ctx, insertCounterSQL, name, delta)
	return err
}

// saveGauge - сохранение метрики типа gauge
func (db *dbStorage) saveGauge(ctx context.Context, name string, value float64) error {
	_, err := db.pool.Exec(ctx, insertGaugeSQL, name, value)
	return err
}

// findCounter - поиск метрики типа counter по идентификатору name
func (db *dbStorage) findCounter(ctx context.Context, name string) (dbCounter, error) {
	var counter dbCounter
	row := db.pool.QueryRow(ctx, findCounterSQL, name)
	err := row.Scan(&counter.ID, &counter.Delta)
	if err != nil {
		return counter, err
	}

	return counter, nil
}

// findGauge - поиск метрики типа gauge по идентификатору name
func (db *dbStorage) findGauge(ctx context.Context, name string) (dbGauge, error) {
	var gauge dbGauge
	row := db.pool.QueryRow(ctx, findGaugeSQL, name)
	err := row.Scan(&gauge.ID, &gauge.Value)
	if err != nil {
		return gauge, err
	}

	return gauge, nil
}

// listCounter - получение всех метрик типа counter
func (db *dbStorage) listCounter(ctx context.Context) ([]dbCounter, error) {
	rows, err := db.pool.Query(ctx, listCounterSQL)
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

// listGauge - получение всех метрик типа gauge
func (db *dbStorage) listGauge(ctx context.Context) ([]dbGauge, error) {
	rows, err := db.pool.Query(ctx, listGaugeSQL)
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
