package billingcycledelivery

import (
	"billing-cycle-ms-app/models"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	BUsercase models.BillingCycleUseCase
}

// New will initialize
func New(e *echo.Echo, bu models.BillingCycleUseCase) {
	handler := &Handler{
		BUsercase: bu,
	}
	// CreateBillingCycle - Creates a billing cycle
	e.POST("/customerBillManagement/v1/billingCycle", handler.CreateBillingCycle)
	// RetrieveBillingCycleById - Retrieves a billing cycle by ID
	e.GET("/customerBillManagement/v1/billingCycle/:billingCycleId", handler.RetrieveBillingCycleById)
}

func (c *Handler) CreateBillingCycle(ctx echo.Context) error {
	billingCycle := models.BillingCycle{}
	defer ctx.Request().Body.Close()
	err := json.NewDecoder(ctx.Request().Body).Decode(&billingCycle)
	if err != nil {
		//log.Debug("Failed to unmarshalling of json %s", err)
		return ctx.String(http.StatusInternalServerError, "")
	}
	response, validationerr := c.BUsercase.ValidateBillingCycle(ctx, billingCycle)
	if validationerr != nil {
		ctx.Logger().Debug("validation error while creating billing cycle")
		ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(ctx.Response()).Encode(response)

	}
	ctx.Logger().Debug("This is your billing related data -> %#v", response)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	ctx.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(ctx.Response()).Encode(response)
}

func (c *Handler) RetrieveBillingCycleById(ctx echo.Context) error {
	billingId := ctx.Param("billingCycleId")
	// below is used to fetch getters params passed as jsons
	//billingCycleId := ctx.QueryParam("billingCycleId")

	ctx.Logger().Info("Fetching billing cycle specification using billingCycleId - %s", billingId)
	billingCycleData, err := c.BUsercase.ValidateBillingIdToFetch(ctx, billingId)
	if err == nil {
		ctx.Logger().Debug("Billing data fetched in delivery layer - ", billingCycleData)
	}

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	ctx.Response().WriteHeader(http.StatusOK)
	if billingCycleData == nil {
		return json.NewEncoder(ctx.Response()).Encode(make(map[string]string))
	}
	return json.NewEncoder(ctx.Response()).Encode(billingCycleData)
}
