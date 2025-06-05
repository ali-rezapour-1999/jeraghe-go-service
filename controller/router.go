package controller

import (
	"go-server/routes/enum"
	"go-server/routes/publicRouter"
	"go-server/types"
)

func GetAllRoutes() []types.PathRouter {
	var allRoutes []types.PathRouter

	// base
	allRoutes = append(allRoutes, publicRouter.CategoryRoutes...)
	allRoutes = append(allRoutes, publicRouter.ContractRoutes...)

	// profile
	allRoutes = append(allRoutes, publicRouter.ProfileRoutes...)
	allRoutes = append(allRoutes, publicRouter.ProfileSkillRoutes...)
	allRoutes = append(allRoutes, publicRouter.ProfileSkillItemsRoutes...)
	allRoutes = append(allRoutes, publicRouter.ProfileSkillItemRoutes...)

	// idea
	allRoutes = append(allRoutes, publicRouter.IdeaRoutes...)
	allRoutes = append(allRoutes, publicRouter.ListIdeaRoutes...)

	allRoutes = append(allRoutes, enum.WorkExperienceLevelRoutes...)

	return allRoutes
}
