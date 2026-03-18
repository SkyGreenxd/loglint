package dynamic

import "log/slog"

func main() {
	dynamic := "value_with_symbols!!!"

	// Полностью динамическая строка должна пропускаться,
	// чтобы избежать ложных срабатываний на placeholder.
	slog.Info(dynamic)

	// При этом конкатенация со статической частью должна анализироваться.
	token := "abc123"
	slog.Info("token: " + token) // want `\[symbols\].*` `\[sensitive\].*`
}
