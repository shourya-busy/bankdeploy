
package handlers

import (
	"github.com/shouryagautam/bankdeploy/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AddNomineeRequest represents the request structure for adding a nominee to an account.
type AddNomineeRequest struct {
	NomineeID     uint      `json:"nominee_id" binding:"required"`
	AccountNumber uuid.UUID `json:"account_number" binding:"required"`
}

// AddNominee adds a nominee to an account.
// @Summary Add a nominee to an account
// @Description Add a nominee to an account
// @Tags Nominees
// @Accept json
// @Produce json
// @Param body body AddNomineeRequest true "Add nominee request"
// @Success 200 {object} map[string]interface{} "Message: Nominee added"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /customer/account/nominee [post]
func AddNominee(context *gin.Context) {
	var input AddNomineeRequest

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	account, err := models.FindAccountByAccountNumber(input.AccountNumber)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	mapping := models.CustomerToAccount{
		CustomerID: input.NomineeID,
		AccountID:  account.ID,
	}

	err = mapping.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Message": "Nominee added"})
}

// DeleteNomineeFromAccountByID deletes a nominee from an account by ID.
// @Summary Delete a nominee from an account by ID
// @Description Delete a nominee from an account by ID
// @Tags Nominees
// @Produce json
// @Param number path string true "Account number"
// @Param id path int true "Nominee ID"
// @Success 200 {object} map[string]interface{} "Message: Nominee deleted"
// @Failure 400 {object} map[string]interface{} "error: Bad request"
// @Router /customer/account/{number}/nominee/{id} [delete]
func DeleteNomineeFromAccountByID(context *gin.Context) {
	number := uuid.MustParse(context.Param("number"))
	id := context.Param("id")
	ID, _ := strconv.ParseUint(id, 10, 0)

	err := models.DeleteNomineeFromAccountByID(number, uint(ID))
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"Message": "Nominee deleted"})
}
