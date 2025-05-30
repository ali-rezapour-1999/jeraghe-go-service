package models

type Category struct {
	BaseModel
	Title string `gorm:"type:varchar(255);not null" json:"title"`
}

func (Category) TableName() string {
	return "base_category"
}
