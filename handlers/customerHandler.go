

package handlers

import (
	"github.com/shouryagautam/bankdeploy/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCustomerRequest represents the request structure for creating a new customer.
type CreateCustomerRequest struct {
	BranchID     uint    `json:"branch_id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	PAN          string  `json:"pan" binding:"required"`
	DOB          string  `json:"dob" binding:"required"`
	Age          uint    `json:"age" binding:"required"`
	Phone        uint    `json:"phone" binding:"required"`
	Address      string  `json:"address" binding:"required"`
	Balance      float64 `json:"balance" binding:"required"`
	AccountType  string  `json:"account_type" binding:"required"`
}

// CreateCustomer creates a new customer and associated account.
// @Summary Create a new customer and account
// @Description Create a new customer and associated account
// @Tags Customers
// @Accept json
// @Produce json
// @Param body body CreateCustomerRequest true "Customer object to be created"
// @Success 201 {object} map[string]interface{} "Customer created successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/customer [post]
func CreateCustomer(context *gin.Context) {
	var input CreateCustomerRequest

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	customer := models.Customer{
		BranchID: input.BranchID,
		Name:     input.Name,
		PAN:      input.PAN,
		DOB:      input.DOB,
		Age:      input.Age,
		Phone:    input.Phone,
		Address:  input.Address,
	}

	account := &models.Account{
		BranchID:    input.BranchID,
		Balance:     input.Balance,
		AccountType: input.AccountType,
	}

	savedCustomer, err := customer.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}

	savedAccount, err := account.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}

	savedCustomer.Account = append(savedCustomer.Account, savedAccount)

	mapping := models.CustomerToAccount{
		CustomerID: savedCustomer.ID,
		AccountID:  savedAccount.ID,
	}

	err = mapping.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"err": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, map[string]interface{}{"Customer": savedCustomer})
}

// GetAllCustomersByBranchID retrieves all customers by branch ID.
// @Summary Get all customers by branch ID
// @Description Retrieve all customers by branch ID
// @Tags Customers
// @Produce json
// @Param id path int true "Branch ID"
// @Success 200 {object} map[string]interface{} "Customer retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/customer/{id} [get]
func GetAllCustomersByBranchID(context *gin.Context) {
	id := context.Param("id")
	ID, _ := strconv.ParseUint(id, 10, 0)

	customers, err := models.FindAllCustomersByBranchID(uint(ID))
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Customer": customers})
}

// GetCustomerByID retrieves a customer by their ID.
// @Summary Get a customer by ID
// @Description Retrieve a customer by their ID
// @Tags Customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]interface{} "Customer retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/customer/{id} [get]
func GetCustomerByID(context *gin.Context) {
	id := context.Param("id")
	ID, _ := strconv.ParseUint(id, 10, 0)
	customer, err := models.FindCustomerByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Customer": customer})
}

// DeleteAllCustomers deletes all customers.
// @Summary Delete all customers
// @Description Delete all customers
// @Tags Customers
// @Produce json
// @Success 200 {object} map[string]interface{} "message: All Customers have been deleted"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/customer [delete]
func DeleteAllCustomers(context *gin.Context) {
	err := models.DeleteAllCustomers()

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"message": "All Customers have been deleted"})
}

// DeleteCustomerByID deletes a customer by their ID.
// @Summary Delete a customer by ID
// @Description Delete a customer by their ID
// @Tags Customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]interface{} "Customer deleted successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/customer/{id} [delete]
func DeleteCustomerByID(context *gin.Context) {
	id := context.Param("id")
	ID, _ := strconv.ParseUint(id, 10, 0)
	customer, err := models.DeleteCustomersByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Customer": customer})
}

// UpdateCustomer updates customer information.
// @Summary Update customer information
// @Description Update customer information
// @Tags Customers
// @Accept json
// @Produce json
// @Param body body models.Customer true "Updated customer object"
// @Success 201 {object} map[string]interface{} "Customer updated successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /manager/customer [put]
func UpdateCustomer(context *gin.Context) {
	var input models.Customer

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	updatedCustomer, err := input.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, map[string]interface{}{"Customer": updatedCustomer})
}

// GetAllNomineesByAccountNumber retrieves all nominees by account number.
// @Summary Get all nominees by account number
// @Description Retrieve all nominees by account number
// @Tags Customers
// @Produce json
// @Param number path string true "Account number"
// @Success 200 {object} map[string]interface{} "Nominees retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /customer/account/{number}/nominee [get]
func GetAllNomineesByAccountNumber(context *gin.Context) {
	number := uuid.MustParse(context.Param("number"))

	customers, err := models.FindAllCustomersByAccountNumber(number)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Nominees": customers})
}
