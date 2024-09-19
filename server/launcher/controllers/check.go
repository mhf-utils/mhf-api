package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"mhf-api/utils/logger"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FileData struct {
	ChecksumHeader string            `json:"checksum_header"`
	Checksums      map[string]string `json:"checksums"`
}

type ControllerCheck struct {
	log       *logger.Logger
	game_data FileData
	file_path string
}

func NewControllerCheck(log *logger.Logger, file_path string) *ControllerCheck {
	controller := &ControllerCheck{
		log:       log,
		file_path: file_path,
		game_data: FileData{
			Checksums: make(map[string]string),
		},
	}

	if err := controller.LoadFolderData(); err != nil {
		log.Fatalf("Failed to load folder data: %v", err)
	}

	return controller
}

func (controller *ControllerCheck) LoadFolderData() error {
	dir_hasher := sha256.New()

	err := filepath.WalkDir(controller.file_path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			if d == nil {
				return fmt.Errorf("invalid root directory")
			}
			return err
		}

		if d.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		hasher := sha256.New()
		if _, err := io.Copy(hasher, file); err != nil {
			return err
		}
		checksum := hex.EncodeToString(hasher.Sum(nil))

		relative_path := strings.ReplaceAll(strings.TrimPrefix(path, controller.file_path), "\\", "/")
		controller.game_data.Checksums[relative_path] = checksum

		dir_hasher.Write([]byte(fmt.Sprintf("%s\t%s\n", checksum, relative_path)))

		return nil
	})

	if err != nil {
		return err
	}

	controller.game_data.ChecksumHeader = fmt.Sprintf("\"%s\"", hex.EncodeToString(dir_hasher.Sum(nil)))
	return nil
}

func (controller *ControllerCheck) CheckFiles(res http.ResponseWriter, req *http.Request) {
	etag := req.Header.Get("If-None-Match")
	if etag == controller.game_data.ChecksumHeader {
		res.WriteHeader(http.StatusNotModified)
		return
	}

	res.Header().Add("ETag", controller.game_data.ChecksumHeader)
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(controller.game_data.Checksums)
}
