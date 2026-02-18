package main

import (
	"errors"
)

type SkewNode struct {
	key   int
	left  *SkewNode
	right *SkewNode
}

type SkewHeap struct {
	root *SkewNode
}

func NewSkewHeap() *SkewHeap {
	return &SkewHeap{}
}

func (h *SkewHeap) isEmpty() bool {
	return h.root == nil
}

func merge(a, b *SkewNode) *SkewNode {
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

	a.left, a.right = a.right, a.left
	return a
}

func (h *SkewHeap) Merge(other *SkewHeap) {
	h.root = merge(h.root, other.root)
	other.root = nil
}

func (h *SkewHeap) Insert(key int) {
	newNode := &SkewNode{key: key}
	h.root = merge(h.root, newNode)
}

func (h *SkewHeap) DeleteMin() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("heap is empty")
	}
	min := h.root.key

	h.root = merge(h.root.left, h.root.right)
	return min, nil
}

func (h *SkewHeap) PeekMin() (int, error) {
	if h.isEmpty() {
		return 0, errors.New("heap is empty")
	}
	return h.root.key, nil
}
