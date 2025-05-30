package base

import (
	"go-server/middleware"
	"go-server/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetCategoryService(c *fiber.Ctx) error {
	db, ok := c.Locals("db").(*gorm.DB)
	if !ok || db == nil {
		err := fiber.NewError(http.StatusInternalServerError, "Database connection unavailable")
		middleware.LogSystemError(db, err, "Database connection unavailable in GetCategoryService")
		return err
	}

	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		middleware.LogSystemError(db, err, "Failed to fetch categories from database")
		return fiber.NewError(http.StatusInternalServerError, "خطا در دریافت داده‌های دسته‌بندی از پایگاه داده")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "اطلاعات دسته‌بندی با موفقیت دریافت شد",
		"data":    categories,
	})
}
