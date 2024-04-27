package example

import "log/slog"

type LogHandler interface {
	slog.Handler
}
