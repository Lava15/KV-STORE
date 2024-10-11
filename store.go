package main

import "sync"

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
