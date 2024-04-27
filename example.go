package example

import "log/slog"

func Exec(log *slog.Logger) {
	log.Info("testing log",
		"foo", "bar",
		"fizz", "buzz")
}
