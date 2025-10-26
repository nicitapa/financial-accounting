package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/nicitapa/financial-accounting/api_gateway/internal/config"
	"github.com/nicitapa/financial-accounting/api_gateway/internal/proxy"
	"net/http"
	"net/url"
)

func mount(r chi.Router, prefix string, target string) {
	u, _ := url.Parse(target)
	h := proxy.New(u)  // ваш http.Handler на базе httputil.NewSingleHostReverseProxy
	r.Mount(prefix, h) // chi сам срежет prefix для h
}

func Setup(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	mount(r, "/auth", cfg.AuthServiceURL)
	mount(r, "/api/employees", cfg.EmployeesServiceURL)

	return r
}
