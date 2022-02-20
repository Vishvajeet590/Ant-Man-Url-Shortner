package key

type Reader interface {
	Get(instance string) (int, int, error)
}
type Writer interface {
	Add(start, last int) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetRangeConfig(instance string) (int, int, error)
	AddKeyOut(start, last int) error
}
