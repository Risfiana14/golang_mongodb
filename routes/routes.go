package routes

import (
	"tugas8/app/services"
	"tugas8/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Login & auth middleware (gunakan token lama)
	api.Post("/login", services.Login)

	protected := api.Group("", middleware.AuthRequired())

	// CRUD ALUMNI
	protected.Get("/alumni", services.GetAllAlumni)
	protected.Get("/alumni/:id", services.GetAlumniByID)
	protected.Post("/alumni", services.CreateAlumni)
	protected.Put("/alumni/:id", services.UpdateAlumni)
	protected.Delete("/alumni/:id", services.DeleteAlumni)

	// CRUD PEKERJAAN
	protected.Get("/pekerjaan", services.GetAllPekerjaan)
	protected.Get("/pekerjaan/:id", services.GetPekerjaanByID)
	protected.Get("/pekerjaan/alumni/:alumni_id", services.GetPekerjaanByAlumniID)
	protected.Post("/pekerjaan", services.CreatePekerjaan)
	protected.Put("/pekerjaan/:id", services.UpdatePekerjaan)
	protected.Delete("/pekerjaan/:id", services.DeletePekerjaan)

	// ---------- FILE UPLOAD ----------
	protected.Post("/upload", services.UploadFile)
	protected.Get("/upload", services.GetAllFiles)
	protected.Get("/upload/:id", services.GetFileByID)
	protected.Delete("/upload/:id", services.DeleteFile)

}

