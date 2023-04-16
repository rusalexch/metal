package filestorage

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/rusalexch/metal/internal/app"
)

type fileStorage struct {
	file    *os.File
	encoder *json.Encoder
	decoder *json.Decoder
}

type store struct {
	Counters map[string]int64   `json:"counters"`
	Gauges   map[string]float64 `json:"gauges"`
}

func New(file string, restore bool) *fileStorage {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(f)
	decoder := json.NewDecoder(f)

	fs := &fileStorage{
		file:    f,
		encoder: encoder,
		decoder: decoder,
	}

	if !restore {
		fs.clear()
	}

	return fs
}

func (fs *fileStorage) Add(m app.Metrics) error {
	if !app.IsMetricType(m.Type) {
		return app.ErrNotFound
	}
	st, err := fs.upload()
	if err != nil {
		return err
	}
	if m.Type == app.Counter {
		delta, isExist := st.Counters[m.ID]
		if isExist {
			st.Counters[m.ID] = delta + *m.Delta
		} else {
			st.Counters[m.ID] = *m.Delta
		}
	}
	if m.Type == app.Gauge {
		st.Gauges[m.ID] = *m.Value
	}

	err = fs.save(st)
	if err != nil {
		return err
	}

	return nil
}

// Get получение метрики с именем name и типом mType
func (fs *fileStorage) Get(name string, mType app.MetricType) (*app.Metrics, error) {
	if !app.IsMetricType(mType) {
		return nil, app.ErrNotFound
	}
	st, err := fs.upload()
	if err != nil {
		return nil, err
	}
	if mType == app.Counter {
		m, isExist := st.Counters[name]
		if isExist {
			return &app.Metrics{
				ID:    name,
				Type:  mType,
				Delta: &m,
			}, nil
		}
	}
	if mType == app.Gauge {
		m, isExist := st.Gauges[name]
		if isExist {
			return &app.Metrics{
				ID:    name,
				Type:  mType,
				Value: &m,
			}, nil
		}
	}

	return nil, app.ErrNotFound
}

// List получения всего списка метрик
func (fs *fileStorage) List() ([]app.Metrics, error) {
	st, err := fs.upload()
	if err != nil {
		return nil, err
	}

	return storeToMetric(st), nil
}

// Ping заглушка
func (fs *fileStorage) Ping() error {
	return nil
}

// Close закрыть файл хранилища
func (fs *fileStorage) Close() {
	fs.file.Close()
}

// save сохранить метрики в файл
func (fs *fileStorage) save(s store) error {
	fs.clear()
	err := fs.encoder.Encode(s)
	defer fs.seekStart()
	if err != nil {
		return err
	}

	return nil
}

// upload выгрузить метрики из файла
func (fs *fileStorage) upload() (store, error) {
	var st store
	err := fs.decoder.Decode(&st)
	defer fs.seekStart()
	if err != nil && !errors.Is(err, io.EOF) {
		return emptyStore(), err
	}

	if st.Counters == nil {
		st.Counters = map[string]int64{}
	}
	if st.Gauges == nil {
		st.Gauges = map[string]float64{}
	}
	return st, nil
}

// clear очистить файл
func (fs *fileStorage) clear() {
	fs.seekStart()
	fs.file.Truncate(0)
}

// seekStart возврат указателя в начало файла
func (fs *fileStorage) seekStart() {
	fs.file.Seek(0, io.SeekStart)
}

// storeToMetric преобразование слайса метрик в структуру файла
func storeToMetric(st store) []app.Metrics {
	m := make([]app.Metrics, 0, len(st.Counters)+len(st.Gauges))
	for name, delta := range st.Counters {
		d := delta
		m = append(m, app.Metrics{
			ID:    name,
			Type:  app.Counter,
			Delta: &d,
		})
	}
	for name, value := range st.Gauges {
		v := value
		m = append(m, app.Metrics{
			ID:    name,
			Type:  app.Gauge,
			Value: &v,
		})
	}

	return m
}

// emptyStore генерация пустой структуры для файла
func emptyStore() store {
	return store{
		Counters: map[string]int64{},
		Gauges:   map[string]float64{},
	}
}
