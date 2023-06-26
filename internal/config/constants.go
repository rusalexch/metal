package config

import "time"

const (
	defaultAddr           = "127.0.0.1:8080"              // адрес сервера по умолчанию
	defaultReportInterval = time.Second * 10              // интервал сбора метрик по умолчанию
	defaultPoolInterval   = time.Second * 2               // интервал отправки метрик по умолчанию
	defaultRestore        = "true"                        // статус восстановления метрик из файлового хранилища по умолчанию
	defaultStoreInterval  = time.Second * 300             // интервал сохранения метрик в файловое хранилище по умолчанию
	defaultStoreFile      = "/tmp/devops-metrics-db.json" // путь к файлу файлового хранилища по умолчанию
	defaultKey            = ""                            // ключ хэш-функции по умолчанию
	defaultRateLimit      = 1                             // количество одновременно исходящих запросов по умолчанию
)
