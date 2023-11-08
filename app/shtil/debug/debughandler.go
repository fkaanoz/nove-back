package debug

import "net/http"

func DebugHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test debug handler"))
}
