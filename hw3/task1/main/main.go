package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk2(n *tree.Tree, ch chan int) {
	if n == nil {
		return
	}
	Walk2(n.Left, ch)
	ch <- n.Value
	Walk2(n.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	Walk2(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for {
		x, ok1 := <-ch1
		y, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if !ok1 || !ok2 {
			return false
		}
		if x != y {
			return false
		}
	}
}
func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
