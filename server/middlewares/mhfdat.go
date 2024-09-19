package middlewares

import (
	"mhf-api/config"
	"mhf-api/server/mhfdat"
	"mhf-api/utils/binary"
	"mhf-api/utils/logger"

	"github.com/gorilla/mux"
)

var binary_files []*binary.BinaryFile

func GetMhfdatRouter(log *logger.Logger, router *mux.Router, prefixes []string) (*mux.Router, []string) {
	mhfdat_configs := map[string]config.Info{
		"/en": config.GlobalConfig.Mhfdat.En,
		"/fr": config.GlobalConfig.Mhfdat.Fr,
		"/jp": config.GlobalConfig.Mhfdat.Jp,
	}

	for locale, mhfdat_config := range mhfdat_configs {
		if mhfdat_config.Enable {
			binary_file := binary.GetBinaryFile(mhfdat_config.FilePath)
			mhfdat_prefix := locale + "/mhfdat"
			mhfdat_sub_router := router.PathPrefix(mhfdat_prefix).Subrouter()
			mhfdat.GetRouter(mhfdat_sub_router, mhfdat_prefix, log, binary_file)
			prefixes = append(prefixes, mhfdat_prefix)
			binary_files = append(binary_files, binary_file)
		}
	}
	return router, prefixes
}

func CloseMhfdatBinaries() {
	if len(binary_files) == 0 {
		return
	}
	for _, binary_file := range binary_files {
		binary_file.Close()
	}
}
