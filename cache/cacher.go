package cache

import "time"

type Cacher interface {
	Get([]byte) ([]byte, error)
	Set([]byte, []byte, time.Duration) error
	Delete([]byte) error
	Update([]byte, []byte) error
}
