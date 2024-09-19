package common

import (
	"fmt"
	"mhf-api/server/types"
	"net/http"
	"reflect"
)

func CreateDynamicHandler(controller interface{}, handlerName string) http.HandlerFunc {
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

func DisplayAvailableRoutes(routes []types.Route, prefix string) {
	fmt.Printf("Routes initialization with prefix='%s'\n", prefix)

	for _, route := range routes {
		fmt.Printf("[%s] -> '%s' -> %s()\n", route.Method, prefix+route.Endpoint, route.Handler)
	}

	fmt.Printf("Routes with prefix='%s' initialized successfully!\n", prefix)
}
