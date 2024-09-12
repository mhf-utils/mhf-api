package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/gorilla/mux"

	"mhf-api/utils/binary"
	"mhf-api/utils/logger"
)

type Route struct {
	Endpoint string
	Handler  string
	Method   string
}

func GetRouter(router *mux.Router, log *logger.Logger, binary_file *binary.BinaryFile) *mux.Router {
	log.Info("MHF-API:middlewares:router:GetRouter")

	GetRouterEquipment(router, log, binary_file)
	GetRouterItem(router, log, binary_file)
	GetRouterQuest(router, log, binary_file)
	GetRouterWeaponMelee(router, log, binary_file)
	GetRouterWeaponRanged(router, log, binary_file)

	return router
}

func RouterKeeper(log *logger.Logger, locales []string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			clientIP := req.RemoteAddr
			endpoint := req.URL.Path
			method := req.Method

			if !isValidEndpoint(locales, endpoint, method) {
				errorMessage := fmt.Sprintf("Sorry, the endpoint '%s' you are trying to access is not available at this time. Please check the URL and try again later.", endpoint)
				http.Error(res, errorMessage, http.StatusInternalServerError)
				log.Error(fmt.Sprintf("Invalid endpoint - Endpoint: %s | Method: %s | ClientIP: %s", endpoint, method, clientIP))
				return
			}

			next.ServeHTTP(res, req)
		})
	}
}

func isValidEndpoint(locales []string, endpoint string, method string) bool {
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
		route_pattern = strings.ReplaceAll(route_pattern, "{type}", ".+")

		for _, locale := range locales {
			full_pattern := "^" + locale + route_pattern + "$"
			matched, err := regexp.MatchString(full_pattern, endpoint)

			if err != nil {
				log.Printf("Error compiling regex for endpoint %s: %s", route.Endpoint, err)
				continue
			}

			if matched {
				return true
			}
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
