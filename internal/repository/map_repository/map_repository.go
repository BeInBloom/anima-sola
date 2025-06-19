package maprepository

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("key not found")
)

type Repo struct {
	mu      sync.RWMutex
	storage map[string]string
}

func New() *Repo {
	const allocMem = 100

	return &Repo{
		storage: make(map[string]string, allocMem),
	}
}

func (r *Repo) Set(key, value string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.storage[key] = value
	return nil
}

func (r *Repo) Get(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	val, ok := r.storage[key]
	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}
