
package handlers

import (
	"github.com/shouryagautam/bankdeploy/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Deposit handles depositing money into an account.
// @Summary Deposit money into an account
// @Description Deposit money into an account
// @Tags Transactions
// @Accept json
// @Produce json
// @Param body body models.Transaction true "Transaction object to be deposited"
// @Success 202 {object} map[string]interface{} "message: Your Transaction has been completed successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /customer/account/deposit [post]
func Deposit(context *gin.Context) {
	var input models.Transaction
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	transaction := models.Transaction{
		AccountID:          input.AccountID,
		Amount:             input.Amount,
		ModeOfPayment:      input.ModeOfPayment,
		TypeOfTransaction: "Deposit",
		Time:               time.Now(),
	}

	err := models.AccountDeposit(transaction.AccountID, transaction.Amount)
	if err != nil {
		context.JSON(http.StatusNotModified, map[string]interface{}{"error": err.Error()})
		return
	}

	savedTransaction, err := transaction.Save()
	if err != nil {
		context.JSON(http.StatusResetContent, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, map[string]interface{}{"message": "Your Transaction has been completed successfully", "data": savedTransaction})
}

// Withdraw handles withdrawing money from an account.
// @Summary Withdraw money from an account
// @Description Withdraw money from an account
// @Tags Transactions
// @Accept json
// @Produce json
// @Param body body models.Transaction true "Transaction object to be withdrawn"
// @Success 202 {object} map[string]interface{} "message: Your Transaction has been completed successfully"
// @Failure 502 {object} map[string]interface{} "error: Bad gateway"
// @Router /customer/account/withdraw [post]
func Withdraw(context *gin.Context) {
	var input models.Transaction
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	transaction := models.Transaction{
		AccountID:          input.AccountID,
		Amount:             input.Amount,
		ModeOfPayment:      input.ModeOfPayment,
		TypeOfTransaction: "Withdraw",
		Time:               time.Now(),
	}

	err := models.AccountWithdrawal(transaction.AccountID, transaction.Amount)
	if err != nil {
		context.JSON(http.StatusBadGateway, map[string]interface{}{"error": err.Error()})
		return
	}

	savedTransaction, err := transaction.Save()
	if err != nil {
		context.JSON(http.StatusResetContent, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, map[string]interface{}{"message": "Your Transaction has been completed successfully", "data": savedTransaction})
}

// Transfer handles transferring money between accounts.
// @Summary Transfer money between accounts
// @Description Transfer money between accounts
// @Tags Transactions
// @Accept json
// @Produce json
// @Param body body models.Transaction true "Transaction object to be transferred"
// @Success 202 {object} map[string]interface{} "message: Your Transaction has been completed successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /customer/account/transfer [post]
func Transfer(context *gin.Context) {
	var input models.Transaction
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	transaction := models.Transaction{
		AccountID:             input.AccountID,
		Amount:                input.Amount,
		ModeOfPayment:         input.ModeOfPayment,
		TypeOfTransaction:     "Transfer",
		ReceiverAccountNumber: input.ReceiverAccountNumber,
		Time:                  time.Now(),
	}

	err := models.AccountTransfer(transaction.AccountID, transaction.ReceiverAccountNumber, transaction.Amount)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	savedTransaction, err := transaction.Save()
	if err != nil {
		context.JSON(http.StatusResetContent, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, map[string]interface{}{"message": "Your Transaction has been completed successfully", "data": savedTransaction})
}

// GetAllTransactionsByAccountNumber retrieves all transactions by account number.
// @Summary Get all transactions by account number
// @Description Retrieve all transactions by account number
// @Tags Transactions
// @Produce json
// @Param number path string true "Account number"
// @Success 200 {object} map[string]interface{} "Transactions retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /customer/account/{number}/transactions [get]
func GetAllTransactionsByAccountNumber(context *gin.Context) {
	number := uuid.MustParse(context.Param("number"))

	transactions, err := models.FindAllTransactionsByAccountNumber(number)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Transactions": transactions})
}

// GetTransactionByID retrieves a transaction by its ID.
// @Summary Get a transaction by ID
// @Description Retrieve a transaction by its ID
// @Tags Transactions
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} map[string]interface{} "Transaction retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /customer/account/transactions/{id} [get]
func GetTransactionByID(context *gin.Context) {
	id := context.Param("id")
	ID, _ := strconv.ParseUint(id, 10, 0)
	transaction, err := models.FindTransactionByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Transaction": transaction})
}
