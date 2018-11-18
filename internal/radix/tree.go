package radix

import (
	"errors"
)

var (
	// ErrNotFound is returned when the key does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidKey is returned if the key is invalid.
	ErrInvalidKey = errors.New("invalid key")
)

type edge struct {
	label byte
	node  *node
}

type node struct {
	value interface{}
	edges []edge
}

func (n *node) findChild(label byte) *node {
	for _, e := range n.edges {
		if e.label == label {
			return e.node
		}
	}

	return nil
}

func (n *node) newChild(label byte) *node {
	child := new(node)
	n.edges = append(n.edges, edge{label, child})
	return child
}

func (n *node) isLeaf() bool {
	return n.value != nil
}

// Tree represents a radix tree.
type Tree struct {
	root *node
}

// NewTree returns an empty Tree.
func NewTree() *Tree {
	return &Tree{
		root: new(node),
	}
}

// Get retrieves the value for a key.
func (t *Tree) Get(key []byte) (interface{}, error) {
	if len(key) == 0 {
		return nil, ErrInvalidKey
	}

	cur := t.root
	for i := 0; i < len(key); i++ {
		child := cur.findChild(key[i])
		if child == nil {
			return nil, ErrNotFound
		}

		cur = child
	}

	if cur.isLeaf() {
		return cur.value, nil
	}

	return nil, ErrNotFound
}

// Set sets the value for a key. If the key already exists, its previous
// value will be overwritten.
func (t *Tree) Set(key []byte, value interface{}) error {
	if len(key) == 0 {
		return ErrInvalidKey
	}

	cur := t.root
	for i := 0; i < len(key); i++ {
		child := cur.findChild(key[i])
		if child == nil {
			child = cur.newChild(key[i])
		}

		cur = child
	}

	cur.value = value
	return nil
}
