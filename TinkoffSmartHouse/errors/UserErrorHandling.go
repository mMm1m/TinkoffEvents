package errors

import (
	"TinkoffSmartHouse/constants"
	"fmt"
)

type APIError struct {
	StatusCode int `json:"status_code"`
	Message    any `json:"message"`
}

func (err APIError) Error() string {
	return fmt.Sprintf("API error catched: %v", err.Message)
}

func NewAPIError(code int, err error) APIError {
	return APIError{
		StatusCode: code,
		Message:    err.Error(),
	}
}

func IncorrectCRC8Code() APIError {
	return NewAPIError(constants.CRC8Error, fmt.Errorf("Incorrect computing of crc8 code"))
}

func EmptyArrayOfBytes() APIError {
	return NewAPIError(constants.ZeroSize, fmt.Errorf("Empty array in input"))
}

func IncorrectSize() APIError {
	return NewAPIError(constants.IncorrectSize, fmt.Errorf("Size of byte array does not match to declared"))
}
