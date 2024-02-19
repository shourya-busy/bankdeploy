

package handlers

import (
	"github.com/shouryagautam/bankdeploy/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



// @BasePath /super
// @Summary Create a new bank
// @Description Create a new bank with the provided name
// @Tags Banks
// @Accept json
// @Produce json
// @Param name body string true "Name of the bank"
// @Success 201 {object} models.Bank "Bank created successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /super/bank [post]
func CreateBank(context *gin.Context) {
	var input models.Bank

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}

	bank := models.Bank{
		Name: input.Name,
	}

	savedBank,err := bank.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Bank":savedBank})
	
}

// @Summary Get all banks
// @Description Retrieve all banks
// @Tags Banks
// @Produce json
// @Success 200 {object} map[string]interface{} "Banks retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /super/bank [get]
func GetAllBanks(context *gin.Context) {

	banks,err := models.FindAllBanks()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Banks":banks})
}

// @Summary Get a bank by ID
// @Description Retrieve a bank by its ID
// @Tags Banks
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {object} models.Bank "Bank retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /super/bank/{id} [get]
func GetBankByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,64)
	bank,err := models.FindBankByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Bank":bank})
}

// @Summary Delete all banks
// @Description Delete all banks
// @Tags Banks
// @Produce json
// @Success 200 {object} map[string]interface{} "message": "All rows have been deleted"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /super/bank [delete]
func DeleteAllBanks(context *gin.Context) {

	err := models.DeleteAllBanks()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"All rows have been deleted"})
}

// @Summary Delete a bank by ID
// @Description Delete a bank by its ID
// @Tags Banks
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {object} models.Bank "Bank deleted successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /super/bank/{id} [delete]
func DeleteBankByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	bank,err := models.DeleteBankByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Bank":bank})
}

// @Summary Update a bank
// @Description Update a bank with new information
// @Tags Banks
// @Accept json
// @Produce json
// @Param name body string true "Name of the bank"
// @Success 201 {object} models.Bank "Bank updated successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /super/bank [put]
func UpdateBank(context *gin.Context) {
	var input models.Bank

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}

	updatedBank,err := input.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Bank":updatedBank})
	
}

