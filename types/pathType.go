package types

import (
	"github.com/gofiber/fiber/v2"
)

type PathRouter struct {
	Path         string
	Method       string
	Handler      fiber.Handler
	RequiresAuth bool
}
