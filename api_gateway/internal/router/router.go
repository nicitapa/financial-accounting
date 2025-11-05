package router

import (
	"api_gateway/internal/config"
	"api_gateway/internal/proxy"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
)

func mount(r chi.Router, prefix string, target string) {
	u, err := url.Parse(target)
	if err != nil {
		panic("invalid proxy target: " + err.Error())
	}
	h := proxy.New(u)  // ваш http.Handler на базе httputil.NewSingleHostReverseProxy
	r.Mount(prefix, h) // chi сам срежет prefix для h
}

func Setup(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	mount(r, "/tx", cfg.TransactionServiceURL)
	mount(r, "/auth", cfg.AuthServiceURL)

	return r
}
