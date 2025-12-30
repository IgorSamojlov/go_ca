package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"

	routes "ca/cmd/http/routes"
	fundsHandler "ca/internal/handler/catalog/funds"
	fundsRepository "ca/internal/repositories/catalog/funds"
)

func Run(r *chi.Mux, db *pgx.Conn) {
	// repos
	rFunds := fundsRepository.New(db)

	// handlers
	hFunds := fundsHandler.New(rFunds)

	// routes
	caRoutes := routes.CARoutes{ChiRouter: r, HFunds: hFunds}
	caRoutes.Mount()

	http.ListenAndServe(":3000", r)
}
