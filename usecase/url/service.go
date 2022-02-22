package url

import "Ant-Man-Url/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) ResolveUrl(shortUrl *entity.Url) (*entity.Url, error) {
	LongUrl, err := s.repo.Resolve(shortUrl)
	if err != nil {
		return nil, err
	}
	return LongUrl, nil
}

func (s *Service) MapUrl(longUrl *entity.Url, isJwt bool, userId *int) (*entity.Url, error) {
	shortUrl, err := s.repo.Link(longUrl, isJwt, userId)
	if err != nil {
		return nil, err
	}
	return shortUrl, nil
}

func (s *Service) DeleteUrl(ShortUrl *entity.Url, isJwt bool, userId *int) (bool, error) {
	res, err := s.repo.Delete(ShortUrl, isJwt, userId)
	if err != nil {
		return false, err
	}
	return res, nil
}
