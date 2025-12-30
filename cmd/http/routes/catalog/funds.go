package catalog

import (
	fundsHandler "ca/internal/handler/catalog/funds"

	"github.com/go-chi/chi/v5"
)

func MountFunds(r *chi.Mux, h *fundsHandler.Handler) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/funds", h.GetCollection)
		r.Get("/funds/{id}", h.GetByID)
	})
}
