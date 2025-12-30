package routes

import (
	"ca/cmd/http/routes/catalog"
	fundsHandler "ca/internal/handler/catalog/funds"

	"github.com/go-chi/chi/v5"
)

type CARoutes struct {
	HFunds    *fundsHandler.Handler
	ChiRouter *chi.Mux
}

func (r *CARoutes) Mount() {
	catalog.MountFunds(r.ChiRouter, r.HFunds)
}
