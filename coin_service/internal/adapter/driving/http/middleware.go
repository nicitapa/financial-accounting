package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prankevich/Auth_service/pkg"

	"net/http"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
)

func (s *Server) checkUserAuthentication(c *gin.Context) {
	token, err := s.extractTokenFromHeader(c, authorizationHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	userID, isRefresh, _, err := pkg.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
		return
	}

	if isRefresh {
		c.AbortWithStatusJSON(http.StatusUnauthorized, CommonError{Error: "inappropriate token"})
		return
	}

	c.Set(userIDCtx, userID)
}
