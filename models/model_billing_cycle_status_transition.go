package models

// BillingCycleStatusTransition - Status transition for the billing cycle statuses 
type BillingCycleStatusTransition struct {

	// Value of the billing cycle status 
	BillingCycleStatus string `json:"billingCycleStatus,omitempty"`

	// List of statuses eligible for transition 
	TransitionToStatus []string `json:"transitionToStatus,omitempty"`
}
