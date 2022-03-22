package http

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
)

func (m *HTTPModule) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	glog.Infof("preflight request for %s", r.URL.Path)
	return
}

func (m *HTTPModule) allowCORS(h http.Handler) http.Handler {
	h = handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"*"}),
		handlers.AllowedHeaders([]string{"*"}),
	)(h)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				m.preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
