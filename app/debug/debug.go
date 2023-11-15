package debug

import (
	"encoding/json"
	"github.com/dimfeld/httptreemux/v5"
	"net/http"
	"net/http/pprof"
)

func debugStandardLibrary() *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()

	mux.Handle(http.MethodGet, "/debug/", pprof.Index)
	mux.Handle(http.MethodGet, "/debug/profile", pprof.Profile)
	mux.Handle(http.MethodGet, "/debug/cmdline", pprof.Cmdline)
	mux.Handle(http.MethodGet, "/debug/symbol", pprof.Symbol)
	mux.Handle(http.MethodGet, "/debug/trace", pprof.Trace)

	return mux
}

func Mux() *httptreemux.ContextMux {
	mux := debugStandardLibrary()
	mux.Handle(http.MethodGet, "/debug/test-handler", TestHandler)

	return mux
}

// TestHandler is a mock handler for debug api. DELETE LATER.
func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	d := struct {
		Status string
	}{
		Status: "test handler ok",
	}

	js, err := json.Marshal(d)
	if err != nil {
		w.Write([]byte("test handler error"))
		return
	}

	w.Write(js)
}
