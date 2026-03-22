package storage

import (
	"sync"
)

type MemoryStorage struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		store: make(map[string]string),
	}
}

func (m *MemoryStorage) Save(shortCode, longURL string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[shortCode] = longURL
	return nil
}

func (m *MemoryStorage) Load(shortCode string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	longURL, ok := m.store[shortCode]
	return longURL, ok
}
