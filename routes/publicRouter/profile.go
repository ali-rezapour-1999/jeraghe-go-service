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
