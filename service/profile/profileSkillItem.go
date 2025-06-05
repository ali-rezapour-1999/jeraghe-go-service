package profile

import (
	"context"
	"errors"
	"go-server/middleware"
	"go-server/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProfileSkillItem(c *fiber.Ctx) error {
	db, ok := c.Locals("db").(*gorm.DB)
	if !ok || db == nil {
		middleware.LogSystemError(db, nil, "Database connection unavailable in GetProfileSkill")
		return fiber.NewError(http.StatusInternalServerError, "خطا در دریافت داده‌های پروفایل از پایگاه داده")
	}

	ctx := context.Background()

	authenticated, userID, err := middleware.IsAuthenticated(c)
	if !authenticated || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "احراز هویت ناموفق بود",
			"data":    nil,
			"error":   err.Error(),
		})
	}

	var profileSkillItem models.ProfileSkillItem

	if err := c.BodyParser(&profileSkillItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	id := profileSkillItem.ID

	if err := db.WithContext(ctx).
		Preload("Tag").
		Joins("JOIN profiles_profile pp ON profiles_skillitems.profile_id = pp.id").
		Joins("left join base_tags bt on bt.id = profiles_skillitems.skill_id").
		Where("pp.user_id = ? AND profiles_skillitems.id = ?", userID, id).
		First(&profileSkillItem).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(fiber.Map{
				"status": "success",
				"data":   nil,
			})
		}

		middleware.LogSystemError(db, err, "Failed to fetch profile skills from database")
		return fiber.NewError(http.StatusInternalServerError, "خطا در دریافت داده‌های پروفایل از پایگاه داده")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "اطلاعات پروفایل با موفقیت دریافت شد",
		"data":    profileSkillItem,
	})
}
