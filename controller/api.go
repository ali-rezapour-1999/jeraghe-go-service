package controller

import (
	"go-server/middleware"
	"go-server/proxys"
	"go-server/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ControllerRoute(app *fiber.App, db *gorm.DB) {
	app.Use(middleware.DBMiddleware(db))

	Group := app.Group("/api/public")
	registerRoutes(Group, GetAllRoutes())

	protectionRules := routes.GetProtectionRules()
	app.All("/api/private/*", middleware.ConditionalAuth(protectionRules), proxys.ProxyToDjango(db))
}
