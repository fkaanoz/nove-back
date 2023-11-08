package testgrp

import (
	"context"
	"net/http"
)

func HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("health check"))
	return nil
}
