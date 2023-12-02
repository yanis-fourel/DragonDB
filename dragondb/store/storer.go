package store

type Storer interface {
	Get(key string) string
	Set(key string, value string)
	Close() error
}
