package postgres

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"hw8/internal/store"
)

type DB struct {
	conn *sqlx.DB
	anime store.TitleRepository
	manga store.TitleRepository
	ranobe store.TitleRepository
}

func NewDB() store.Store {
	return &DB{}
}

func (db *DB) Connect(url string) error {
	conn, err := sqlx.Connect("pgx", url)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	db.conn = conn
	return nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) Anime() store.TitleRepository {
	if db.anime == nil{
		db.anime = NewAnimeRepository(db.conn)
	}
	return db.anime
}
func (db *DB) Manga() store.TitleRepository {
	if db.manga == nil{
		db.manga = NewMangaRepository(db.conn)
	}
	return db.manga
}
func (db *DB) Ranobe() store.TitleRepository {
	if db.ranobe == nil{
		db.ranobe = NewRanobeRepository(db.conn)
	}
	return db.ranobe
}