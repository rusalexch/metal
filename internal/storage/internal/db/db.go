package db

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rusalexch/metal/internal/app"
)

type dbStorage struct {
	pool *pgxpool.Pool
	sync.Mutex
}

type dbCounter struct {
	ID    string `db:"id"`
	Delta int64  `db:"delta"`
}

type dbGauge struct {
	ID    string  `db:"id"`
	Value float64 `db:"value"`
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
	db.Lock()
	defer db.Unlock()
	switch m.Type {
	case app.Counter:
		return db.saveCounter(m.ID, *m.Delta)
	case app.Gauge:
		return db.saveGauge(m.ID, *m.Value)
	default:
		return app.ErrIncorrectType
	}
}

func (db *dbStorage) AddList(m []app.Metrics) error {
	db.Lock()
	defer db.Unlock()
	ctx := context.Background()
	tx, err := db.pool.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Prepare(ctx, "gaugeInsert", insertGaugeSQL)
	if err != nil {
		return err
	}
	_, err = tx.Prepare(ctx, "counterInsert", insertCounterSQL)
	if err != nil {
		return err
	}
	for _, v := range m {
		switch v.Type {
		case app.Counter:
			{
				if _, err = tx.Exec(ctx, "counterInsert", v.ID, *v.Delta); err != nil {
					if err = tx.Rollback(ctx); err != nil {
						return err
					}
					return err
				}
			}
		case app.Gauge:
			{
				if _, err = tx.Exec(ctx, "gaugeInsert", v.ID, *v.Value); err != nil {
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

// Get получение метрики name с типом mType
func (db *dbStorage) Get(name string, mType app.MetricType) (app.Metrics, error) {
	db.Lock()
	defer db.Unlock()
	switch mType {
	case app.Counter:
		{
			counter, err := db.findCounter(name)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return app.Metrics{}, app.ErrNotFound
				}
				return app.Metrics{}, err
			}
			return app.Metrics{
				ID:    counter.ID,
				Type:  app.Counter,
				Delta: &counter.Delta,
			}, nil
		}
	case app.Gauge:
		{
			gauge, err := db.findGauge(name)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return app.Metrics{}, app.ErrNotFound
				}
				return app.Metrics{}, err
			}
			return app.Metrics{
				ID:    gauge.ID,
				Type:  app.Gauge,
				Value: &gauge.Value,
			}, nil
		}
	default:
		return app.Metrics{}, app.ErrIncorrectType
	}
}

// List получение списка всех метрик
func (db *dbStorage) List() ([]app.Metrics, error) {
	db.Lock()
	defer db.Unlock()
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
	_, err := db.pool.Exec(context.Background(), insertCounterSQL, name, delta)
	return err
}

// saveGauge сохранение метрики типа gauge
func (db *dbStorage) saveGauge(name string, value float64) error {
	_, err := db.pool.Exec(context.Background(), insertGaugeSQL, name, value)
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
