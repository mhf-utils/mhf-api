package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/controllers"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

var weapon_melee_routes = []Route{
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

func GetRouterWeaponMelee(router *mux.Router, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:middlewares:weapon_melee:GetRouterWeaponMelee")
	controller := controllers.NewControllerWeaponMelee(log, binary_file)

	for _, route := range weapon_melee_routes {
		handler := createDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
