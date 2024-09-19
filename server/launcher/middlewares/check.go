package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/common"
	"mhf-api/server/launcher/controllers"
	"mhf-api/server/types"
	"mhf-api/utils/logger"
)

func GetCheckRoutes() []types.Route {
	return []types.Route{
		{
			Endpoint: "/checks",
			Handler:  "CheckFiles",
			Method:   "GET",
		},
	}
}

func GetCheckRouter(router *mux.Router, log *logger.Logger, file_path string) *mux.Router {
	log.Info("MHF-API:launcher:middlewares:check:GetCheckRouter")
	controller := controllers.NewControllerCheck(log, file_path)

	for _, route := range GetCheckRoutes() {
		handler := common.CreateDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
