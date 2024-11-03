package store

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

type Store interface {
	Get(key string) (string, bool)
	Set(key string, value string)
	Delete(key string)
	SaveFile(filename string) error
	LoadFile(filename string) error
}

type store struct {
	sync.RWMutex
	data map[string]string
}

var (
	instance Store
	once     sync.Once
)

func GetStore() Store {
	once.Do(func() {
		instance = &store{
			data: make(map[string]string),
		}
	})
	return instance
}

func (s store) Get(key string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok := s.data[key]
	return value, ok
}

func (s store) Set(key string, value string) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}

func (s store) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.data, key)
}

func (s store) SaveFile(filename string) error {
	s.RLock()
	defer s.RUnlock()
	data, err := json.Marshal(s.data)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (s store) LoadFile(filename string) error {
	s.Lock()
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
