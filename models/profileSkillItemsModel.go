package models

import "go-server/models/enums"

type ProfileSkillItem struct {
	BaseModel
	ProfileId uint                   `gorm:"not null;index" json:"profile_id"`
	Skill     uint                   `gorm:"column:skill_id;not null;index" json:"skill_id"`
	Level     enums.ProficiencyLevel `gorm:"type:VARCHAR(20);not null" json:"level"`

	Tag Tag `gorm:"foreignKey:Skill;references:ID" json:"tag"`
}

func (ProfileSkillItem) TableName() string {
	return "profiles_skillitems"
}
