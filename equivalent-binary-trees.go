package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if r1, r2 := <-ch1, <-ch2; r1 != r2 {
			return false
		}
	}
	return true
}

func EquivalentBinaryTrees() {
	test1 := Same(tree.New(1), tree.New(1))
	test2 := !Same(tree.New(1), tree.New(2))
	switch {
	case test1 && test2:
		fmt.Println("ALL TESTS PASSED")
	case test1 && !test2:
		fmt.Println("TEST 2 FAILED")
	case !test1 && test2:
		fmt.Println("TEST 1 FAILED")
	default:
		fmt.Println("ALL TESTS FAILED")
	}
}
