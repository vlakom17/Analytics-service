package service

import "github.com/vlakom17/analytics-service/internal/domain/song"

type SongService struct {
	Repo song.SongRepository
}

func NewSongService(repo song.SongRepository) *SongService {
	return &SongService{Repo: repo}
}

func (s *SongService) CreateSong(song *song.Song) error {
	return s.Repo.Create(song)
}

func (s *SongService) GetPopularSongs(limit int) ([]song.PopularSong, error) {
	return s.Repo.GetPopularSongs(limit)
}
