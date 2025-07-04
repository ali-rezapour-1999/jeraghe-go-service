package publicRouter

import (
	"go-server/service/profile"
	"go-server/types"
)

var ProfileRoutes = []types.PathRouter{
	{
		Path:         "/get-profile/",
		Method:       "GET",
		Handler:      profile.GetProfileData,
		RequiresAuth: true,
	},
}

var ProfileSkillRoutes = []types.PathRouter{
	{
		Path:         "/get-profile-skill/",
		Method:       "GET",
		Handler:      profile.GetProfileSkill,
		RequiresAuth: true,
	},
}

var ProfileSkillItemsRoutes = []types.PathRouter{
	{
		Path:         "/get-profile-skill-items/",
		Method:       "GET",
		Handler:      profile.GetProfileSkillItems,
		RequiresAuth: true,
	},
}

var ProfileSkillItemRoutes = []types.PathRouter{
	{
		Path:         "/get-profile-skill-item/",
		Method:       "POST",
		Handler:      profile.GetProfileSkillItem,
		RequiresAuth: true,
	},
}
