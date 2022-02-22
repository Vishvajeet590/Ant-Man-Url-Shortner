package entity

import "time"

type Url struct {
	Long_link  string
	Keyval     string
	Redirects  int
	OwnerId    int
	Created_at time.Time
}

func NewUrl(Longlink, shortLink string, createdAt time.Time, redirects, ownerId int) *Url {
	return &Url{
		Long_link:  Longlink,
		Keyval:     shortLink,
		Redirects:  redirects,
		OwnerId:    ownerId,
		Created_at: createdAt,
	}
}
