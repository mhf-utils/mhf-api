package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/controllers"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

var item_routes = []Route{
	{
		Endpoint: "/items",
		Handler:  "List",
		Method:   "GET",
	},
	{
		Endpoint: "/items/{id}",
		Handler:  "Read",
		Method:   "GET",
	},
}

func GetRouterItem(router *mux.Router, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:middlewares:item:GetRouterItem")
	controller := controllers.NewControllerItem(log, binary_file)

	for _, route := range item_routes {
		handler := createDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
