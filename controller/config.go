package controller

import (
	"go-server/middleware"
	"go-server/types"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(router fiber.Router, routes []types.PathRouter) {
	for _, route := range routes {
		var handlers []fiber.Handler

		if route.RequiresAuth {
			handlers = append(handlers, middleware.JWTMiddleware)
		}

		handlers = append(handlers, route.Handler)

		switch route.Method {
		case "GET":
			router.Get(route.Path, handlers...)
		default:
			router.All(route.Path, handlers...)
		}
	}
}
