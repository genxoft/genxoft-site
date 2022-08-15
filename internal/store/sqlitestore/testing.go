package sqlitestore

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
	"testing"
)

func TestStore(t *testing.T) (*Store, func(...string)) {
	t.Helper()

	s := New("./../../../test/site.db")
	db := s.GetDb()
	if err := s.db.Ping(); err != nil {
		t.Fatal(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./../../../migrations",
	}

	if _, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			for _, tt := range tables {
				if _, err := db.Exec(fmt.Sprintf("DELETE FROM %s", tt)); err != nil {
					t.Fatal(err)
				}
			}

		}

		if err := db.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
