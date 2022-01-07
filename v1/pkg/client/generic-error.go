// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

package client

// NewGenericOpenAPIError creates a GenericOpenAPIError from the provided
// parameters.
func NewGenericOpenAPIError(body []byte, error string, model interface{}) GenericOpenAPIError {
	return GenericOpenAPIError{
		body:  body,
		error: error,
		model: model,
	}
}

// Message returns the short error message for display purpose
// If e.model is set and is of type ErrorResponse, then returns the value in ErrorResponse.Message,
// otherwise, returns the entire response body in string format.  
func (e GenericOpenAPIError) Message() string {
	msg := string(e.Body())

	model := e.Model()
	if model != nil {
		res, ok := model.(ErrorResponse)
		if ok {
			msg = res.Message
		}
	}
	return msg
}
