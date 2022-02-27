package models

// BillingCycle - Cycle for a specific bill cycle specification and a specific bill run
type BillingCycle struct {

	// Unique identifier of the billing cycle
	Id string `json:"id,omitempty"`

	// Hypertext reference to the billing cycle
	//Href string `json:"href,omitempty"`

	// Start date of the billing cycle
	StartDate string `json:"startDate" validate:"required"`

	// End date of the billing cycle
	EndDate string `json:"endDate" validate:"required"`

	// Frequency of the billing cycle (for example, monthly)
	Frequency string `json:"frequency" validate:"required"`

	// Date on which the end-of-cycle process started running for this billing cycle
	BillingDate string `json:"billingDate,omitempty"`

	// Status of the bill cycle run
	Status string `json:"status,omitempty"`

	// Identifier of the billing cycle in an external system
	ExternalId string `json:"externalId,omitempty"`

	// Indicates whether the cycle is locked for rate changes
	//LockedForRateChange string `json:"lockedForRateChange,omitempty"`

	// Serial occurrence of the cycle instance in a calendar year. For example, for a monthly cycle, it contains the month of the year, and for a quarterly cycle, it contains a number between 1 and 4.
	CycleInstance string `json:"cycleInstance,omitempty"`

	// Year of the cycle end date
	CycleYear string `json:"cycleYear,omitempty"`

	//PrevBillingCycle BillingCycleRef `json:"prevBillingCycle,omitempty"`

	//NextBillingCycle BillingCycleRef `json:"nextBillingCycle,omitempty"`

	CycleSpecification BillingCycleSpecificationRef `json:"cycleSpecification,omitempty"`
}
