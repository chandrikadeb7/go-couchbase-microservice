package main

import (
	billingcyclebusiness "billing-cycle-ms-app/handlers/business"
	billingcycledelivery "billing-cycle-ms-app/handlers/delivery"
	billingcyclerepository "billing-cycle-ms-app/handlers/repository"
	billingcycleusecase "billing-cycle-ms-app/handlers/usecase"
	"billing-cycle-ms-app/logger"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	billingcyclerepo := billingcyclerepository.New()
	billingcyclebusiness := billingcyclebusiness.New(e, billingcyclerepo)
	billingcyclecase := billingcycleusecase.New(e, billingcyclebusiness)
	billingcycledelivery.New(e, billingcyclecase)

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	loggerConfiguration := logger.EchoLoggerConfig()
	e.Use(middleware.LoggerWithConfig(loggerConfiguration))
	e.Logger.SetLevel(log.DEBUG)

	/*

		End points to be developed and implemented in delivery layer

		// CreateBillingCycle - Creates a billing cycle
		 e.POST("/customerBillManagement/v1/billingCycle", c.CreateBillingCycle)

		// GetBillingCycles - Retrieves the billing cycles
		e.GET("/customerBillManagement/v1/billingCycle", c.GetBillingCycles)

		// RetrieveBillingCycleById - Retrieves a billing cycle by ID
		e.GET("/customerBillManagement/v1/billingCycle/:billingCycleId", c.RetrieveBillingCycleById)

		// UpdateBillingCycle - Updates the status of a billing cycle
		e.PATCH("/customerBillManagement/v1/billingCycle/:billingCycleId", c.UpdateBillingCycle)

	*/

	// Start server
	//e.Logger.Fatal(e.Start(":8081"))

	// enabling TLS
	// os.Getenv("PORT")
	if err := e.StartTLS(":"+os.Getenv("PORT"), "server.crt", "server.key"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
