package internal

import (
	"coin_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (ctrl *Controller) CreateTransaction(c *gin.Context) {
	var tx models.Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: " + err.Error()})
		return
	}
	if tx.Timestamp.IsZero() {
		tx.Timestamp = time.Now()
	}
	if err := ctrl.svc.CreateTransactions(tx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save transaction: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "saved",
	})
}
func (ctrl *Controller) GetAllTransaction(c *gin.Context) {
	txs, err := ctrl.svc.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": txs,
		"count":        len(txs),
	})
}
