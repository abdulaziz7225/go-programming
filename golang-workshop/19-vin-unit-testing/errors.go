package vin

import "errors"

var (
	ErrInvalidVINLength = errors.New("VIN must be 17 characters long")
	ErrEmptyVIN         = errors.New("VIN cannot be empty")
	ErrInvalidCharacter = errors.New("VIN contains invalid characters")
)
