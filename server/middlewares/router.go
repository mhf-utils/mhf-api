package middlewares

import (
	"log"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Endpoint string
	Handler  string
	Method   string
}

func GetRouter(log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:middlewares:router:GetRouter")
	router := mux.NewRouter()

	router_equipment := GetRouterEquipment(log, binary_file)
	router_item := GetRouterItem(log, binary_file)
	router_quest := GetRouterQuest(log, binary_file)
	router_weapon_melee := GetRouterWeaponMelee(log, binary_file)
	router_weapon_ranged := GetRouterWeaponRanged(log, binary_file)

	router.PathPrefix("/equipments").Handler(router_equipment)
	router.PathPrefix("/items").Handler(router_item)
	router.PathPrefix("/quests").Handler(router_quest)
	router.PathPrefix("/weapons/melee").Handler(router_weapon_melee)
	router.PathPrefix("/weapons/ranged").Handler(router_weapon_ranged)

	return router
}

func IsValidEndpoint(endpoint string, method string) bool {
	var routes []Route
	routes = append(routes, equipment_routes...)
	routes = append(routes, item_routes...)
	routes = append(routes, quest_routes...)
	routes = append(routes, weapon_melee_routes...)
	routes = append(routes, weapon_ranged_routes...)

	for _, route := range routes {
		if route.Method != method {
			continue
		}

		route_pattern := route.Endpoint
		route_pattern = strings.ReplaceAll(route_pattern, "{id}", "[0-9]+")
		matched, err := regexp.MatchString("^"+route_pattern+"$", endpoint)

		if err != nil {
			log.Printf("Error compiling regex for endpoint %s: %s", route.Endpoint, err)
			continue
		}

		if matched {
			return true
		}
	}

	return false
}

func createDynamicHandler(controller interface{}, handlerName string) http.HandlerFunc {
	method := reflect.ValueOf(controller).MethodByName(handlerName)
	if !method.IsValid() {
		return nil
	}

	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		method.Call([]reflect.Value{
			reflect.ValueOf(res),
			reflect.ValueOf(req),
		})
	}
}
