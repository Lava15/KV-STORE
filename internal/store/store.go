package store

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

type Store struct {
	sync.RWMutex
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key string, value string) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.data, key)
}

func (s *Store) SaveFile(filename string) error {
	s.RLock()
	defer s.RUnlock()
	data, err := json.Marshal(s.data)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (s *Store) LoadFile(filename string) error {
	s.RLock()
	defer s.Unlock()
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.data)
}
