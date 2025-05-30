package enum

import (
	"go-server/service/enums"
	"go-server/types"
)

var WorkExperienceLevelRoutes = []types.PathRouter{
	{
		Path:         "/enum/work-experience-levels/",
		Method:       "GET",
		Handler:      enums.GetProficiencyEnum,
		RequiresAuth: false,
	},
}
