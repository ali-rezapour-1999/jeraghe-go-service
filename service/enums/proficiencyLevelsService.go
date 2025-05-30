package enums

import (
	"go-server/models/enums"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetProficiencyEnum(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "لیست سطوح مهارت",
		"data":    enums.GetProficiencyLevels(),
	})
}
