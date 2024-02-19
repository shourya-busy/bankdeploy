package routes

import (
	"github.com/shouryagautam/bankdeploy/handlers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)

// @BasePath /super
// @BasePath /admin
// @BasePath /manager
// @BasePath /customer
func Router() {

	router := gin.Default()

	superRoutes := router.Group("/super")
	superRoutes.POST("/bank", handlers.CreateBank)
	superRoutes.GET("/bank", handlers.GetAllBanks)
	superRoutes.GET("/bank/:id", handlers.GetBankByID)
	superRoutes.PUT("/bank", handlers.UpdateBank)
	superRoutes.DELETE("/bank/:id", handlers.DeleteBankByID)
	superRoutes.DELETE("/bank", handlers.DeleteAllBanks)

	adminRoutes := router.Group("/admin")
	adminRoutes.POST("/branch", handlers.CreateBranch)
	adminRoutes.GET("/bank/:id/branch", handlers.GetAllBranchesByBankID)
	adminRoutes.GET("branch/:id", handlers.GetBranchByID)
	adminRoutes.PUT("/branch", handlers.UpdateBranch)
	adminRoutes.DELETE("/branch/:id", handlers.DeleteBranchByID)
	adminRoutes.DELETE("/branch", handlers.DeleteAllBranches)

	managerRoutes := router.Group("/manager")
	managerRoutes.POST("/customer", handlers.CreateCustomer)
	managerRoutes.POST("/account", handlers.CreateAccount)
	managerRoutes.GET("branch/:id/account", handlers.GetAllAccountsByBranchID)
	managerRoutes.GET("/account/:id", handlers.GetAccountById)
	managerRoutes.GET("/branch/:id/customer", handlers.GetAllCustomersByBranchID)
	managerRoutes.GET("/customer/:id", handlers.GetCustomerByID)
	managerRoutes.PUT("/account", handlers.UpdateAccount)
	managerRoutes.PUT("/customer", handlers.UpdateCustomer)
	managerRoutes.DELETE("/account", handlers.DeleteAllAccounts)
	managerRoutes.DELETE("/account/:id", handlers.DeleteAccountByID)
	managerRoutes.DELETE("/customer", handlers.DeleteAllCustomers)
	managerRoutes.DELETE("/customer/:id", handlers.DeleteCustomerByID)

	userRoutes := router.Group("/customer")
	userRoutes.POST("/account/deposit", handlers.Deposit)
	userRoutes.POST("/account/withdraw", handlers.Withdraw)
	userRoutes.POST("/account/transfer", handlers.Transfer)
	userRoutes.GET("/account/:number/nominee", handlers.GetAllNomineesByAccountNumber)
	userRoutes.GET("/:id/account", handlers.GetAllAccountsByCustomerID)
	userRoutes.GET("/account/:number", handlers.GetAccountByAccountNumber)
	userRoutes.GET("/account/:number/transactions", handlers.GetAllTransactionsByAccountNumber)
	userRoutes.GET("/account/transactions/:id", handlers.GetTransactionByID)
	userRoutes.PUT("/account/nominee", handlers.AddNominee)
	userRoutes.DELETE("/account/:number/nominee/:id", handlers.DeleteNomineeFromAccountByID)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
