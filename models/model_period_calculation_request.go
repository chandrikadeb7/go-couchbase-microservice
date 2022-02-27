package models

type PeriodCalculationRequest struct {

	// The assumed cycle/period start day
	Day int32 `json:"day,omitempty"`

	// The assumed cycle/period start month
	Month int32 `json:"month,omitempty"`

	// The assumed cycle/period start year
	Year int32 `json:"year,omitempty"`
}
