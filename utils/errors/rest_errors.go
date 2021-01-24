package errors

import "net/http"

type RestErr struct {
	Message string `json:"message""`
	Status  int    `json:"status""`
	Error   string `json:"error""`
}

func NewBadRequestError(bad_request string) *RestErr {
	return &RestErr{
		Message: bad_request,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
