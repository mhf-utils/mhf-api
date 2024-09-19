package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/common"
	"mhf-api/server/mhfdat/controllers"
	"mhf-api/server/types"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

func GetWeaponRangedRoutes() []types.Route {
	return []types.Route{
		{
			Endpoint: "/weapons/ranged",
			Handler:  "List",
			Method:   "GET",
		},
		{
			Endpoint: "/weapons/ranged/{id}",
			Handler:  "Read",
			Method:   "GET",
		},
	}
}

func GetWeaponRangedRouter(router *mux.Router, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:mhfdat:middlewares:weapon_ranged:GetRouterWeaponRanged")
	controller := controllers.NewControllerWeaponRanged(log, binary_file)

	for _, route := range GetWeaponRangedRoutes() {
		handler := common.CreateDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
