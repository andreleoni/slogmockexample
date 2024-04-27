package example

import (
	"context"
	"example/mocks"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExec(t *testing.T) {
	loghandler := mocks.NewLogHandler(t)

	loghandler.EXPECT().Enabled(context.Background(), slog.LevelInfo).Return(true)

	loghandler.EXPECT().Handle(context.Background(),
		mock.MatchedBy(func(r slog.Record) bool {
			assert.Equal(t, "testing log", r.Message)

			r.Attrs(func(attr slog.Attr) bool {
				if attr.Key == "foo" {
					assert.Equal(t, "bar", attr.Value.String())
				}

				if attr.Key == "fizz" {
					assert.Equal(t, "buzz", attr.Value.String())
				}

				return true
			})

			return true
		})).Return(nil)

	logger := slog.New(loghandler)

	Exec(logger)
}
