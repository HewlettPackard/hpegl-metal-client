// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package model

// ErrorResponse used to return the failures of REST APIs
// Per GL API standards, the error message must contain details of the problem and must contain fields like
// 'message', 'recommendedActions', 'details', 'errorCode' in the response.
// GL API standards - https://developer.greenlake.hpe.com/docs/greenlake/api/internal/http/#error-handling
// For now, limiting the error response to send only 'messages', 'details' (optional) and, deferring
// adding remaining fields for the future.
type ErrorResponse struct {
	// message: clear and concise description of the error condition, suitable for display to an end user
	Message string `json:"message"`
	//details: optional verbose description of the error condition, suitable for display to an end user
	Details string `json:"details,omitempty"`
}
