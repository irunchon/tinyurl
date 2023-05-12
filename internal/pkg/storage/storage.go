package storage

type Storage interface {
	GetShortURLbyLong(string) (string, error)
	GetLongURLbyShort(string) (string, error)
	SetShortAndLongURLs(string, string) error
}
