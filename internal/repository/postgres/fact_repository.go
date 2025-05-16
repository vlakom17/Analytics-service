package postgres

import (
	"database/sql"

	"github.com/vlakom17/analytics-service/internal/domain/fact"
)

type FactRepo struct {
	DB *sql.DB
}

func NewFactRepo(db *sql.DB) *FactRepo {
	return &FactRepo{DB: db}
}

func (r *FactRepo) Insert(f *fact.ListenFact) error {
	_, err := r.DB.Exec(`
		INSERT INTO fact_listens (user_id, song_id, album_id, artist_id, genre_id, listened_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, f.UserID, f.SongID, f.AlbumID, f.ArtistID, f.GenreID, f.ListenedAt)
	return err
}
