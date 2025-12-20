package atm

import "errors"

var (
	ErrInsufficientFunds     = errors.New("insufficient funds")
	ErrInsufficientCashInATM = errors.New("insufficient cash in ATM")
	ErrInvalidAmount         = errors.New("invalid amount")
	ErrInvalidPIN            = errors.New("invalid PIN")
	ErrAccountNotFound       = errors.New("account not found")
	ErrCardNotFound          = errors.New("card not found")
)
