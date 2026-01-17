package logger

import "context"

type NoopLogger struct{}

func (NoopLogger) Log(ctx context.Context, log AppLog) {}

func (NoopLogger) Error(ctx context.Context, event, msg string, err error, meta map[string]interface{}) {}
