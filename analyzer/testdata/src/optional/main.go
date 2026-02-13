package optional

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// FAIL
	l, _ := zap.NewProduction()
	sugar := l.Sugar()

	slog.Info("user password is invalid") // want `\[sensitive\].*`
	slog.Info("user admin skygreenxd")    // want `\[sensitive\].*` `\[sensitive\].*` `\[sensitive\].*`

	ctx := context.Background()
	slog.InfoContext(ctx, "skygreenxd")                  // want `\[sensitive\].*`
	slog.Log(ctx, slog.LevelInfo, "skygreenxd the best") // want `\[sensitive\].*`

	password := "sadasdsa"
	apiKey := "shfjsdhf14814"
	adminVal := "hdsjfhdsjfdsh"
	slog.Info("user test: " + password) // want `\[sensitive\].*` `\[sensitive\].*`
	slog.Debug("test=" + apiKey)        // want `\[sensitive\].*`
	slog.Info("admin: " + adminVal)     // want `\[sensitive\].*`
	l.Info("тут тесты хех")             // want `\[sensitive\].*`

	// OK
	slog.Info("connection established")
	slog.Info("request completed")
	l.Info("server started on port 8080")
	sugar.Infow("database connected", "latency", 42)
}
