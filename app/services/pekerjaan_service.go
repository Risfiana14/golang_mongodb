package services

import (
	"context"
	"tugas8/app/model"
	"tugas8/app/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllPekerjaan - ambil semua data pekerjaan
func GetAllPekerjaan(c *fiber.Ctx) error {
	data, err := repository.GetAllPekerjaan(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// GetPekerjaanByID - ambil pekerjaan berdasarkan ID
func GetPekerjaanByID(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repository.GetPekerjaanByID(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
	}
	return c.JSON(data)
}

// GetPekerjaanByAlumniID - ambil semua pekerjaan berdasarkan alumni_id
func GetPekerjaanByAlumniID(c *fiber.Ctx) error {
	alumniID := c.Params("alumni_id")
	data, err := repository.GetPekerjaanByAlumniID(context.Background(), alumniID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// CreatePekerjaan - tambah pekerjaan baru
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


// UpdatePekerjaan - ubah data pekerjaan
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

// DeletePekerjaan - hapus pekerjaan
func DeletePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repository.DeletePekerjaan(context.Background(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Pekerjaan berhasil dihapus"})
}
