package funds

import (
	"context"

	"ca/internal/types/catalog"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll(ctx context.Context, page, perPage uint64) ([]catalog.Fund, error) {
	query, args, err := buildQueryGetCollection(page, perPage)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	funds, err := pgx.CollectRows(rows, pgx.RowToStructByName[Fund])
	if err != nil {
		return nil, err
	}

	return FillCollectionFromDo(funds), nil
}

func (r *Repository) GetByID(ctx context.Context, id uint64) (catalog.Fund, error) {
	query, args, err := buildQueryGetByID(id)
	if err != nil {
		return catalog.Fund{}, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return catalog.Fund{}, err
	}

	fund, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Fund])
	if err != nil {
		return catalog.Fund{}, err
	}

	return FillFromModel(fund), nil
}
