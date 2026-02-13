package loggers

import "errors"

var (
	ErrLoggerRegistered = errors.New("logger already registered")
)
