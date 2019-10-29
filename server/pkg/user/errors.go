package user

import "errors"

//Errors for this package are defined here

var (
	errED = errors.New("Empty Dataset: NewUser")
	errFC = errors.New("Could not establish connection")
	errFD = errors.New("Failed to delete data")
	errME = errors.New("Missing Field: Email")
	errMF = errors.New("Missing Field: FirstName")
	errML = errors.New("Missing Field: LastName")
	errMP = errors.New("Missing Field: Password")
	errUN = errors.New("Missing Field: UserName")
	errWE = errors.New("Could not Insert data")

	errID = errors.New("_id is either invalid or empty")
)
