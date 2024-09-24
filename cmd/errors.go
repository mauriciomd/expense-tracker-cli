package cmd

import "errors"

var (
	ErrInvalidDescription = errors.New("an expense description was not informed")
	ErrInvalidAmount      = errors.New("amount shoud be greater than zero")
	ErrInvalidId          = errors.New("id must be greater than zero")
	ErrExpenseNotFound    = errors.New("expense not found")
	ErrInvalidMonth       = errors.New("month must be between 1 and 12")
)
