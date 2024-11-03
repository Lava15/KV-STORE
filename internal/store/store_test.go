package store

import "testing"

func TestStore(t *testing.T) {
	s := NewStore()
	s.Set("test_key", "test_value")
	val, ok := s.Get("test_key")
	if !ok || val != "test_value" {
		t.Errorf("Expected val=test_value, got=%v", val)
	}
	s.Delete("test_key")
	_, ok = s.Get("test_key")
	if ok {
		t.Errorf("Expected val=nil, got=%v", val)
	}
}

func TestPersistence(t *testing.T) {
	s := NewStore()
	s.Set("foo", "bar")
	err := s.SaveFile("testdata.json")
	if err != nil {
		t.Errorf("Error saving to file: %v", err)
	}

	s2 := NewStore()
	err = s2.LoadFile("testdata.json")
	if err != nil {
		t.Errorf("Error loading from file: %v", err)
	}

	val, ok := s2.Get("foo")
	if !ok || val != "bar" {
		t.Errorf("Expected 'bar', got '%s'", val)
	}
}
