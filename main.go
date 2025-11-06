package main

import (
	"fmt"
	"log"
	"tugas8/database"
	"tugas8/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	routes.UserRoutes(app)
	fmt.Println("🚀 Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
