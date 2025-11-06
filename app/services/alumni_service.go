package services

import (
	"context"
	"tugas8/app/model"
	"tugas8/app/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllAlumni - ambil semua alumni
func GetAllAlumni(c *fiber.Ctx) error {
	data, err := repository.GetAllAlumni(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// GetAlumniByID - ambil data alumni berdasarkan ID
func GetAlumniByID(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repository.GetAlumniByID(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
	}
	return c.JSON(data)
}

// CreateAlumni - tambah alumni baru
func CreateAlumni(c *fiber.Ctx) error {
	var input model.Alumni
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}
	if err := repository.CreateAlumni(context.Background(), input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Alumni berhasil ditambahkan"})
}

// UpdateAlumni - ubah data alumni
func UpdateAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	var input model.Alumni
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}
	if err := repository.UpdateAlumni(context.Background(), id, input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Alumni berhasil diperbarui"})
}

// DeleteAlumni - hapus alumni
func DeleteAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repository.DeleteAlumni(context.Background(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Alumni berhasil dihapus"})
}
