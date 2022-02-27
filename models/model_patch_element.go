package models

// PatchElement - *PatchElement* object 
type PatchElement struct {

	// Operation for this *PatchElement*. Valid value&#58; *replace*. 
	Op string `json:"op,omitempty"`

	// Path to this *PatchElement* (for example, *_/status*) 
	Path string `json:"path,omitempty"`

	// Value of this *PatchElement* 
	Value string `json:"value,omitempty"`
}
