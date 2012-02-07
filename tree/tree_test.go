package tree

import (
	"log"
	"testing"
)

func cmp(a, b Item) int { return a.(int) - b.(int) }

func TestTree(t *testing.T) {
	tree := NewTree(cmp)
	for _, j := range []int{2, 4, 6, 8, 1, 3, 5, 7, 9} {
		tree.Insert(j)
	}
	log.Print("iterating")
	for i := range tree.Iterate(InOrder) {
		log.Print(i.(int))
	}
}
