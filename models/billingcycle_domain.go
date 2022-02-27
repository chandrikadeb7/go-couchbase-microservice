package models

import "github.com/labstack/echo/v4"

type BillingCycleRepository interface {
	InsertData(BillingCycle) (interface{}, error)
	CheckForBillingId(id string) (interface{}, error)
	GetBillingCycleById(id string) (interface{}, error)
}

type BillingCycleUseCase interface {
	ValidateBillingCycle(echo.Context, BillingCycle) (interface{}, error)
	ValidateBillingIdToFetch(echo.Context, string) (interface{}, error)
}

type BillingCycleBusiness interface {
	SetBillingData(echo.Context, BillingCycle) (interface{}, error)
	FetchBillingDataForID(echo.Context, string) (interface{}, error)
}
