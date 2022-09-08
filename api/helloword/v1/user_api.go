package v1

import (
	"fastgo/internal/common/valid"
	"github.com/gofiber/fiber/v2"
)

type Job struct {
	Type   int `json:"type" validate:"required,min=3,max=32"`
	Salary int `json:"salary" validate:"required,number"`
}

type User struct {
	Name string `validate:"required,min=3,max=32"`
	// use `*bool` here otherwise the validation will fail for `false` values
	// Ref: https://github.com/go-playground/validator/issues/319#issuecomment-339222389
	IsActive *bool  `validate:"required"`
	Email    string `validate:"required,email,min=6,max=32"`
	Job      Job    `json:"job" validate:"dive"`
}

func AddUser(c *fiber.Ctx) error {
	//Connect to database
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}
	errors := valid.ValidateStruct(*user)
	if errors != nil {
		return errors
	}

	//Do something else here

	//Return user
	return c.JSON(user)
}
