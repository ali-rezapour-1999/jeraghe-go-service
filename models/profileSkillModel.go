package models

type ProfileSkills struct {
	BaseModel
	ProfileId  uint      `gorm:"not null;index" json:"profile_id"`
	Title      string    `gorm:"type:varchar(255);not null" json:"title"`
	CategoryID uint      `gorm:"not null;index" json:"-"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category"`
}

func (*ProfileSkills) TableName() string {
	return "profiles_profileskill"
}
