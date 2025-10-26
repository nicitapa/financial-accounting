package internal

import (
	"coin_service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	router *gin.Engine
	svc    *service.TransactionService
}

func NewController(r *gin.Engine, svc *service.TransactionService) *Controller {
	return &Controller{
		router: r,
		svc:    svc,
	}
}

func (ctrl *Controller) RegisterEndpoints() {
	ctrl.router.GET("/ping", ctrl.Ping)
	ctrl.router.POST("/transactions", ctrl.CreateTransaction)
	ctrl.router.GET("/transactions", ctrl.GetAllTransaction)
	//ctrl.router.GET("/transactions/:id", ctrl.GetTransactionByID)
	//ctrl.router.PUT("/transactions/:id", ctrl.UpdateTransactionByID)
	//ctrl.router.DELETE("/transactions/:id", ctrl.DeleteTransactionByID)
}

func (ctrl *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, 200)
}

func (ctrl *Controller) RunServer(address string) error {
	ctrl.RegisterEndpoints()
	return ctrl.router.Run(address)
}
