package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"hw8/internal/models"
	"hw8/internal/store"
)

type MangaRepository struct {
	conn *sqlx.DB
}

func NewMangaRepository(conn *sqlx.DB) store.TitleRepository {
	return &MangaRepository{
		conn:      conn,
	}
}

func (t* MangaRepository) Create(ctx context.Context, category *models.Title) error {
	_, err := t.conn.Exec(`insert into manga (name, name_english, release, final, status, type) values ($1, $2, $3, $4, $5, $6)`,
		category.Name, category.NameEnglish, category.Release, category.Final, category.Status, category.Type)
	if err != nil {
		return err
	}
	return nil
}

func (t* MangaRepository) All(ctx context.Context, filter *models.TitleFilter) ([]*models.Title, error) {
	titles := make([]*models.Title, 0)
	basicQuery := "SELECT * FROM manga"
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

func (t* MangaRepository) ByID(ctx context.Context, id int) (*models.Title, error) {
	title := new(models.Title)
	if err := t.conn.Get(title, "SELECT * FROM manga WHERE id = $1", id); err != nil {
		return nil, err
	}

	return title, nil
}

func (t* MangaRepository) Update(ctx context.Context, title *models.Title) error {
	_, err := t.conn.Exec("UPDATE manga SET name = $1, name_english = $2, release = $3, final = $4, status = $5, type = $7 WHERE id = $8", title.Name, title.NameEnglish, title.Release, title.Final, title.Status, title.Type, title.Id)
	if err != nil {
		return err
	}

	return nil
}

func (t* MangaRepository) Delete(ctx context.Context, id int) error {
	_, err := t.conn.Exec("DELETE FROM manga WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
