package serverErr

import "errors"

var (
	// ErrBadRequest is returned when when there is something wrong with the request itself (i.e. invalid request body)
	// Should be returned as 400 HTTP Status
	ErrBadRequest = errors.New("Bad request")

	// ErrNotAuthenticated is returned when user is not authenticated (authentication token is missing or is invalid)
	// Should be returned as 401
	ErrNotAuthenticated = errors.New("Not authenticated")

	// ErrForbidden is returned when user tries to access parts of system he isn't allowed to
	// Should be returned as 403
	ErrForbidden = errors.New("Forbidden")

	// ErrInvalidAPICall is returned when API call issued to server is invalid (API is not supported or URI is invalid)
	// Should be returned as 404 HTTP Status
	ErrInvalidAPICall = errors.New("Invalid API call")

	// ErrResourceNotFound is returned when API is called with a request for nonexisting system resource
	// Should be returned as 404 HTTP Status
	ErrResourceNotFound = errors.New("Resource not found")

	// ErrMethodNotAllowed is returned when request API does not support HTTP method used for request
	// Should be returned as 405 HTTP Status
	ErrMethodNotAllowed = errors.New("Method not allowed")

	// ErrInternal is returned when something bad happened during processing API request
	// Should be returned as 500 HTTP Status
	ErrInternal = errors.New("Internal error")
)
