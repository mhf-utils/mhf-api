package core

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Entity interface{}

func List[T Entity](
	from,
	to,
	length int,
	extractFunc func(int) T,
) []T {
	count := (to-from)/length + 1
	entries := make([]T, count)
	log.Printf("[COUNT]: %d\n", count)

	for index := 0; index < int(count); index++ {
		entries[index] = extractFunc(index)
	}

	return entries
}

func Read[T Entity](
	index int,
	extractFunc func(int) T,
) T {
	return extractFunc(index)
}

func Paginate[T any](
	res http.ResponseWriter,
	req *http.Request,
	from int64,
	to int64,
	length int64,
	extractFunc func(index int) T,
) {
	query := req.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit == 0 {
		limit = 10
	}
	page, _ := strconv.Atoi(query.Get("page"))
	if page == 0 {
		page = 1
	}
	sort := query.Get("sort")
	if sort == "" {
		sort = "id"
	}

	total_items := (to - from) / length
	offset := int64((page - 1) * limit)
	if offset > total_items {
		offset = total_items
	}

	total_pages := (total_items + int64(limit) - 1) / int64(limit)

	items := []T{}
	for i := 0; i < limit && offset+int64(i) < total_items; i++ {
		item := extractFunc(int(offset) + i)
		items = append(items, item)
	}

	response := map[string]interface{}{
		"results":      items,
		"count":        total_items,
		"current_page": page,
		"pages":        total_pages,
		"limit":        limit,
		"sort":         sort,
	}

	json.NewEncoder(res).Encode(response)
}

func ReadItem[T any](
	res http.ResponseWriter,
	req *http.Request,
	extractFunc func(index int) T,
) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	item := extractFunc(id)

	response := map[string]interface{}{
		"results": item,
	}

	json.NewEncoder(res).Encode(response)
}
