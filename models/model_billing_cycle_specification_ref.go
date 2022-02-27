package models

// BillingCycleSpecificationRef - Information about the billing cycle initiation and various attributes of the billing cycle
type BillingCycleSpecificationRef struct {

	// Unique identifier of the billing cycle specification
	Id string `json:"id" validate:"required"`

	// Hypertext reference to the billing cycle specification
	//Href string `json:"href,omitempty"`

	// Frequency of the billing cycle (for example, monthly)
	Frequency string `json:"frequency" validate:"required"`
}
