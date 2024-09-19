package middlewares

import (
	"fmt"

	"github.com/gorilla/mux"

	"mhf-api/server/common"
	"mhf-api/server/launcher/controllers"
	"mhf-api/server/types"
	"mhf-api/utils/logger"
)

func GetFileRoutes() []types.Route {
	return []types.Route{
		{
			Endpoint: "/files/{type}/{filename}",
			Handler:  "List",
			Method:   "GET",
		},
		{
			Endpoint: "/files/{type}",
			Handler:  "List",
			Method:   "GET",
		},
		{
			Endpoint: "/files",
			Handler:  "List",
			Method:   "GET",
		},
	}
}

func GetFileRouter(router *mux.Router, log *logger.Logger, file_path string) *mux.Router {
	log.Info("MHF-API:launcher:middlewares:file:GetFileRouter")
	fmt.Println("File path:", file_path)
	controller := controllers.NewControllerFile(log, file_path)

	for _, route := range GetFileRoutes() {
		handler := common.CreateDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
