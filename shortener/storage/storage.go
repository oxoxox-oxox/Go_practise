package storage

type Storage interface {
	Save(shortCode, longURL string) error
	Load(shortCode string) (string, bool)
}