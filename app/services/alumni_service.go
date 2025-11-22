package services

import (
	"context"
	"tugas8/app/model"
	"tugas8/app/repository"

	"github.com/gofiber/fiber/v2"
)

// GetAllAlumni godoc
// @Summary      Ambil semua data alumni
// @Description  Mengembalikan daftar semua alumni
// @Tags         Alumni
// @Security     BearerAuth
// @Produce      json
// @Success      200 {array} model.Alumni
// @Failure      500 {object} map[string]string
// @Router       /alumni [get]
func GetAllAlumni(c *fiber.Ctx) error {
	data, err := repository.GetAllAlumni(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// GetAlumniByID godoc
// @Summary      Ambil alumni berdasarkan ID
// @Description  Mengembalikan data alumni berdasarkan ID
// @Tags         Alumni
// @Security     BearerAuth
// @Produce      json
// @Param        id path string true "ID Alumni"
// @Success      200 {object} model.Alumni
// @Failure      404 {object} map[string]string
// @Router       /alumni/{id} [get]
func GetAlumniByID(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repository.GetAlumniByID(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
	}
	return c.JSON(data)
}

// CreateAlumni godoc
// @Summary      Tambah alumni baru
// @Description  Membuat data alumni baru
// @Tags         Alumni
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        data body model.Alumni true "Data Alumni"
// @Success      201 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /alumni [post]
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

// UpdateAlumni godoc
// @Summary      Update data alumni
// @Description  Mengubah data alumni berdasarkan ID
// @Tags         Alumni
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path string      true  "ID Alumni"
// @Param        data body model.Alumni true "Data Alumni Baru"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /alumni/{id} [put]
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

// DeleteAlumni godoc
// @Summary      Hapus alumni
// @Description  Menghapus data alumni berdasarkan ID
// @Tags         Alumni
// @Security     BearerAuth
// @Param        id path string true "ID Alumni"
// @Success      200 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /alumni/{id} [delete]
func DeleteAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repository.DeleteAlumni(context.Background(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Alumni berhasil dihapus"})
}