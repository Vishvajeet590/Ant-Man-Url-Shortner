package entity

type Key struct {
	Keyval     string
	url_linked string
	keyid      int
}

func NewKey(keyval string, url string, id int) *Key {
	key := &Key{
		Keyval:     keyval,
		url_linked: url,
		keyid:      id,
	}
	return key
}
