package publicRouter

import (
	"go-server/service/base"
	"go-server/types"
)

var CategoryRoutes = []types.PathRouter{
	{
		Path:         "/category/",
		Method:       "GET",
		Handler:      base.GetCategoryService,
		RequiresAuth: false,
	},
}
