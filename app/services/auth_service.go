package services

import (
	"context"
	"tugas8/app/model"
	"tugas8/app/repository"
	"tugas8/utils"

	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary      Login user
// @Description  Login dengan username dan password, mengembalikan JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body model.LoginRequest true "Login credentials"
// @Success      200 {object} model.LoginResponse
// @Failure      400 {object} map[string]string "Invalid request"
// @Failure      401 {object} map[string]string "Username atau password salah"
// @Router       /login [post]
func Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid JSON"})
	}

	if req.Username == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Username dan password wajib diisi"})
	}

	user, err := repository.FindUserByUsername(context.Background(), req.Username)
	if err != nil || user == nil {
		return c.Status(401).JSON(fiber.Map{"message": "Username atau password salah"})
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return c.Status(401).JSON(fiber.Map{"message": "Username atau password salah"})
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal buat token"})
	}

	return c.JSON(model.LoginResponse{
		Token: token,
		User:  *user,
	})
}