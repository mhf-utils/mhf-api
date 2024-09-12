package middlewares

import (
	"github.com/gorilla/mux"

	"mhf-api/server/controllers"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

var quest_routes = []Route{
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

func GetRouterQuest(log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:middlewares:quest:GetRouterQuest")
	controller := controllers.NewControllerQuest(log, binary_file)
	router := mux.NewRouter()

	for _, route := range quest_routes {
		handler := createDynamicHandler(controller, route.Handler)
		if handler != nil {
			router.HandleFunc(route.Endpoint, handler).Methods(route.Method)
		}
	}

	return router
}
