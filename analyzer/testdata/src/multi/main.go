package multi

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	slog.Info("Ошибка подключения! 🔥") // want `\[lowercase\].*` `\[english\].*` `\[symbols\].*`

	logger, _ := zap.NewProduction()
	logger.Info("Starting server on port 8080") // want `\[lowercase\].*`

	sugar := logger.Sugar()

	sugar.Errorf("Пароль неверный 🛑") // want `\[lowercase\].*` `\[english\].*` `\[symbols\].*`

	slog.Info("good message but with emoji 🚀")            // want `\[symbols\].*`
	slog.DebugContext(context.Background(), "bad -> ctx") // want `\[symbols\].*`

	slog.Info("all systems nominal") // OK

	slog.Info("user password is invalid") // want `\[sensitive\].*`
	slog.Error("error with pwd reset")    // want `\[sensitive\].*`

	logger.Warn("api_key rotation done")     // want `\[symbols\].*` `\[sensitive\].*`
	sugar.Infof("client secret was rotated") // want `\[sensitive\].*`

	slog.Info("auth-token expired")   // want `\[symbols\].*` `\[sensitive\].*`
	logger.Info("access_key missing") // want `\[symbols\].*` `\[sensitive\].*`

	slog.Info("Authorization: Bearer my.token.value==")                                             // want `\[lowercase\].*` `\[symbols\].*` `\[sensitive\].*` `\[sensitive\].*`
	slog.Info("session_id: 5f4dcc3b5aa765d61d8327deb882cf99")                                       // want `\[symbols\].*` `\[sensitive\].*`
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Password: abc123!", slog.String("k", "v")) // want `\[lowercase\].*` `\[symbols\].*` `\[sensitive\].*`

	slog.Info("!token expired")           // want `\[symbols\].*` `\[sensitive\].*`
	sugar.Warnw("Token expired", "id", 1) // want `\[lowercase\].*` `\[sensitive\].*`

	slog.Info("1st ошибка подключения") // want `\[english\].*`
	sugar.Debugf("Ошибка debugf")       // want `\[lowercase\].*` `\[english\].*`
}
