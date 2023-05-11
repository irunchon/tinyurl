package storage

type Storage interface {
	Get(string) (string, error)
	Set(string, string)
}
