package main

import (
	"github.com/gofiber/fiber/v2" // Import the Fiber framework
	"github.com/gofiber/fiber/v2/middleware/cors" // Import the CORS middleware
	routes "github.com/heronhoga/sapa-usaha-api/routes" // Import the routes package
	dbConfig "github.com/heronhoga/sapa-usaha-api/config"
)

func main() {
	app := fiber.New()

	routes.Setup(app)
	app.Use(cors.New())

	dbConfig.DbConnectionMigration()

	app.Listen(":3000")
}