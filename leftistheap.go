package main

import (
	"errors"
)

type LeftistNode struct {
	key   int
	left  *LeftistNode
	right *LeftistNode
	dist  int
}

type LeftistHeap struct {
	root *LeftistNode
}

func NewLeftistHeap() *LeftistHeap {
	return &LeftistHeap{}
}

func (h *LeftistHeap) isEmpty() bool {
	return h.root == nil
}

func getDist(node *LeftistNode) int {
	if node == nil {
		return 0
	}
	return node.dist
}

func merge(a, b *LeftistNode) *LeftistNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	if a.key > b.key {
		a, b = b, a
	}

	a.right = merge(a.right, b)

	if getDist(a.left) < getDist(a.right) {
		a.left, a.right = a.right, a.left
	}

	a.dist = getDist(a.right) + 1
	return a
}

// Merge сливает две кучи.
func (h *LeftistHeap) Merge(other *LeftistHeap) {
	h.root = merge(h.root, other.root)
	other.root = nil
}

func (h *LeftistHeap) Insert(key int) {
	newNode := &LeftistNode{
		key:  key,
		dist: 1,
	}
	h.root = merge(h.root, newNode)
}

func (h *LeftistHeap) DeleteMin() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("heap is empty")
	}
	min := h.root.key

	h.root = merge(h.root.left, h.root.right)
	return min, nil
}

func (h *LeftistHeap) PeekMin() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("heap is empty")
	}
	return h.root.key, nil
}
