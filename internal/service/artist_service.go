package service

import (
	"github.com/vlakom17/analytics-service/internal/domain/artist"
)

type ArtistService struct {
	Repo artist.ArtistRepository
}

func NewArtistService(repo artist.ArtistRepository) *ArtistService {
	return &ArtistService{Repo: repo}
}

func (s *ArtistService) CreateArtist(artist *artist.Artist) error {
	return s.Repo.Create(artist)
}
