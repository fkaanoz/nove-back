package testgrp

import (
	"context"
	"net/http"
	"time"
)

func HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	time.Sleep(time.Millisecond * 1500)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("health check"))
	return nil
}
