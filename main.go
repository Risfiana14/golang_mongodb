// @title           TUGAS8 - Sistem Informasi Alumni
// @version         1.0
// @description     API untuk manajemen alumni, pekerjaan, dan upload file
// @host            localhost:3000
// @BasePath       /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token dengan format: Bearer {token}

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	_ "tugas8/docs"
	"tugas8/database"
	"tugas8/app/repository"  // ← TAMBAHKAN INI
	"tugas8/routes"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	_ = godotenv.Load()

	// 1. Koneksi MongoDB
	database.Connect()

	// 2. Inisialisasi collection SETELAH koneksi berhasil ← INI YANG PENTING!
	repository.InitCollections()

	app := fiber.New()

	// CORS
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(204)
		}
		return c.Next()
	})

	// Swagger UI
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// API Routes
	api := app.Group("/api")
	routes.UserRoutes(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server jalan → http://localhost:" + port)
	log.Println("Swagger UI  → http://localhost:" + port + "/swagger/index.html")
	log.Fatal(app.Listen(":" + port))
}