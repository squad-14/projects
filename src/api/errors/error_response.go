package errors

import (
	"fmt"
)

//ErrResponse struct
type ErrResponse struct {
	ErrCode string
	Message string
}

//NewErrResponse constructor.
func NewErrResponse(err error, errCode string) *ErrResponse {
	return &ErrResponse{ErrCode: errCode, Message: err.Error()}
}

//Error taking the cause.
func (e *ErrResponse) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("Cause error %v", e.Error())
}
