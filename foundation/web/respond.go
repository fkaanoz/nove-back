package web

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, v interface{}, statusCode int) error {

	js, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	if _, err := w.Write(js); err != nil {
		return err
	}

	return nil
}
