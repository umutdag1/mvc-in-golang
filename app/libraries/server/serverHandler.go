package server

import (
	"net/http"

	"github.com/rest-api/config"
)

func setHandlers() {
	for _, route := range config.ROUTES {
		handlerFunc := authHandler(route.Handler, route.Method)
		mux.HandleFunc(route.Path, handlerFunc)
	}
}

func authHandler(f http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.NotFound(w, r)
			return
		}
		setHeaders(w)
		f(w, r)
	}
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
