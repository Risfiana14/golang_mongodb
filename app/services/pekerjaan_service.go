package services

import (
	"context"
	"tugas8/app/model"
	"tugas8/app/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllPekerjaan godoc
// @Summary      Ambil semua data pekerjaan
// @Tags         Pekerjaan
// @Security     BearerAuth
// @Produce      json
// @Success      200 {array} model.Pekerjaan
// @Failure      500 {object} map[string]string
// @Router       /pekerjaan [get]
func GetAllPekerjaan(c *fiber.Ctx) error {
	data, err := repository.GetAllPekerjaan(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// GetPekerjaanByID godoc
// @Summary      Ambil pekerjaan berdasarkan ID
// @Tags         Pekerjaan
// @Security     BearerAuth
// @Produce      json
// @Param        id path string true "ID Pekerjaan"
// @Success      200 {object} model.Pekerjaan
// @Failure      404 {object} map[string]string
// @Router       /pekerjaan/{id} [get]
func GetPekerjaanByID(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repository.GetPekerjaanByID(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
	}
	return c.JSON(data)
}

// GetPekerjaanByAlumniID godoc
// @Summary      Ambil pekerjaan berdasarkan alumni_id
// @Tags         Pekerjaan
// @Security     BearerAuth
// @Produce      json
// @Param        alumni_id path string true "ID Alumni"
// @Success      200 {array} model.Pekerjaan
// @Failure      500 {object} map[string]string
// @Router       /pekerjaan/alumni/{alumni_id} [get]
func GetPekerjaanByAlumniID(c *fiber.Ctx) error {
	alumniID := c.Params("alumni_id")
	data, err := repository.GetPekerjaanByAlumniID(context.Background(), alumniID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// CreatePekerjaan godoc
// @Summary      Tambah pekerjaan baru
// @Tags         Pekerjaan
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        data body model.Pekerjaan true "Data Pekerjaan"
// @Success      201 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Router       /pekerjaan [post]
func CreatePekerjaan(c *fiber.Ctx) error {
	var input model.Pekerjaan
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	_, err := repository.CreatePekerjaan(context.Background(), input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Pekerjaan berhasil ditambahkan",
	})
}

// UpdatePekerjaan godoc
// @Summary      Update pekerjaan
// @Tags         Pekerjaan
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path string         true  "ID Pekerjaan"
// @Param        data body model.Pekerjaan true "Data Pekerjaan Baru"
// @Success      200 {object} map[string]string
// @Router       /pekerjaan/{id} [put]
func UpdatePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	var input model.Pekerjaan
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	_, err := repository.UpdatePekerjaan(context.Background(), id, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Pekerjaan berhasil diperbarui",
	})
}

// DeletePekerjaan godoc
// @Summary      Hapus pekerjaan
// @Tags         Pekerjaan
// @Security     BearerAuth
// @Param        id path string true "ID Pekerjaan"
// @Success      200 {object} map[string]string
// @Router       /pekerjaan/{id} [delete]
func DeletePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repository.DeletePekerjaan(context.Background(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Pekerjaan berhasil dihapus"})
}