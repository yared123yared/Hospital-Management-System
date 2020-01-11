package utils

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error ")
	ErrInsufficientBalance = errors.New("Insufficient Balance ")
	ErrAlreadyReserved     = errors.New("Room already reserved in the time you select try another time ")
)
