package constant

import (
	"errors"
)

var errorCodeMap = map[string]string{
	"BSS-BILL-CBCI-40000001": "Missing mandatory input parameters -",
}

type errorString struct {
	Code    string
	Message string
}

func GetMessageForErrorCode(code string) string {
	switch code {
	case "BSS-BILL-CBCI-40000001":
		return errorCodeMap[code]
	default:
		return "code not declared"
	}
}

func New(errCode string, message error) interface{} {
	return &errorString{
		Code:    errCode,
		Message: GetMessageForErrorCode(errCode) + message.Error(),
	}
}

func (e *errorString) Error() interface{} {
	return &errorString{
		Code:    e.Code,
		Message: e.Message,
	}
}

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

const (
	LayoutISO  = "2006-01-02"
	LayoutUS   = "January 2, 2006"
	Underscore = "_"
)
