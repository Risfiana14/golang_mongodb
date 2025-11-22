package services

import "github.com/gofiber/fiber/v2"

type userService struct{}

var UserService = &userService{}

// HandleGetAllUsers godoc
// @Summary Dapatkan semua user
// @Description Mengambil daftar semua user dari database
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func (s *userService) HandleGetAllUsers(c *fiber.Ctx) error {
	return c.JSON([]map[string]interface{}{
		{"id": "1", "name": "User 1"},
		{"id": "2", "name": "User 2"},
	})
}

// HandleGetUserByID godoc
// @Summary Dapatkan user berdasarkan ID
// @Description Mengambil data user spesifik berdasarkan ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [get]
func (s *userService) HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(map[string]interface{}{
		"id":   id,
		"name": "User " + id,
	})
}

// HandleCreateUser godoc
// @Summary Buat user baru
// @Description Membuat user baru di database
// @Tags Users
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "User data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [post]
func (s *userService) HandleCreateUser(c *fiber.Ctx) error {
	body := make(map[string]interface{})
	c.BodyParser(&body)
	return c.Status(201).JSON(body)
}

// HandleUpdateUser godoc
// @Summary Update user
// @Description Memperbarui data user berdasarkan ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param body body map[string]interface{} true "User data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [put]
func (s *userService) HandleUpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	body := make(map[string]interface{})
	c.BodyParser(&body)
	body["updated_id"] = id
	return c.JSON(body)
}

// HandleDeleteUser godoc
// @Summary Hapus user
// @Description Menghapus user berdasarkan ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [delete]
func (s *userService) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(map[string]interface{}{
		"message": "User deleted",
		"id":      id,
	})
}
