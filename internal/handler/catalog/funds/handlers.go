package funds

import (
	"context"
	"fmt"
	"net/http"

	"ca/internal/types/catalog"
	"ca/pkg/handlers/helpers"

	"github.com/go-chi/render"
)

const perPage = uint64(25)

type fundsRepository interface {
	GetAll(ctx context.Context, page uint64, perPage uint64) ([]catalog.Fund, error)
	GetByID(ctx context.Context, id uint64) (catalog.Fund, error)
	// Delete(ctx context.Context, id uint64) (catalog.Fund, error)
}

type Handler struct {
	fundsRepository fundsRepository
}

func New(r fundsRepository) *Handler {
	return &Handler{
		fundsRepository: r,
	}
}

func (h *Handler) GetCollection(w http.ResponseWriter, r *http.Request) {
	paginator := helpers.GetPaginator(r)

	funds, err := h.fundsRepository.GetAll(r.Context(), paginator.Page, perPage)
	if err != nil {
		fmt.Println(err)
	}
	render.Render(w, r, &HttpResponse{Collection: funds})
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := helpers.GetIDFromParams(r)

	fund, err := h.fundsRepository.GetByID(r.Context(), id)
	if err != nil {
		fmt.Println(err)
	}
	render.Render(w, r, &HttpResponse{Record: fund})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := helpers.GetIDFromParams(r)

	fund, err := h.fundsRepository.GetByID(r.Context(), id)
	if err != nil {
		fmt.Println(err)
	}
	render.Render(w, r, &HttpResponse{Record: fund})
}
