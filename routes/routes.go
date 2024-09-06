package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/heronhoga/sapa-usaha-api/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Test)
}
