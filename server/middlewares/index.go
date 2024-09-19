package middlewares

import (
	"fmt"
	"log"
	"mhf-api/server/launcher"
	"mhf-api/server/mhfdat"
	"mhf-api/server/types"
	"mhf-api/utils/logger"
	"net/http"
	"regexp"
	"strings"
)

type Middleware func(http.Handler) http.Handler

func RouterKeeper(log *logger.Logger, prefixes []string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			clientIP := req.RemoteAddr
			endpoint := req.URL.Path
			method := req.Method

			if !isValidEndpoint(prefixes, endpoint, method) {
				errorMessage := fmt.Sprintf("Sorry, the endpoint '%s' you are trying to access is not available at this time. Please check the URL and try again later.", endpoint)
				http.Error(res, errorMessage, http.StatusNotImplemented)
				log.Error(fmt.Sprintf("Invalid endpoint - Endpoint: %s | Method: %s | ClientIP: %s", endpoint, method, clientIP))
				return
			}

			next.ServeHTTP(res, req)
		})
	}
}

func GetRoutes() []types.Route {
	var routes []types.Route

	routes = append(routes, launcher.GetRoutes()...)
	routes = append(routes, mhfdat.GetRoutes()...)

	return routes
}

func isValidEndpoint(prefixes []string, endpoint string, method string) bool {
	for _, prefix := range prefixes {
		for _, route := range GetRoutes() {
			if route.Method != method {
				continue
			}

			route_pattern := route.Endpoint
			route_pattern = strings.ReplaceAll(route_pattern, "{id}", "[0-9]+")
			route_pattern = strings.ReplaceAll(route_pattern, "{type}", ".+")
			route_pattern = strings.ReplaceAll(route_pattern, "{filename}", "[^/]*")

			full_pattern := "^" + prefix + route_pattern + "$"
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

func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
