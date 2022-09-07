package types

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type ServiceError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ServiceError) Error() string {
	return e.Message
}

func NewServiceError(message string, code ...int) error {
	cd := 0
	if len(code) > 0 {
		cd = code[0]
	}

	serr := &ServiceError{
		Code:    cd,
		Message: message,
	}
	jbytes, err := json.Marshal(serr)
	if err != nil {
		return err
	}

	return fiber.NewError(fiber.StatusBadRequest, string(jbytes))
}

type Page[T any] struct {
	Total int64 `json:"total"`
	List  []T   `json:"list"`
	Extra any   `json:"extra"`
}
