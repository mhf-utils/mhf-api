package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/common"
	"mhf-api/server/mhfdat/controllers"
	"mhf-api/server/types"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

func GetQuestRoutes() []types.Route {
	return []types.Route{
		{
			Endpoint: "/quests/{type}",
			Handler:  "List",
			Method:   "GET",
		},
		{
			Endpoint: "/quests/{type}/{id}",
			Handler:  "Read",
			Method:   "GET",
		},
	}
}

func GetQuestRouter(router *mux.Router, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:mhfdat:middlewares:quest:GetRouterQuest")
	controller := controllers.NewControllerQuest(log, binary_file)

	for _, route := range GetQuestRoutes() {
		handler := common.CreateDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
