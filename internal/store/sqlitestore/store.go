package sqlitestore

import (
	"database/sql"
	"genxoft.dev/internal/store"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

type Store struct {
	db                 *sql.DB
	fcmTokenRepository store.FcmRepository
}

func New(f string) *Store {
	db, err := sql.Open("sqlite3", "file:"+f)
	if err != nil {
		panic(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	m, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Migrations done %d", m)

	return &Store{
		db: db,
	}
}

func (s *Store) GetDb() *sql.DB {
	return s.db
}

func (s *Store) FcmToken() store.FcmRepository {
	if s.fcmTokenRepository != nil {
		return s.fcmTokenRepository
	}
	s.fcmTokenRepository = &FcmRepository{
		db: s.db,
	}

	return s.fcmTokenRepository
}
