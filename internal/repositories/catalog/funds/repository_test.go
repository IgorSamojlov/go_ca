package funds

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"ca/internal/types/catalog"
	"ca/pkg/testhelper"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func seeds(db *sql.DB) error {
	fmt.Println("=========Seeds==========")

	files := []string{
		"../../../../testing/fixtures/shared/archives.yml",
		"../../../../testing/fixtures/shared/funds.yml",
	}

	loader, err := testhelper.AddFiles(db, files)
	if err != nil {
		return err
	}

	err = loader.Load()
	if err != nil {
		return err
	}
	return nil
}

func clearTables(db *sql.DB) error {
	fmt.Println("=========ClearDb==========")
	err := testhelper.TrancuteTables(db, []string{"archives", "funds"})
	if err != nil {
		return err
	}
	return nil
}

func TestRepositoryGetAll(t *testing.T) {
	ctx := context.Background()
	cfg, err := testhelper.NewFromFile("../../../../.test_db.yaml")
	if err != nil {
		assert.NoError(t, err)
	}

	db, err := sql.Open(cfg.Database.Dialect, cfg.Database.ConString)
	if err != nil {
		assert.NoError(t, err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			assert.NoError(t, err)
		}
	}()

	rdb, err := pgx.Connect(ctx, cfg.Database.ConString)
	if err != nil {
		assert.NoError(t, err)
	}

	defer rdb.Close(ctx)

	err = seeds(db)
	if err != nil {
		assert.NoError(t, err)
	}

	type args struct {
		page    uint64
		perPage uint64
	}
	tests := []struct {
		setup   func()
		name    string
		args    args
		want    []catalog.Fund
		wantErr error
	}{
		{
			name: "First page collection",
			args: args{page: 1, perPage: 2},
			want: []catalog.Fund{
				{
					ID:                catalog.FundID(2),
					Code:              "4",
					CodeOrder:         "      000004      ",
					Title:             "5-е городское начальное женское училище",
					NumberOfArchFiles: 24,
					StartYear:         1885,
					EndYear:           1916,
					ArchiveID:         1,
					FundType:          1,
					DSP:               false,
					CreatedAt:         time.Date(2025, time.August, 29, 21, 18, 45, 0, time.UTC),
					UpdatedAt:         time.Date(2025, time.August, 29, 21, 18, 45, 0, time.UTC),
				},
				{
					ID:                catalog.FundID(3),
					Code:              "5",
					CodeOrder:         "      000005      ",
					Title:             "Департамент шоссейных и грунтовых дорог Всевеликого войска Донского",
					NumberOfArchFiles: 83,
					StartYear:         1918,
					EndYear:           1921,
					ArchiveID:         1,
					FundType:          1,
					DSP:               false,
					CreatedAt:         time.Date(2025, time.August, 29, 21, 18, 45, 0, time.UTC),
					UpdatedAt:         time.Date(2025, time.August, 29, 21, 18, 45, 0, time.UTC),
				},
			},
		},
		{
			name: "When third page collection",
			args: args{page: 3, perPage: 1},
			want: []catalog.Fund{
				{
					ID:                catalog.FundID(4),
					Code:              "8",
					CodeOrder:         "      000008      ",
					Title:             "Пробст 2-го Южно-Российского евангелическо-лютеранского округа, г. Ростов-на-Дону",
					NumberOfArchFiles: 123,
					StartYear:         1835,
					EndYear:           1917,
					ArchiveID:         1,
					FundType:          1,
					DSP:               false,
					CreatedAt:         time.Date(2025, time.August, 29, 21, 18, 45, 0, time.UTC),
					UpdatedAt:         time.Date(2025, time.August, 29, 21, 18, 45, 0, time.UTC),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: rdb,
			}
			got, err := r.GetAll(ctx, tt.args.page, tt.args.perPage)

			if tt.wantErr != nil {
				assert.Equal(t, err.Error(), tt.wantErr.Error())
			} else {
				assert.ElementsMatch(t, got, tt.want)
			}
		})
	}

	err = clearTables(db)
	if err != nil {
		assert.NoError(t, err)
	}
}
