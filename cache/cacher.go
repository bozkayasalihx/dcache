package cache

type Cacher interface {
	Get([]byte) ([]byte, error)
	Set([]byte, []byte) error
	Delete([]byte) ([]byte, error)
	Update([]byte, []byte) error
}
