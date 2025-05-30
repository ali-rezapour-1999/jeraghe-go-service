package publicRouter

import (
	"go-server/service/base"
	"go-server/types"
)

var ContractRoutes = []types.PathRouter{
	{
		Path:         "/contact/",
		Method:       "GET",
		Handler:      base.GetContactUserService,
		RequiresAuth: true,
	},
}
