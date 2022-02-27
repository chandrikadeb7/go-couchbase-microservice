package models

// BillingCycleRef - Information regarding the billing cycle  
type BillingCycleRef struct {

	// Unique identifier of the billing cycle 
	Id string `json:"id,omitempty"`

	// Cycle Instance 
	CycleInstance string `json:"cycleInstance,omitempty"`
}
