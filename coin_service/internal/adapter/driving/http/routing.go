package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) endpoints() {
	s.router.GET("/ping", s.Ping)

	txG := s.router.Group("/tx")
	txG.POST("/create", s.CreateTransaction)
	txG.GET("/get-all", s.GetAllTransactions)
	txG.GET("/get", s.GetTransactionByID)
	txG.GET("/total", s.GetTotalBalance)
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
