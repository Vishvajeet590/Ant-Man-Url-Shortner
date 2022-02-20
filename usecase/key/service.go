package key

import (
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetRangeConfig(instance string) (int, int, error) {
	start, last, err := s.repo.Get(instance)
	if err != nil {
		return -999, -999, err
	}
	if start > last {
		return -999, -999, errors.New("Error : Range error.")
	}
	return start, last, nil
}

func (s *Service) AddKeyOut(start, last int) error {
	err := s.repo.Add(start, last)
	if err != nil {
		return err
	}
	return nil
}
