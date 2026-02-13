package loggers

import (
	"sync"

	"github.com/SkyGreenxd/loglint/pkg/e"
)

var (
	globalLoggerRegistry = &Registry{
		loggers: make(map[string]Logger),
		mu:      sync.RWMutex{},
	}
)

type Registry struct {
	loggers map[string]Logger
	mu      sync.RWMutex
}

// Register регистрирует логгер в глобальном реестре
func Register(logger Logger) {
	globalLoggerRegistry.mu.Lock()
	defer globalLoggerRegistry.mu.Unlock()

	if _, exists := globalLoggerRegistry.loggers[logger.Name()]; exists {
		panic(e.Wrap(logger.Name(), ErrLoggerRegistered))
	}

	globalLoggerRegistry.loggers[logger.Name()] = logger
}

// GetRegistry возвращает глобальный реестр логгеров
func GetRegistry() *Registry {
	return globalLoggerRegistry
}

// GetAll возвращает все зарегистрированные логгеры
func (r *Registry) GetAll() []Logger {
	r.mu.RLock()
	defer r.mu.RUnlock()

	loggers := make([]Logger, 0, len(r.loggers))
	for _, logger := range r.loggers {
		loggers = append(loggers, logger)
	}
	return loggers
}

// GetByNames возвращает логгеры по списку имен
func (r *Registry) GetByNames(names []string) []Logger {
	if len(names) == 0 {
		return r.GetAll()
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	loggers := make([]Logger, 0, len(names))
	for _, name := range names {
		if logger, ok := r.loggers[name]; ok {
			loggers = append(loggers, logger)
		}
	}
	return loggers
}
