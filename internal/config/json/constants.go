package json

import "time"

const (
	// адрес сервера по умолчанию.
	defaultAddr = "127.0.0.1:8080"
	// интервал сбора метрик по умолчанию.
	defaultReportInterval = time.Second * 10
	// интервал отправки метрик по умолчанию.
	defaultPoolInterval = time.Second * 2
	// статус восстановления метрик из файлового хранилища по умолчанию.
	defaultRestore = "true"
	// интервал сохранения метрик в файловое хранилище по умолчанию.
	defaultStoreInterval = time.Second * 300
	// путь к файлу файлового хранилища по умолчанию.
	defaultStoreFile = "/tmp/devops-metrics-db.json"
	// ключ хэш-функции по умолчанию.
	defaultKey = ""
	// количество одновременно исходящих запросов по умолчанию.
	defaultRateLimit = 1
)
