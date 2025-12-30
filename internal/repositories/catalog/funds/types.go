//         Column        |              Type              | Collation | Nullable |
// ----------------------+--------------------------------+-----------+----------+
//  id                   | bigint                         |           | not null |
//  code                 | character varying(20)          |           | not null |
//  code_order           | character varying              | C         | not null |
//  title                | character varying(1023)        |           |          |
//  number_of_arch_files | integer                        |           |          |
//  start_year           | integer                        |           |          |
//  end_year             | integer                        |           |          |
//  description          | text                           |           |          |
//  archive_id           | bigint                         |           | not null |
//  user_id              | bigint                         |           |          |
//  fund_type            | integer                        |           | not null |
//  dsp                  | boolean                        |           | not null |
//  created_at           | timestamp(6) without time zone |           | not null |
//  updated_at           | timestamp(6) without time zone |           | not null |

package funds

import (
	"time"

	"ca/internal/types/catalog"
)

type Fund struct {
	ID                int64     `db:"id"`
	Code              string    `db:"code"`
	CodeOrder         string    `db:"code_order"`
	Title             *string   `db:"title"`
	NumberOfArchFiles *int      `db:"number_of_arch_files"`
	StartYear         *int      `db:"start_year"`
	EndYear           *int      `db:"end_year"`
	ArchiveID         int64     `db:"archive_id"`
	UserID            *int64    `db:"user_id"`
	FundType          int       `db:"fund_type"`
	DSP               bool      `db:"dsp"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

func FillCollectionFromDo(dos []Fund) []catalog.Fund {
	funds := make([]catalog.Fund, 0, len(dos))

	for _, f := range dos {
		funds = append(funds, FillFromModel(f))
	}

	return funds
}

func FillFromModel(f Fund) catalog.Fund {
	fund := catalog.Fund{
		ID:        catalog.FundID(f.ID),
		Code:      f.Code,
		CodeOrder: f.CodeOrder,
		FundType:  f.FundType,
		DSP:       f.DSP,
		ArchiveID: f.ArchiveID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}

	if f.Title != nil {
		fund.Title = *f.Title
	}

	if f.StartYear != nil {
		fund.StartYear = *f.StartYear
	}

	if f.NumberOfArchFiles != nil {
		fund.NumberOfArchFiles = *f.NumberOfArchFiles
	}

	if f.EndYear != nil {
		fund.EndYear = *f.EndYear
	}

	if f.UserID != nil {
		fund.EndYear = *f.EndYear
	}

	return fund
}
