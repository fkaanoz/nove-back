package system

import (
	"context"
	"encoding/json"
	"net/http"
)

func HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	d := struct {
		Health bool
	}{
		Health: true,
	}

	js, err := json.Marshal(d)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return nil
}
