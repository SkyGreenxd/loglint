package lowercase

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	const ErrLog = "Err Bad message"

	slog.Info("Bad message") // want `\[lowercase\].*`
	slog.Info("good message")
	slog.Error(ErrLog) // want `\[lowercase\].*`

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Bad message") // want `\[lowercase\].*`

	sugar := logger.Sugar()
	sugar.Errorf("Another Bad message")         // want `\[lowercase\].*`
	sugar.Debugf("Bad message in debug")        // want `\[lowercase\].*`
	sugar.Warnw("Bad message in warnw", "k", 1) // want `\[lowercase\].*`
	sugar.Panicf("Bad message panic skygreenxd")           // want `\[lowercase\].*`

	slog.Info("Bad " + "message") // want `\[lowercase\].*`

	slog.Info("1st is a number")
	slog.Info("!bang but ok for lowercase rule")
	slog.Info("  Bad after spaces")

	// OK
	logger.Warn("good message")
	logger.DPanic("Bad message") // want `\[lowercase\].*`
	logger.Fatal("Bad message")  // want `\[lowercase\].*`
}
