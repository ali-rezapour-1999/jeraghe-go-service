package models

type Contact struct {
	BaseModel
	UserID     uint   `gorm:"not null;index" json:"user_id"`
	SlugID     string `gorm:"type:varchar(100);unique;not null" json:"slug_id"`
	IsActive   bool   `gorm:"default:true" json:"is_active"`
	Platform   string `gorm:"type:varchar(50);not null" json:"platform"`
	Link       string `gorm:"type:varchar(255);not null" json:"link"`
	IsVerified bool   `gorm:"default:false" json:"is_verified"`
}
