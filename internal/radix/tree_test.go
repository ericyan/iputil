package radix

import (
	"testing"
)

func testSet(tree *Tree, key []byte, val interface{}, expected error, t *testing.T) {
	if err := tree.Set(key, val); err != expected {
		t.Errorf("unexpected error: got '%v', want '%v'", err, expected)
	}
}

func testGet(tree *Tree, key []byte, val interface{}, expected error, t *testing.T) {
	got, err := tree.Get(key)
	if err != expected {
		t.Errorf("unexpected error: got '%v', want '%v'", err, expected)
	}
	if got != val {
		t.Errorf("unexpected value: got '%v', want '%v'", got, val)
	}
}

func TestTree(t *testing.T) {
	tree := NewTree()

	// Insert a new node
	testGet(tree, []byte("hello"), nil, ErrNotFound, t)
	testSet(tree, []byte("hello"), "world", nil, t)
	testGet(tree, []byte("hello"), "world", nil, t)
	testGet(tree, []byte("he"), nil, ErrNotFound, t)

	// Update existing node
	testSet(tree, []byte("hello"), "eric", nil, t)
	testGet(tree, []byte("hello"), "eric", nil, t)

	// Handle errors
	testSet(tree, nil, nil, ErrInvalidKey, t)
	testGet(tree, nil, nil, ErrInvalidKey, t)
}
