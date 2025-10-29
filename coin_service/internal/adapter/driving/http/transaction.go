package http

import (
	"coin_service/internal/domain"
	"coin_service/internal/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateTransactionRequest struct {
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
}

func (s *Server) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errs.ErrUserNotFound) || errors.Is(err, errs.ErrNotfound):
		c.JSON(http.StatusNotFound, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidTransactionID) || errors.Is(err, errs.ErrInvalidRequestBody):
		c.JSON(http.StatusBadRequest, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrIncorrectUsernameOrPassword) || errors.Is(err, errs.ErrInvalidToken):
		c.JSON(http.StatusUnauthorized, CommonError{Error: err.Error()})
	case errors.Is(err, errs.ErrInvalidFieldValue) ||
		errors.Is(err, errs.ErrUsernameAlreadyExists):
		c.JSON(http.StatusUnprocessableEntity, CommonError{Error: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, CommonError{Error: err.Error()})
	}
}

func (s *Server) CreateTransaction(c *gin.Context) {
	var transaction CreateTransactionRequest
	if err := c.ShouldBindJSON(&transaction); err != nil {
		s.handleError(c, errors.Join(errs.ErrInvalidFieldValue, err))
		return
	}
	if transaction.Category == "" || transaction.Currency == "" || transaction.Amount == 0 {
		s.handleError(c, errs.ErrInvalidFieldValue)
		return
	}
	if err := s.uc.TransactionCreator.CreateTransaction(c, domain.Transaction{
		Category:    domain.Category(transaction.Category),
		Amount:      transaction.Amount,
		Currency:    transaction.Currency,
		Description: transaction.Description,
	}); err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, CommonResponse{Message: "Transaction created successfully!"})
}

func (s *Server) GetAllTransactions(c *gin.Context) {
	transactions, err := s.uc.TransactionList.GetAllTransactions(c)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (s *Server) GetTransactionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		s.handleError(c, errs.ErrInvalidTransactionID)
		return
	}

	transaction, err := s.uc.TransactionGetter.GetTransactionByID(c, id)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": transaction})
}

func (s *Server) GetTotalBalance(c *gin.Context) {
	total, err := s.uc.TotalBalance.GetTotalBalance(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, total)
		return
	}
	c.JSON(http.StatusOK, total)
}
