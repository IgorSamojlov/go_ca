package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Paginator struct {
	Page uint64
}

func GetPaginator(r *http.Request) Paginator {
	page, _ := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)

	if page == 0 {
		page = 1
	}

	return Paginator{Page: page}
}

func GetIDFromParams(r *http.Request) uint64 {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	fmt.Println(chi.URLParam(r, "id"))

	return id
}
