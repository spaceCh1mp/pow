package utils

import "errors"

//Error definition for this project are made here

var (
	// ErrED is returned when an empty dataset is either passed or gotten as a response
	ErrED = errors.New("Empty Dataset")

	// ErrFC is returned when the operation fails due to a bad connection
	ErrFC = errors.New("Could not establish connection")

	// ErrFD is returned when the delete operation fails
	ErrFD = errors.New("Failed to delete data")

	// ErrMF is returned if there's a missing field in the dataset that was passed
	ErrMF = errors.New("Missing Field")

	// ErrWE is returnes if an Insert operation fails
	ErrWE = errors.New("Could not Insert data")

	// ErrID is returned if the object.ID passe is invalid or empty
	ErrID = errors.New("_id is either invalid or empty")

	// ErrInternalServerError is returned if an operation during IPCs
	ErrInternalServerError = errors.New("Internal Server Error")
)
