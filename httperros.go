package httperors

import (
	"net/http"
)
var HttperrorHnadler HttpInterface = &httperrorhandler{}

type HttpInterface interface{
	NewBadRequestError(message string) *HttpError
	NewNotFoundError(message string) *HttpError
	NewSuccessMessage(message string) *HttpSuccess
}
type httperrorhandler struct{

}
func (h httperrorhandler)NewBadRequestError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad request",
	}
}
func (h httperrorhandler)NewNotFoundError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "Not Found",
	}
}
func (h httperrorhandler)NewSuccessMessage(message string) *HttpSuccess {
	return &HttpSuccess{
		Message: message,
		Code:    http.StatusOK,
		Error:   "Delete success",
	}
}