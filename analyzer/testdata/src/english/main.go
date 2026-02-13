package english

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// FAIL
	slog.Info("привет мир")                // want `\[english\].*`
	slog.Info("cafe is spelled caf\u00e9") // want `\[english\].*`

	ctx := context.Background()
	slog.DebugContext(ctx, "тест debug")                                    // want `\[english\].*`
	slog.WarnContext(ctx, "ошибка warn")                                    // want `\[english\].*`
	slog.LogAttrs(ctx, slog.LevelInfo, "лог attrs", slog.String("id", "1")) // want `\[english\].*`

	logger, _ := zap.NewProduction()
	logger.Warn("ошибка базы данных") // want `\[english\].*`

	sugar := logger.Sugar()
	sugar.Infof("пользователь %s не найден", "skygreenxd") // want `\[english\].*`
	sugar.Debugw("ошибка debugw", "k", 1)                  // want `\[english\].*`

	// OK: английский в Infow
	sugar.Infow("user logout", "id", 123)
	slog.Error("connection failed")
}
