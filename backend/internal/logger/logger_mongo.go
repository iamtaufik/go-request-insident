package logger

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Level string

const (
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
)

type AppLog struct {
	TS       time.Time              `bson:"ts"`
	Level    Level                  `bson:"level"`
	Service  string                 `bson:"service"`
	Event    string                 `bson:"event"`              // e.g. "ORDER_CREATE_FAILED"
	Message  string                 `bson:"message"`            // human readable
	UserID   *string                `bson:"user_id,omitempty"`  // optional
	Entity   *string                `bson:"entity,omitempty"`   // e.g. "order"
	EntityID *string                `bson:"entity_id,omitempty"`
	Meta     map[string]interface{} `bson:"meta,omitempty"`
	Err      *LogError              `bson:"error,omitempty"`
}

type LogError struct {
	Type    string `bson:"type,omitempty"`
	Message string `bson:"message,omitempty"`
}

type AppLogger struct {
	coll    *mongo.Collection
	service string
}

func NewAppLogger(coll *mongo.Collection, service string) *AppLogger {
	return &AppLogger{coll: coll, service: service}
}

func (l *AppLogger) Log(ctx context.Context, log AppLog) {
	log.TS = time.Now()
	log.Service = l.service

	cctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	_, _ = l.coll.InsertOne(cctx, log)
}

func (l *AppLogger) Error(ctx context.Context, event, msg string, err error, meta map[string]interface{}) {
	le := &LogError{}
	if err != nil {
		le.Type = "error"
		le.Message = err.Error()
	}
	l.Log(ctx, AppLog{
		Level:   ERROR,
		Event:   event,
		Message: msg,
		Meta:    meta,
		Err:     le,
	})
}

func strPtr(s string) *string { return &s }
