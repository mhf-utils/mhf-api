package mhfdat

import (
	"github.com/gorilla/mux"

	"mhf-api/server/common"
	"mhf-api/server/mhfdat/middlewares"
	"mhf-api/server/types"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

func GetRoutes() []types.Route {
	var routes []types.Route

	routes = append(routes, middlewares.GetEquipmentRoutes()...)
	routes = append(routes, middlewares.GetItemRoutes()...)
	routes = append(routes, middlewares.GetQuestRoutes()...)
	routes = append(routes, middlewares.GetWeaponMeleeRoutes()...)
	routes = append(routes, middlewares.GetWeaponRangedRoutes()...)

	return routes
}

func GetRouter(router *mux.Router, prefix string, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:mhfdat:middlewares:GetRouter")

	common.DisplayAvailableRoutes(GetRoutes(), prefix)

	middlewares.GetEquipmentRouter(router, log, binary_file)
	middlewares.GetItemRouter(router, log, binary_file)
	middlewares.GetQuestRouter(router, log, binary_file)
	middlewares.GetWeaponMeleeRouter(router, log, binary_file)
	middlewares.GetWeaponRangedRouter(router, log, binary_file)

	return router
}
