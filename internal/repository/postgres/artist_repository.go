package postgres

import (
	"database/sql"

	"github.com/vlakom17/analytics-service/internal/domain/artist"
)

type ArtistRepo struct {
	DB *sql.DB
}

func NewArtistRepo(db *sql.DB) *ArtistRepo {
	return &ArtistRepo{DB: db}
}

func (r *ArtistRepo) Create(artist *artist.Artist) error {
	_, err := r.DB.Exec("INSERT INTO artists (name) VALUES ($1)", artist.Name)
	return err
}
