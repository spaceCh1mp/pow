package user

import "errors"

//Errors for this package are defined here

var (
	errED = errors.New("Empty Dataset: NewUser")
	errMF = errors.New("Missing Field: FirstName")
	errML = errors.New("Missing Field: LastName")
	errME = errors.New("Missing Field: Email")
	errMP = errors.New("Missing Field: Password")
	errFC = errors.New("Could not establish connection")
	errWE = errors.New("Could not Insert data")
)
