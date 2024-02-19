package handlers

import (
	"github.com/shouryagautam/bankdeploy/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateAccountRequest represents the request structure for creating a new account.
type CreateAccountRequest struct {
	CustomerID  uint    `json:"customer_id" binding:"required"`
	Balance     float64 `json:"balance" binding:"required"`
	AccountType string  `json:"account_type" binding:"required"`
	NomineeID   uint    `json:"nominee_id"`
}

// CreateAccount creates a new account for a customer.
// @Summary Create a new account
// @Description Create a new account for a customer
// @Tags Accounts
// @Accept json
// @Produce json
// @Param body body CreateAccountRequest true "Account object to be created"
// @Success 201 {object} map[string]interface{} "Account created successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/account [post]
func CreateAccount(context *gin.Context) {
	var input CreateAccountRequest

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	customer, err := models.FindCustomerByID(input.CustomerID)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	account := models.Account{
		BranchID:    customer.BranchID,
		Balance:     input.Balance,
		AccountType: input.AccountType,
	}
	account.Customer = append(account.Customer, customer)

	if input.NomineeID != 0 {
		nominee, err := models.FindCustomerByID(input.NomineeID)
		if err != nil {
			context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		account.Customer = append(account.Customer, nominee)
	}

	savedAccount, err := account.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	mapping := models.CustomerToAccount{
		CustomerID: customer.ID,
		AccountID:  savedAccount.ID,
	}
	err = mapping.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}

	if input.NomineeID != 0 {
		mapping := models.CustomerToAccount{
			CustomerID: input.NomineeID,
			AccountID:  savedAccount.ID,
		}
		err = mapping.Save()
		if err != nil {
			context.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
			return
		}
	}

	context.JSON(http.StatusCreated, map[string]interface{}{"Account": savedAccount})
}

// GetAllAccountsByBranchID retrieves all accounts by branch ID.
// @Summary Get all accounts by branch ID
// @Description Retrieve all accounts by branch ID
// @Tags Accounts
// @Produce json
// @Param id path int true "Branch ID"
// @Success 200 {object} map[string]interface{} "Account retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/branch/{id}/account [get]
func GetAllAccountsByBranchID(context *gin.Context) {
	id := context.Param("id")
	ID, _ := strconv.ParseUint(id, 10, 0)
	accounts, err := models.FindAllAccountsByBranchID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Account": accounts})
}

// GetAccountById retrieves an account by its ID.
// @Summary Get an account by ID
// @Description Retrieve an account by its ID
// @Tags Accounts
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} map[string]interface{} "Account retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/account/{id} [get]
func GetAccountById(context *gin.Context) {
	id := context.Param("id")
	ID, _ := strconv.ParseUint(id, 10, 0)
	account, err := models.FindAccountByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Account": account})
}

// GetAllAccountsByCustomerID retrieves all accounts by customer ID.
// @Summary Get all accounts by customer ID
// @Description Retrieve all accounts by customer ID
// @Tags Accounts
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]interface{} "Account retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /customer/{id}/account [get]
func GetAllAccountsByCustomerID(context *gin.Context) {
	id := context.Param("id")
	ID, _ := strconv.ParseUint(id, 10, 0)

	accounts, err := models.FindAllAccountsByCustomerID(uint(ID))
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Account": accounts})
}

// GetAccountByAccountNumber retrieves an account by its account number.
// @Summary Get an account by account number
// @Description Retrieve an account by its account number
// @Tags Accounts
// @Produce json
// @Param number path string true "Account number"
// @Success 200 {object} map[string]interface{} "Account retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /customer/account/{number} [get]
func GetAccountByAccountNumber(context *gin.Context) {
	id := uuid.MustParse(context.Param("number"))

	account, err := models.FindAccountByAccountNumber(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Account": account})
}

// DeleteAllAccounts deletes all accounts.
// @Summary Delete all accounts
// @Description Delete all accounts
// @Tags Accounts
// @Produce json
// @Success 200 {object} map[string]interface{} "message: All tables have been deleted"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/account [delete]
func DeleteAllAccounts(context *gin.Context) {
	err := models.DeleteAllAccounts()

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"message": "All tables have been deleted"})
}

// DeleteAccountByID deletes an account by its ID.
// @Summary Delete an account by ID
// @Description Delete an account by its ID
// @Tags Accounts
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} map[string]interface{} "Branches deleted successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/account/{id} [delete]
func DeleteAccountByID(context *gin.Context) {
	id := context.Param("id")
	ID, _ := strconv.ParseUint(id, 10, 0)
	branches, err := models.DeleteAccountByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Branches": branches})
}

// UpdateAccount updates account information.
// @Summary Update account information
// @Description Update account information
// @Tags Accounts
// @Accept json
// @Produce json
// @Param body body models.Account true "Updated account object"
// @Success 201 {object} map[string]interface{} "Account updated successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/account [put]
func UpdateAccount(context *gin.Context) {
	var input models.Account

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	updatedAccount, err := input.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, map[string]interface{}{"Account": updatedAccount})
}
