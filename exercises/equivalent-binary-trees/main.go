package main

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walkRecursive(t.Left, ch)
	ch <- t.Value
	walkRecursive(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	println("Values in the tree")
	for i := range ch {
		println(i)
	}

	same11 := Same(tree.New(1), tree.New(1)) // Should return true
	println("1 and 1 are same:", same11)

	same12 := Same(tree.New(1), tree.New(2)) // Should return false
	println("1 and 2 are same:", same12)

}
