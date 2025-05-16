package service

import "github.com/vlakom17/analytics-service/internal/domain/fact"

type FactService struct {
	Repo fact.FactRepository
}

func NewFactService(repo fact.FactRepository) *FactService {
	return &FactService{Repo: repo}
}

func (s *FactService) Store(fact *fact.ListenFact) error {
	return s.Repo.Insert(fact)
}
