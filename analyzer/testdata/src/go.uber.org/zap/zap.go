// Пакет zap предоставляет минимальную заглушку библиотеки Uber Zap.
//
// Это необходимо, потому что инструмент тестирования 'analysistest' запускает
// проверки в изолированном окружении и ищет зависимости только внутри
// папки 'testdata/src', игнорируя системный GOMOD/GOPATH.

package zap

type Logger struct{}

func (l *Logger) Info(msg string, fields ...any)   {}
func (l *Logger) Error(msg string, fields ...any)  {}
func (l *Logger) Warn(msg string, fields ...any)   {}
func (l *Logger) Debug(msg string, fields ...any)  {}
func (l *Logger) DPanic(msg string, fields ...any) {}
func (l *Logger) Panic(msg string, fields ...any)  {}
func (l *Logger) Fatal(msg string, fields ...any)  {}
func (l *Logger) Sync() error                      { return nil }

func NewProduction() (*Logger, error)  { return &Logger{}, nil }
func NewDevelopment() (*Logger, error) { return &Logger{}, nil }

func (l *Logger) Sugar() *SugaredLogger { return &SugaredLogger{} }

type SugaredLogger struct{}

func (s *SugaredLogger) Info(args ...any)                         {}
func (s *SugaredLogger) Infof(template string, args ...any)       {}
func (s *SugaredLogger) Infow(msg string, keysAndValues ...any)   {}
func (s *SugaredLogger) Debugf(template string, args ...any)      {}
func (s *SugaredLogger) Debugw(msg string, keysAndValues ...any)  {}
func (s *SugaredLogger) Warnf(template string, args ...any)       {}
func (s *SugaredLogger) Warnw(msg string, keysAndValues ...any)   {}
func (s *SugaredLogger) Errorf(template string, args ...any)      {}
func (s *SugaredLogger) Errorw(msg string, keysAndValues ...any)  {}
func (s *SugaredLogger) DPanicf(template string, args ...any)     {}
func (s *SugaredLogger) DPanicw(msg string, keysAndValues ...any) {}
func (s *SugaredLogger) Panicf(template string, args ...any)      {}
func (s *SugaredLogger) Panicw(msg string, keysAndValues ...any)  {}
func (s *SugaredLogger) Fatalf(template string, args ...any)      {}
func (s *SugaredLogger) Fatalw(msg string, keysAndValues ...any)  {}
