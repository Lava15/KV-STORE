package store

import (
	"os"
	"testing"
)

func TestStoreOperations(t *testing.T) {
	s := GetStore()
	// Ensure the store is empty initially
	_, exists := s.Get("test_key")
	if exists {
		t.Errorf("Expected key 'test_key' to not exist initially")
	}

	// Test Set and Get
	s.Set("test_key", "test_value")
	val, exists := s.Get("test_key")
	if !exists {
		t.Errorf("Expected key 'test_key' to exist")
	}
	if val != "test_value" {
		t.Errorf("Expected value 'test_value', got '%s'", val)
	}

	// Test Delete
	s.Delete("test_key")
	_, exists = s.Get("test_key")
	if exists {
		t.Errorf("Expected key 'test_key' to be deleted")
	}
}

func TestStorePersistence(t *testing.T) {
	s := GetStore()

	s.Set("foo", "bar")

	// Save to file
	err := s.SaveFile("testdata.json")
	if err != nil {
		t.Fatalf("Error saving to file: %v", err)
	}
	defer os.Remove("testdata.json") // Clean up

	// Create a new store instance
	s2 := GetStore()

	s2.Delete("foo")
	err = s2.LoadFile("testdata.json")
	if err != nil {
		t.Fatalf("Error loading from file: %v", err)
	}

	val, exists := s2.Get("foo")
	if !exists {
		t.Errorf("Expected key 'foo' to exist after loading")
	}
	if val != "bar" {
		t.Errorf("Expected value 'bar', got '%s'", val)
	}
}
