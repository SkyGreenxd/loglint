package sensitive

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	// FAIL
	l, _ := zap.NewProduction()
	sugar := l.Sugar()

	slog.Info("user password is invalid")      // want `\[sensitive\].*`
	slog.Error("error with pwd reset")         // want `\[sensitive\].*`
	slog.Info("Password changed successfully") // want `\[sensitive\].*`
	slog.Info("user passwd expired")           // want `\[sensitive\].*`

	l.Warn("api_key rotation completed")     // want `\[sensitive\].*`
	sugar.Infof("client secret was rotated") // want `\[sensitive\].*`
	slog.Info("api-key is invalid")          // want `\[sensitive\].*`
	slog.Info("secret_key not found")        // want `\[sensitive\].*`

	slog.Info("auth-token expired")                  // want `\[sensitive\].*`
	l.Info("access_key missing")                     // want `\[sensitive\].*`
	slog.Info("token refresh failed")                // want `\[sensitive\].*`
	sugar.Debugw("token revoked", "token", "abc123") // want `\[sensitive\].*`

	slog.Info("user credentials are wrong") // want `\[sensitive\].*`
	slog.Warn("invalid creds provided")     // want `\[sensitive\].*`

	slog.Info("Authorization: Bearer token") // want `\[sensitive\].*` `\[sensitive\].*`
	slog.Info("Bearer abcDEF123456==")       // want `\[sensitive\].*`
	l.DPanic("Bearer a.b.c==")               // want `\[sensitive\].*`

	slog.Info("session_id: 5f4dcc3b5aa765d61d8327deb882cf99")   // want `\[sensitive\].*`
	slog.Info("hash: 0123456789abcdef0123456789abcdef01234567") // want `\[sensitive\].*`
	l.Error("hash: 5f4dcc3b5aa765d61d8327deb882cf99")           // want `\[sensitive\].*`

	ctx := context.Background()
	slog.InfoContext(ctx, "api_key rotation")                                  // want `\[sensitive\].*`
	slog.Log(ctx, slog.LevelInfo, "secret rotation complete")                  // want `\[sensitive\].*`
	slog.LogAttrs(ctx, slog.LevelWarn, "token expired", slog.String("k", "v")) // want `\[sensitive\].*`

	password := "sadasdsa"
	apiKey := "shfjsdhf14814"
	tokenVal := "hdsjfhdsjfdsh"
	slog.Info("user password: " + password) // want `\[sensitive\].*`
	slog.Debug("api_key=" + apiKey)         // want `\[sensitive\].*`
	slog.Info("token: " + tokenVal)         // want `\[sensitive\].*`

	// OK
	slog.Info("connection established")
	slog.Info("request completed")
	l.Info("server started on port 8080")
	sugar.Infow("database connected", "latency", 42)
}
