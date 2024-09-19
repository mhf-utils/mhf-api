package launcher

import (
	"github.com/gorilla/mux"

	"mhf-api/server/common"
	"mhf-api/server/launcher/middlewares"
	"mhf-api/server/types"
	"mhf-api/utils/logger"
)

func GetRoutes() []types.Route {
	var routes []types.Route

	routes = append(routes, middlewares.GetCheckRoutes()...)
	routes = append(routes, middlewares.GetFileRoutes()...)

	return routes
}

func GetRouter(router *mux.Router, prefix string, log *logger.Logger, file_path string) *mux.Router {
	log.Info("MHF-API:launcher:middlewares:GetRouter")

	common.DisplayAvailableRoutes(GetRoutes(), prefix)

	middlewares.GetCheckRouter(router, log, file_path)
	middlewares.GetFileRouter(router, log, file_path)

	return router
}
