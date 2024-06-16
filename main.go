package main

import (
	"os"

	"github.com/secrethook/backend/pkg/configs"
	"github.com/secrethook/backend/pkg/middleware"
	"github.com/secrethook/backend/pkg/routes"
	"github.com/secrethook/backend/pkg/utils"
	"github.com/secrethook/backend/platform/database"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.ConnectToMongodb()

	config := configs.FiberConfig()

	app := fiber.New(config)

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, POST request received")
	})

	middleware.FiberMiddleware(app)
	
	routes.WebhookRoutes(app.Group("/api/v1/webhook"))
	
	routes.NotFoundRoute(app)
	
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
