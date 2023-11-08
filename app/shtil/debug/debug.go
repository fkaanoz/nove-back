package debug

import (
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

func DebugApi() *httptreemux.ContextMux {
	mux := debugStandardLibrary()
	mux.Handle(http.MethodGet, "/debug/test-handler", DebugHandler)

	return mux
}
