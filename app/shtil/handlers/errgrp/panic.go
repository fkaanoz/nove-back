package errgrp

import (
	"context"
	"net/http"
)

func PanicSimulation(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	panic("panic simulation")
}
