package url

import "Ant-Man-Url/entity"

type Reader interface {
	Resolve(shortUrl *entity.Url) (*entity.Url, error)
}

type Writer interface {
	Link(LongUrl *entity.Url, isJwt bool, userId *int) (*entity.Url, error)
}

type Deleter interface {
	Delete(ShortUrl *entity.Url, isJwt bool, userId *int) (bool, error)
}

type Repository interface {
	Reader
	Writer
	Deleter
}

type UseCase interface {
	ResolveUrl(ShortUrl *entity.Url) (*entity.Url, error)
	MapUrl(LongUrl *entity.Url, isJwt bool, userId *int) (*entity.Url, error)
	DeleteUrl(ShortUrl *entity.Url, isJwt bool, userId *int) (bool, error)
}
