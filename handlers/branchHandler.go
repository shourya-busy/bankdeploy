

package handlers

import (
	"github.com/shouryagautam/bankdeploy/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBranch creates a new branch.
// @Summary Create a new branch
// @Description Create a new branch
// @Tags Branches
// @Accept json
// @Produce json
// @Param body body models.Branch true "Branch object to be created"
// @Success 201 {object} models.Branch "Branch created successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /admin/branch [post]
func CreateBranch(context *gin.Context) {
	var input models.Branch

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	savedBranch,err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Branch":savedBranch})
}

// GetAllBranchesByBankID retrieves all branches by bank ID.
// @Summary Get all branches by bank ID
// @Description Retrieve all branches by bank ID
// @Tags Branches
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {array} models.Branch "Branches retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /admin/branch/{id} [get]
func GetAllBranchesByBankID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)

	branches,err := models.FindAllBranchesByBankID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Branches":branches})
}

// GetBranchByID retrieves a branch by its ID.
// @Summary Get a branch by ID
// @Description Retrieve a branch by its ID
// @Tags Branches
// @Produce json
// @Param id path int true "Branch ID"
// @Success 200 {object} models.Branch "Branch retrieved successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /admin/branch/{id} [get]
func GetBranchByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	branch,err := models.FindBranchByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Branch":branch})
}

// DeleteAllBranches deletes all branches.
// @Summary Delete all branches
// @Description Delete all branches
// @Tags Branches
// @Produce json
// @Success 200 {string} message "All branches have been deleted"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /admin/branch [delete]
func DeleteAllBranches(context *gin.Context) {
	err := models.DeleteAllBranches()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"All branches have been deleted"})
}

// DeleteBranchByID deletes a branch by its ID.
// @Summary Delete a branch by ID
// @Description Delete a branch by its ID
// @Tags Branches
// @Produce json
// @Param id path int true "Branch ID"
// @Success 200 {object} models.Branch "Branch deleted successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /admin/branch/{id} [delete]
func DeleteBranchByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	branches,err := models.DeleteBranchByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Branch":branches})
}

// UpdateBranch updates a branch.
// @Summary Update a branch
// @Description Update a branch with new information
// @Tags Branches
// @Accept json
// @Produce json
// @Param body body models.Branch true "Branch object to be updated"
// @Success 201 {object} models.Branch "Branch updated successfully"
// @Failure 400 {object} map[string]interface{} "error": "Bad request"
// @Router /admin/branch [put]
func UpdateBranch(context *gin.Context) {
	var input models.Branch

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}

	updatedBranch,err := input.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Branch":updatedBranch})
}

