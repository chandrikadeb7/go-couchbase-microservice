package models

// PeriodCalculationResponse - The calculated cycle/period start date and end date 
type PeriodCalculationResponse struct {

	// The calculated cycle/period start date. In most cases it will be the same date as in the input. 
	StartDate string `json:"startDate,omitempty"`

	// The calculated cycle/period end date 
	EndDate string `json:"endDate,omitempty"`
}
