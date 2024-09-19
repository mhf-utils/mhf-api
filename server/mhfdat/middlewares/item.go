package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/common"
	"mhf-api/server/mhfdat/controllers"
	"mhf-api/server/types"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

func GetItemRoutes() []types.Route {
	return []types.Route{
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
}

func GetItemRouter(router *mux.Router, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:mhfdat:middlewares:item:GetRouterItem")
	controller := controllers.NewControllerItem(log, binary_file)

	for _, route := range GetItemRoutes() {
		handler := common.CreateDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
