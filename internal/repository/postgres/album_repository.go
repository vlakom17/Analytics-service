package postgres

import (
	"database/sql"

	"github.com/vlakom17/analytics-service/internal/domain/album"
)

type AlbumRepo struct {
	DB *sql.DB
}

func NewAlbumRepo(db *sql.DB) *AlbumRepo {
	return &AlbumRepo{DB: db}
}

func (r *AlbumRepo) Create(album *album.Album) error {
	_, err := r.DB.Exec("INSERT INTO albums (name) VALUES ($1)", album.Name)
	return err
}
