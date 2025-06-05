package models

type Tag struct {
	BaseModel
	Title string `gorm:"type:varchar(255);not null" json:"title"`
}

func (*Tag) TableName() string {
	return "base_tags"
}
