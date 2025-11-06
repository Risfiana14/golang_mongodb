package services

import (
	"context"
	"fmt"
	"tugas8/app/model"
	"tugas8/app/repository"
	"tugas8/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var loginData model.LoginRequest
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"success": false,
		})
	}

	user, err := repository.FindUserByUsername(context.Background(), loginData.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username tidak ditemukan",
			"success": false,
		})
	}

	// 🟡 Debug (sementara)
	fmt.Println(">> Password dari user:", loginData.Password)
	fmt.Println(">> Hash dari DB:", user.PasswordHash)

	if !utils.CheckPasswordHash(loginData.Password, user.PasswordHash) {
		fmt.Println(">> CheckPasswordHash: FALSE")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Password salah",
			"success": false,
		})
	}

	fmt.Println(">> CheckPasswordHash: TRUE")

	token, err := utils.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat token",
			"success": false,
		})
	}

	return c.JSON(model.LoginResponse{
		Token: token,
		User:  user,
	})
}
