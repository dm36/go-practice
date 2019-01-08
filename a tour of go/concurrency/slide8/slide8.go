package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

// Think pre-order, inorder, post-order traversal- this
// example is inorder traversal (left, root then right)
func WalkTree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkTree(t.Right, ch)
	}
}

// Doing this so we can close the channel and signal so to the receiving
// channel
func Walk(t *tree.Tree, ch chan int) {
	WalkTree(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var a1 []int
	var a2 []int
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch)
	go Walk(t2, ch2)

	for i := range ch {
		a1 = append(a1, i)
	}

	for i := range ch2 {
		a2 = append(a2, i)
	}

	fmt.Println(a1)
	fmt.Println(a2)

	for ind, _ := range a1 {
		if a1[ind] != a2[ind] {
			return false
		}
	}
	return true

	// Not using an array:
	// for i := range ch1 {
	// 	if i != <-ch2 {
	// 		return false
	// 	}
	// }
	// return true
}

func main() {
	// Make a channel of integers of size 10, run a goroutine
	// that calls the function go with a tree and a channel
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)

	// Read values off the channel
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
