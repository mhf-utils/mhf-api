package middlewares

import (
	"mhf-api/config"
	"mhf-api/server/launcher"
	"mhf-api/utils/logger"

	"github.com/gorilla/mux"
)

func GetLauncherRouter(log *logger.Logger, router *mux.Router, prefixes []string) (*mux.Router, []string) {
	launcher_configs := map[string]config.Info{
		"/en": config.GlobalConfig.Launcher.En,
		"/fr": config.GlobalConfig.Launcher.Fr,
		"/jp": config.GlobalConfig.Launcher.Jp,
	}

	for locale, launcher_config := range launcher_configs {
		if launcher_config.Enable {
			launcher_prefix := locale + "/launcher"
			launcher_sub_router := router.PathPrefix(launcher_prefix).Subrouter()
			launcher.GetRouter(launcher_sub_router, launcher_prefix, log, launcher_config.FilePath)
			prefixes = append(prefixes, launcher_prefix)
		}
	}
	return router, prefixes
}
