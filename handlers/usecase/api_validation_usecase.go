package billingcycleusecase

import (
	"billing-cycle-ms-app/constant"
	helper "billing-cycle-ms-app/handlers"
	"billing-cycle-ms-app/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type billingUseCase struct {
	//billingRepo models.BillingCycleRepository
	billingBusiness models.BillingCycleBusiness
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	// fmt.Println("I am validating mandatory parameters")
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New(e *echo.Echo, billingBusiness models.BillingCycleBusiness) models.BillingCycleUseCase {
	e.Validator = &CustomValidator{validator: validator.New()}
	return &billingUseCase{
		billingBusiness: billingBusiness,
	}
}

func (bu *billingUseCase) ValidateBillingCycle(ctx echo.Context, billingData models.BillingCycle) (interface{}, error) {
	// lets do mandatory check validations
	err := ctx.Validate(billingData)
	if err != nil {
		return constant.New("BSS-BILL-CBCI-40000001", err), err
	}
	validation_response, err_in_date := ValidateBillingStartEndDate(billingData.StartDate, billingData.EndDate)
	if err_in_date != nil {
		switch validation_response {
		case "BSS-BILL-CBCI-40000004":
			return constant.New("BSS-BILL-CBCI-40000004", err_in_date), err_in_date
		case "BSS-BILL-CBCI-40000005":
			return constant.New("BSS-BILL-CBCI-40000005", err_in_date), err_in_date
		case "BSS-BILL-CBCI-40000003":
			return constant.New("BSS-BILL-CBCI-40000003", err_in_date), err_in_date
		case "BSS-BILL-CBCI-40000007":
			return constant.New("BSS-BILL-CBCI-40000007", err_in_date), err_in_date
		}
	}
	ctx.Logger().Debug("Start and End date validated to %s", validation_response)
	valid_frequency_response, err_freq := ValidateFrequency(billingData.CycleSpecification.Frequency)
	if err_freq != nil {
		return constant.New("BSS-BILL-CBCI-40000009", err_freq), err_freq
	}
	/*
		valid_cycle_instance, err_instance := ValidateCycleInstance(billingData.CycleInstance)
		if err_instance == nil {
			return constant.New("BSS-BILL-CBCI-400000010", err_instance), err_instance
		}
	*/
	ctx.Logger().Debug("Frequency validated to %t", valid_frequency_response)
	// ctx.Logger().Debug("CycleInsatance validated to %t", valid_cycle_instance)
	return bu.billingBusiness.SetBillingData(ctx, billingData)

}

func (bu *billingUseCase) ValidateBillingIdToFetch(ctx echo.Context, billingId string) (interface{}, error) {
	validationCheck, errValidation := SplitIdToVerifyDate(billingId)
	if validationCheck {
		return bu.billingBusiness.FetchBillingDataForID(ctx, billingId)
	}
	return constant.New("BSS-BILL-CBCI-40000002", errValidation), errValidation
}

func SplitIdToVerifyDate(billingId string) (bool, error) {
	sliceOfId := strings.Split(billingId, constant.Underscore)
	if len(sliceOfId) < 3 {
		return false, fmt.Errorf("Invalid ID passed should be of format - id_shiftfate_date .")
	}
	inputEndDate := sliceOfId[2]
	endDate, err := ValidateDate(inputEndDate)
	if err != nil {
		return false, fmt.Errorf("Invalid date passed as input should of format yyyy-mm-dd . Passed - %s", endDate)
	}
	return true, nil
}

func ValidateDate(ds string) (time.Time, error) {
	t, err := time.Parse(constant.LayoutISO, ds)
	if err != nil {
		return t, err
	}
	return t, nil
}

// ValidateBillingStartEndDate checks is startdate <= enddate
func ValidateBillingStartEndDate(startdate, enddate string) (string, error) {

	tstart, err := ValidateDate(startdate)
	if err != nil {
		return "BSS-BILL-CBCI-40000004", fmt.Errorf("cannot parse startdate: %v", err)
	}
	tend, err := ValidateDate(enddate)
	if err != nil {
		return "BSS-BILL-CBCI-40000005", fmt.Errorf("cannot parse enddate: %v", err)
	}

	if tstart.After(tend) {
		return "BSS-BILL-CBCI-40000003", fmt.Errorf("startdate > enddate - please set proper data boundaries")
	}

	// validates if the days difference between start and endate is greater than 30.
	stDay := helper.Date(tstart.Year(), int(tstart.Month()), tstart.Day())
	ltDay := helper.Date(tend.Year(), int(tend.Month()), tend.Day())
	days := ltDay.Sub(stDay).Hours() / 24
	if days > 30 {
		return "BSS-BILL-CBCI-40000007", fmt.Errorf("startdate < enddate - but cannot be greater than 30 days.")
	}
	return "true", err
}

func ValidateFrequency(frequency string) (bool, error) {
	frequency_conv := strings.ToUpper(frequency)
	if frequency_conv == "MONTHLY" || frequency_conv[0] == 'M' {
		return true, nil
	}
	return false, fmt.Errorf("values other than monthly or m are not allowed passed %s", frequency)
}

/*
func ValidateCycleInstance(cycleInstance string) (bool, error) {
	cycleInstanceNum, _ := strconv.Atoi(cycleInstance)
	if cycleInstanceNum >= 1 && cycleInstanceNum <= 12 {
		return true, nil
	}
	return false, fmt.Errorf("cycle instance cannot be less than 1 and greater than 12 passed %d", cycleInstance)
}
*/
