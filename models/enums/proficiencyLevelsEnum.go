package enums

type ProficiencyLevel string

const (
	NOVICE       ProficiencyLevel = "NOVICE"
	BEGINNER     ProficiencyLevel = "BEGINNER"
	INTERMEDIATE ProficiencyLevel = "INTERMEDIATE"
	ADVANCED     ProficiencyLevel = "ADVANCED"
	EXPERT       ProficiencyLevel = "EXPERT"
)

type EnumItem struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

func GetProficiencyLevels() []EnumItem {
	return []EnumItem{
		{Value: string(NOVICE), Label: "تازه‌کار / آشنایی اولیه (Novice)"},
		{Value: string(BEGINNER), Label: "مبتدی (Beginner)"},
		{Value: string(INTERMEDIATE), Label: "متوسط (Intermediate)"},
		{Value: string(ADVANCED), Label: "پیشرفته (Advanced)"},
		{Value: string(EXPERT), Label: "تحصیل (Expert)"},
	}
}
