package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/common"
	"mhf-api/server/mhfdat/controllers"
	"mhf-api/server/types"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

func GetEquipmentRoutes() []types.Route {
	return []types.Route{
		{
			Endpoint: "/equipments/{type}",
			Handler:  "List",
			Method:   "GET",
		},
		{
			Endpoint: "/equipments/{type}/{id}",
			Handler:  "Read",
			Method:   "GET",
		},
	}
}

func GetEquipmentRouter(router *mux.Router, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:mhfdat:middlewares:equipment:GetRouterEquipment")
	controller := controllers.NewControllerEquipment(log, binary_file)

	for _, route := range GetEquipmentRoutes() {
		handler := common.CreateDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
