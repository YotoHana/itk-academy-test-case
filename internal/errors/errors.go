package errors

import "errors"

var (
	ErrInvalidAmount        = errors.New("invalid amount")
	ErrInvalidOperationType = errors.New("invalid operation type")
	ErrUnsufficientBalance  = errors.New("unsufficient balance")
	ErrWalletNotFound       = errors.New("wallet not found")
	ErrInvalidPathParameter = errors.New("invalid path parameter")
)
