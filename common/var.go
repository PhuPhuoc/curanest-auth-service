package common

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEmailNotFound  = errors.New("email not found")
	ErrWrongPassword  = errors.New("wrong password")
)
