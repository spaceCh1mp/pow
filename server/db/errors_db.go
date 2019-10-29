package db

import "errors"

var (
	//ErrNotModified is returned when an update opertaion completes but does not modify any document
	ErrNotModified = errors.New("No Modified Document")

	//ErrNoMatchedDocument is returned when an operation doesn't return any document
	ErrNoMatchedDocument = errors.New("No Matched Documents")

	//ErrInvalidHex is returned when the ObjectId hex is invalid
	ErrInvalidHex = errors.New("Invalid Hex")
)
