package billingcyclebusiness

import (
	"billing-cycle-ms-app/constant"
	helper "billing-cycle-ms-app/handlers"
	"billing-cycle-ms-app/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type billingcycleBusiness struct {
	billingRepo models.BillingCycleRepository
}

func New(e *echo.Echo, billingRepo models.BillingCycleRepository) models.BillingCycleBusiness {
	return &billingcycleBusiness{
		billingRepo: billingRepo,
	}
}

func (bu *billingcycleBusiness) SetBillingData(ctx echo.Context, billingData models.BillingCycle) (interface{}, error) {
	billingData.Id = GenrateId(billingData.CycleSpecification.Id, billingData.StartDate, billingData.EndDate)
	billingData.BillingDate = PopulateBillDate(billingData.EndDate)
	billingData.CycleYear = PopulateCycleYear(billingData.StartDate)
	db_response, db_err := bu.billingRepo.InsertData(billingData)
	if db_err != nil {
		return db_response, db_err
	}
	return billingData, nil
	// Getting null as immediately calling the getter method.
	// return bu.billingRepo.CheckForBillingId(billingData.Id)
}

func (bu *billingcycleBusiness) FetchBillingDataForID(ctx echo.Context, billingId string) (interface{}, error) {
	// bu.billingRepo.CheckForBillingId(billingId)
	return bu.billingRepo.GetBillingCycleById(billingId)
}

func GenrateId(id string, startDate string, endDate string) string {
	startDayStr, _ := helper.HelperDate(startDate)
	startDayNum := startDayStr.Day()
	idLogic := [3]string{id, strconv.Itoa(startDayNum), endDate}
	fmt.Println(strings.Join(idLogic[:], "_"))
	return strings.Join(idLogic[:], "_")
}

func PopulateBillDate(endDate string) string {
	endDateStr, _ := helper.HelperDate(endDate)
	billDate := endDateStr.AddDate(0, 0, 1)
	return billDate.Format(constant.LayoutISO)
}

func PopulateCycleYear(startDate string) string {
	startDateStr, _ := helper.HelperDate(startDate)
	cycleYear := startDateStr.Year()
	return strconv.Itoa(cycleYear)
}
