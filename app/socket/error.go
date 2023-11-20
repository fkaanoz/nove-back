package socket

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse is helper to respond with errors when websocket connection failed.
func ErrorResponse(w http.ResponseWriter, message string) {
	e := struct {
		Message string
	}{
		Message: message,
	}

	js, _ := json.Marshal(e)

	w.WriteHeader(http.StatusNotFound)
	w.Write(js)
}
