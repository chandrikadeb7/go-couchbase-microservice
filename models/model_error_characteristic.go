package models

// ErrorCharacteristic - Parameter that is associated with the error entity 
type ErrorCharacteristic struct {

	// Non-localized name identifier of the characteristic 
	Name string `json:"name,omitempty"`

	// Value of the characteristic 
	Value string `json:"value,omitempty"`
}
