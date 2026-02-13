package loggers

const (
	SlogName = "slog"
	SlogPath = "log/slog"
)

func init() {
	Register(NewSlogLogger())
}

type SlogLogger struct {
	BaseLogger
}

func NewSlogLogger() Logger {
	methods := []string{
		"Debug",
		"DebugContext",
		"Info",
		"InfoContext",
		"Warn",
		"WarnContext",
		"Error",
		"ErrorContext",
		"Log",
		"LogAttrs",
	}

	return &SlogLogger{
		BaseLogger: NewBaseLogger(SlogName, SlogPath, methods),
	}
}
