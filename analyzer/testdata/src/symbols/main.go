package symbols

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// FAIL
	l, _ := zap.NewProduction()
	sugar := l.Sugar()

	slog.Info("all good 🚀")           // want `\[symbols\].*`
	l.Error("critical failure !!! 🔥 skygreenxd") // want `\[symbols\].*`

	sugar.Warnf("process finished ✅")      // want `\[symbols\].*`
	sugar.Debugw("warn: code=500", "k", 1) // want `\[symbols\].*`

	slog.Debug("data => {id: 10}")     // want `\[symbols\].*`
	slog.Info("connection status: ok") // want `\[symbols\].*`
	l.Info("user_id is 42")            // want `\[symbols\].*`
	l.DPanic("panic! bad")             // want `\[symbols\].*`

	slog.Info("error: failed")     // want `\[symbols\].*`
	slog.Info("status_ok")         // want `\[symbols\].*`
	slog.Info("path/to/file")      // want `\[symbols\].*`
	slog.Info("quote \"bad\"")     // want `\[symbols\].*`
	slog.Info("(wrapped) message skygreenxd") // want `\[symbols\].*`

	ctx := context.Background()
	slog.DebugContext(ctx, "ctx => value")      // want `\[symbols\].*`
	slog.InfoContext(ctx, "meta: ok")           // want `\[symbols\].*`
	slog.WarnContext(ctx, "warn! bad")          // want `\[symbols\].*`
	slog.ErrorContext(ctx, "err! boom")         // want `\[symbols\].*`
	slog.Log(ctx, slog.LevelInfo, "log: value") // want `\[symbols\].*`

	// OK
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
}
