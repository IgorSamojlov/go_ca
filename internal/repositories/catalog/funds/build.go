package funds

import (
	"ca/pkg/repository/helpers"

	sq "github.com/Masterminds/squirrel"
)

var fundsFields = []string{
	"id",
	"code",
	"code_order",
	"title",
	"number_of_arch_files",
	"start_year",
	"end_year",
	"archive_id",
	"user_id",
	"fund_type",
	"dsp",
	"created_at",
	"updated_at",
}

func buildQueryGetCollection(page, perPage uint64) (string, []any, error) {
	return helpers.PsqlBuilder().
		Select(fundsFields...).
		From("funds").
		OrderBy("code_order").
		Limit(uint64(perPage)).
		Offset(uint64((page - 1) * perPage)).
		ToSql()
}

func buildQueryGetByID(id uint64) (string, []any, error) {
	return helpers.PsqlBuilder().
		Select(fundsFields...).
		From("funds").
		Where(sq.Eq{"id": id}).
		ToSql()
}
