package handler

import (
	"backend/model"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Type     string `json:"type"`
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Error on login request",
			"data":    err,
		})
	}

	category := input.Type
	identity := input.Identity
	pass := input.Password
	if category == "worker" {
		userModel, err := new(model.Worker), *new(error)
	} else if category == "poster" {
		userModel, err := new(model.Poster), *new(error)
	}

}
