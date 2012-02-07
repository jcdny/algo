package tree

import (
	"log"
)

type Item interface{}

type Element struct {
	Item
	l, r *Element
}

type Tree struct {
	root *Element
	len  int
	cmp  CompareFunc
}

type CompareFunc func(a, b Item) int

// NewTree returns a new empty root element
func NewTree(cmp CompareFunc) *Tree {
	return &Tree{cmp: cmp}
}

func (t *Tree) Init(cmp CompareFunc) {
	t.len = 0
	t.root = nil
	t.cmp = cmp
}

func (t *Tree) Insert(item Item) {
	log.Print("insert")
	if t.root == nil {
		t.root = &Element{Item: item}
		t.len++
	} else {
		parent := t.root
		for {
			cmp := t.cmp(parent.Item, item)
			if cmp > 0 {
				if parent.l == nil {
					parent.l = &Element{Item: item}
					t.len++
					break
				}
				parent = parent.l
			} else if cmp < 0 {
				if parent.r == nil {
					parent.r = &Element{Item: item}
					t.len++
					break
				}
				parent = parent.r
			} else {
				// replace
				parent.Item = item
				break
			}
		}
	}
}

func (t *Tree) Iterate(f func(*Element, chan<- Item)) <-chan Item {
	c := make(chan Item)
	go func() {
		f(t.root, c)
		close(c)
	}()

	return c
}

// Iterate: in order traversal.
func InOrder(e *Element, c chan<- Item) {
	if e == nil {
		return
	}
	InOrder(e.l, c)
	c <- e.Item
	InOrder(e.r, c)
}

// Iterate: preorder traversal.
func PreOrder(e *Element, c chan<- Item) {
	if e == nil {
		return
	}
	c <- e.Item
	PreOrder(e.l, c)
	PreOrder(e.r, c)
}

// Iterate: postorder traversal.
func PostOrder(e *Element, c chan<- Item) {
	if e == nil {
		return
	}
	c <- e.Item
	PostOrder(e.l, c)
	PostOrder(e.r, c)
}
