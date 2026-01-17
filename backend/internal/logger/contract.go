package logger

import "context"

type Logger interface {
	Error(ctx context.Context, event, msg string, err error, meta map[string]interface{})
	Log(ctx context.Context, log AppLog)
}
