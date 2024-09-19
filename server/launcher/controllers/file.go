package controllers

import (
	"fmt"
	"mhf-api/server/launcher/views"
	"mhf-api/utils/logger"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type ControllerFile struct {
	log       *logger.Logger
	file_path string
}

func NewControllerFile(log *logger.Logger, file_path string) *ControllerFile {
	return &ControllerFile{
		log,
		file_path,
	}
}

type FileInfo struct {
	Name string
	Link string
	Type string
}

func (controller *ControllerFile) List(res http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	locale := ""
	if len(parts) > 1 {
		locale = parts[1]
	}

	relative_path := strings.TrimPrefix(req.URL.Path, fmt.Sprintf("/%s/launcher/files", locale))
	if relative_path == "" {
		relative_path = "/"
	}
	abs_path := filepath.Join(controller.file_path, relative_path)

	file_info, err := os.Stat(abs_path)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(res, "File not found", http.StatusNotFound)
		} else {
			http.Error(res, "Error accessing file", http.StatusInternalServerError)
		}
		return
	}

	if file_info.IsDir() {
		controller.serveDirectory(res, req, abs_path, locale, relative_path)
	} else {
		controller.serveFile(res, req, abs_path, file_info.Name())
	}
}

func (controller *ControllerFile) serveDirectory(res http.ResponseWriter, req *http.Request, abs_path, locale, relative_path string) {
	dir, err := os.Open(abs_path)
	if err != nil {
		http.Error(res, "Failed to open directory", http.StatusInternalServerError)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		http.Error(res, "Failed to read directory", http.StatusInternalServerError)
		return
	}

	var folders, files_list []FileInfo
	for _, file := range files {
		name := file.Name()
		link := filepath.Join(req.URL.Path, name)
		if file.IsDir() {
			folders = append(folders, FileInfo{Name: name, Link: link, Type: "folder"})
		} else {
			files_list = append(files_list, FileInfo{Name: name, Link: link, Type: "file"})
		}
	}

	sort.Slice(folders, func(i, j int) bool {
		return folders[i].Name < folders[j].Name
	})
	sort.Slice(files_list, func(i, j int) bool {
		return files_list[i].Name < files_list[j].Name
	})

	file_list := append(folders, files_list...)

	parent_link := fmt.Sprintf("/%s/launcher/files", locale)
	if relative_path != "/" {
		parent_link = filepath.Dir(req.URL.Path)
		if parent_link == "/" {
			parent_link = fmt.Sprintf("/%s/launcher/files", locale)
		}
	}

	data := struct {
		Locale     string
		Path       string
		Files      []FileInfo
		ParentLink string
	}{
		Locale:     locale,
		Path:       req.URL.Path,
		Files:      file_list,
		ParentLink: parent_link,
	}

	res.Header().Set("Content-Type", "text/html")

	tmpl, err := template.New("dirList").Parse(views.FileSystemTemplate)
	if err != nil {
		http.Error(res, "Failed to generate directory listing", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(res, data); err != nil {
		http.Error(res, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func (controller *ControllerFile) serveFile(res http.ResponseWriter, req *http.Request, abs_path, filename string) {
	res.Header().Set("Content-Disposition", "attachment; filename="+filename)
	res.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(res, req, abs_path)
}
