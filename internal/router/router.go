package router

import (
	"mystats-server/internal/auth"
	"mystats-server/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(h *handler.Handler, app *fiber.App) {
	// Public routes
	app.Post("/auth/login", h.Login)
	app.Post("/auth/signup", h.Signup)
	app.Get("/auth/validate", h.ValidateToken)

	// Test routes (개발 환경에서만 사용)
	app.Post("/test/create-user", h.CreateTestUser)
	app.Post("/test/create-game", h.CreateTestGame)

	api := app.Group("/api")
	api.Use(auth.AuthMiddleware())
	{
		api.Get("/profile", h.GetProfile)
		api.Get("/records", h.GetRecords)
		api.Post("/records/batter", h.CreateBatterRecord)
		api.Post("/records/pitcher", h.CreatePitcherRecord)
	}
}
