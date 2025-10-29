package http

import (
	"net/http"
	"time"

	"github.com/prankevich/Auth_service/internal/config"
	"github.com/prankevich/Auth_service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	uc     *usecase.UseCases
}

const httpServerReadHeaderTimeout = 70 * time.Second

func New(
	cfg *config.Config,
	uc *usecase.UseCases,
) *http.Server {
	r := gin.New()

	srv := &Server{
		router: r,
		cfg:    cfg,
		uc:     uc,
	}

	srv.endpoints()

	httpServer := &http.Server{
		Addr:              cfg.HTTPPort,
		Handler:           srv,
		ReadHeaderTimeout: httpServerReadHeaderTimeout,
	}

	// srv.log.Info(fmt.Sprintf("HTTP server is initialized on port: %v", cfg.HTTPPort))

	return httpServer
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
