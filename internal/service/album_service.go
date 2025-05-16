package service

import "github.com/vlakom17/analytics-service/internal/domain/album"

type AlbumService struct {
	Repo album.AlbumRepository
}

func NewAlbumService(repo album.AlbumRepository) *AlbumService {
	return &AlbumService{Repo: repo}
}

func (s *AlbumService) CreateAlbum(album *album.Album) error {
	return s.Repo.Create(album)
}
