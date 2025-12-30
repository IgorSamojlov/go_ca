package testhelper

import (
	"database/sql"
	"fmt"
	"html/template"
	"math/rand"
	"strings"

	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
)

func AddDir(db *sql.DB, dir string) (*testfixtures.Loader, error) {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(dir),
		testfixtures.UseAlterConstraint(),
	)
	if err != nil {
		return nil, err
	}

	return fixtures, nil
}

func AddFiles(db *sql.DB, files []string) (*testfixtures.Loader, error) {
	fixtures, err := testfixtures.New(
		testfixtures.Template(),
		testfixtures.TemplateFuncs(funcMap),
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Files(files...),
		testfixtures.UseAlterConstraint(),
	)
	if err != nil {
		return nil, err
	}

	return fixtures, nil
}

func AddPaths(db *sql.DB, paths []string) (*testfixtures.Loader, error) {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Paths(paths...),
		testfixtures.UseAlterConstraint(),
	)
	if err != nil {
		return nil, err
	}

	return fixtures, nil
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var funcMap = template.FuncMap{
	"randomText": func() string {
		b := make([]rune, 10)

		for i := range b {
			b[i] = characterRunes[rand.Intn(len(characterRunes))]
		}
		return string(b)
	},
}

func AddFilesMultiTables(db *sql.DB, files []string) (*testfixtures.Loader, error) {
	fixtures, err := testfixtures.New(
		testfixtures.Template(),
		testfixtures.TemplateFuncs(funcMap),
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.FilesMultiTables(files...),
		testfixtures.UseAlterConstraint(),
	)
	if err != nil {
		return nil, err
	}

	return fixtures, nil
}

func TrancuteTables(db *sql.DB, tables []string) error {
	query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", strings.Join(tables, ", "))
	_, err := db.Query(query)
	if err != nil {
		return err
	}

	return nil
}
