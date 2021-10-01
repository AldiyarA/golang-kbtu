package main

import "fmt"

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func add(t *Tree, x int) *Tree {
	if t == nil {
		return &Tree{Value: x}
	}
	if t.Value == x {
		return t
	}
	if t.Value < x {
		t.Right = add(t.Right, x)
	} else {
		t.Left = add(t.Left, x)
	}
	return t
}
func (t *Tree) print() {
	ch := make(chan int)
	go Walk(t, ch)
	for {
		x, ok := <-ch
		if !ok {
			return
		}
		fmt.Println(x)
	}
}

func Walk2(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk2(t.Left, ch)
	ch <- t.Value
	Walk2(t.Right, ch)
}

func Walk(t *Tree, ch chan int) {
	Walk2(t, ch)
	close(ch)
}

func Same(t1, t2 *Tree) bool {
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
	var t1 *Tree
	var t2 *Tree
	for _, n := range []int{3, 4, 5, 1, 2} {
		t1 = add(t1, n)
	}
	for _, n := range []int{3, 1, 2, 4, 5} {
		t2 = add(t2, n)
	}
	fmt.Println("Tree 1")
	t1.print()
	fmt.Println("Tree 2")
	t2.print()
	fmt.Println(Same(t1, t2))
}
