package store

type Cache interface {
	Get(code string) (string, bool)
	Set(code string, url string) error
}
