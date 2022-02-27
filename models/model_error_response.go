package models

// ErrorResponse - Error response. This object is used when an API throws an error, typically with HTTP error response codes 3xx, 4xx, and 5xx. 
type ErrorResponse struct {

	// Error code relevant to an application, defined in the API or in a common list 
	Code string `json:"code"`

	// Reason for the error that can be displayed to a user 
	Reason string `json:"reason"`

	// Additional information about the error and corrective actions related to the error that can be displayed to a user 
	Message string `json:"message,omitempty"`

	// HTTP error code extension 
	Status string `json:"status,omitempty"`

	// URI of the documentation describing the error 
	ReferenceError string `json:"referenceError,omitempty"`

	// Unique tracker ID that is used to facilitate the troubleshooting 
	TraceId string `json:"traceId,omitempty"`

	// Type of the error 
	Type string `json:"@type,omitempty"`

	// Characteristic parameters that are associated with the error 
	Characteristic []ErrorCharacteristic `json:"characteristic,omitempty"`
}
