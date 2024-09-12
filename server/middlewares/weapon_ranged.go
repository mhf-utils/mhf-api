package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/controllers"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

var weapon_ranged_routes = []Route{
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

func GetRouterWeaponRanged(router *mux.Router, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:middlewares:weapon_ranged:GetRouterWeaponRanged")
	controller := controllers.NewControllerWeaponRanged(log, binary_file)

	for _, route := range weapon_ranged_routes {
		handler := createDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
