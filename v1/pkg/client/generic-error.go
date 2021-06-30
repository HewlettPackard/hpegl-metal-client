// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

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
