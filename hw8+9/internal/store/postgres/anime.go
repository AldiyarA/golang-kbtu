package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"hw8/internal/models"
	"hw8/internal/store"
)

type AnimeRepository struct {
	conn *sqlx.DB
}

func NewAnimeRepository(conn *sqlx.DB) store.TitleRepository {
	return &AnimeRepository{
		conn:      conn,
	}
}

func (t* AnimeRepository) Create(ctx context.Context, category *models.Title) error {
	_, err := t.conn.Exec(`insert into anime (name, name_english, release, final, status, type) values ($1, $2, $3, $4, $5, $6)`,
		category.Name, category.NameEnglish, category.Release, category.Final, category.Status, category.Type)
	if err != nil {
		return err
	}
	return nil
}

func (t* AnimeRepository) All(ctx context.Context, filter *models.TitleFilter) ([]*models.Title, error) {
	titles := make([]*models.Title, 0)
	basicQuery := "SELECT * FROM anime"
	if filter.Query != nil {
		basicQuery = fmt.Sprintf("%s WHERE name_english ILIKE $1", basicQuery)

		if err := t.conn.Select(&titles, basicQuery, "%"+*filter.Query+"%"); err != nil {
			return nil, err
		}
		return titles, nil
	}
	if err := t.conn.Select(&titles, basicQuery); err != nil {
		return nil, err
	}

	return titles, nil
}

func (t* AnimeRepository) ByID(ctx context.Context, id int) (*models.Title, error) {
	title := new(models.Title)
	if err := t.conn.Get(title, "SELECT * FROM anime WHERE id = $1", id); err != nil {
		return nil, err
	}

	return title, nil
}

func (t* AnimeRepository) Update(ctx context.Context, title *models.Title) error {
	_, err := t.conn.Exec("UPDATE anime SET name = $1, name_english = $2, release = $3, final = $4, status = $5, type = $6 WHERE id = $7", title.Name, title.NameEnglish, title.Release, title.Final, title.Status, title.Type, title.Id)
	if err != nil {
		return err
	}

	return nil
}

func (t* AnimeRepository) Delete(ctx context.Context, id int) error {
	_, err := t.conn.Exec("DELETE FROM anime WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
