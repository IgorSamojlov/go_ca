package helpers

import (
	sq "github.com/Masterminds/squirrel"
)

func PsqlBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
