package web

import (
	"context"
)

type keyType string

var CtxKey keyType = "custom-context"

type Values struct {
	TraceID    string
	StatusCode int
}

func GetTraceID(ctx context.Context) string {
	val, ok := ctx.Value(CtxKey).(Values)
	if !ok {
		return "000-000-000"
	}
	return val.TraceID
}

func GetStatusCode(ctx context.Context) int {
	val, ok := ctx.Value(CtxKey).(Values)
	if !ok {
		return 400
	}
	return val.StatusCode
}

func SetStatusCode(ctx context.Context, statusCode int) context.Context {
	val, ok := ctx.Value(CtxKey).(Values)
	if !ok {
		return nil
	}
	return context.WithValue(context.Background(), CtxKey, Values{
		TraceID:    val.TraceID,
		StatusCode: statusCode,
	})
}
