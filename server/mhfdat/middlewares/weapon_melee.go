package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/common"
	"mhf-api/server/mhfdat/controllers"
	"mhf-api/server/types"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

func GetWeaponMeleeRoutes() []types.Route {
	return []types.Route{
		{
			Endpoint: "/weapons/melee",
			Handler:  "List",
			Method:   "GET",
		},
		{
			Endpoint: "/weapons/melee/{id}",
			Handler:  "Read",
			Method:   "GET",
		},
	}
}

func GetWeaponMeleeRouter(router *mux.Router, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:mhfdat:middlewares:weapon_melee:GetRouterWeaponMelee")
	controller := controllers.NewControllerWeaponMelee(log, binary_file)

	for _, route := range GetWeaponMeleeRoutes() {
		handler := common.CreateDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
