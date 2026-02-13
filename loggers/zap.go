package loggers

const (
	ZapName = "zap"
	ZapPath = "go.uber.org/zap"
)

func init() {
	Register(NewZapLogger())
}

type ZapLogger struct {
	BaseLogger
}

func NewZapLogger() Logger {
	methods := []string{
		"Debug",
		"Debugf",
		"Debugw",
		"Info",
		"Infof",
		"Infow",
		"Warn",
		"Warnf",
		"Warnw",
		"Error",
		"Errorf",
		"Errorw",
		"DPanic",
		"DPanicf",
		"DPanicw",
		"Panic",
		"Panicf",
		"Panicw",
		"Fatal",
		"Fatalf",
		"Fatalw",
	}

	return &ZapLogger{
		BaseLogger: NewBaseLogger(ZapName, ZapPath, methods),
	}
}
