package user

import "Ant-Man-Url/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) SignUpUser(userName, userPassword, userEmail, userRole string) (*entity.User, error) {
	user, err := s.repo.SignUp(userName, userPassword, userEmail, userRole)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) LoginUser(email, password string) (*entity.User, error) {
	user, err := s.repo.Login(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetUrlStat(user_id int, keyval string) (*entity.Url, error) {
	urlStat, err := s.repo.Get(user_id, keyval)
	if err != nil {
		return nil, err
	}
	return urlStat, nil
}

func (s *Service) GetUrlList(user_id int) ([]*entity.Url, error) {
	urlList, err := s.repo.List(user_id)
	if err != nil {
		return nil, err
	}
	return urlList, nil
}
