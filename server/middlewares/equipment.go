package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/controllers"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

var equipment_routes = []Route{
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

func GetRouterEquipment(log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:middlewares:equipment:GetRouterEquipment")
	controller := controllers.NewControllerEquipment(log, binary_file)
	router := mux.NewRouter()

	for _, route := range equipment_routes {
		handler := createDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
