package publicRouter

import (
	"go-server/service/idea"
	"go-server/types"
)

var ListIdeaRoutes = []types.PathRouter{
	{
		Path:         "/get-idea-list/",
		Method:       "GET",
		Handler:      idea.GetUserIdeaListService,
		RequiresAuth: true,
	},
}

var IdeaRoutes = []types.PathRouter{
	{
		Path:         "/get-idea/",
		Method:       "GET",
		Handler:      idea.GetIdeaUserService,
		RequiresAuth: true,
	},
}
