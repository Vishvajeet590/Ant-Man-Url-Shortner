package user

import (
	"Ant-Man-Url/entity"
)

type Reader interface {
	Login(email, password string) (*entity.User, error)
	Get(user_id int, keyval string) (*entity.Url, error)
	List(user_id int) ([]*entity.Url, error)
}

type Writer interface {
	SignUp(userName, userPassword, userEmail, userRole string) (*entity.User, error)
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	SignUpUser(userName, userPassword, userEmail, userRole string) (*entity.User, error)
	LoginUser(email, password string) (*entity.User, error)
	GetUrlStat(user_id int, keyval string) (*entity.Url, error)
	GetUrlList(user_id int) ([]*entity.Url, error)
}
