package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree) {
	if t == nil {
		return
	}
	Walk(t.Left)
	fmt.Println(t.Value)
	Walk(t.Right)
}

func Walk_with_channel(t *Tree, ch chan int) {
	if t != nil {
		if t.Left != nil {
			Walk_with_channel(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil {
			Walk_with_channel(t.Right, ch)
		}
	}
}

func same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer close(ch1)
		Walk_with_channel(t1, ch1)
	}()
	go func() {
		defer close(ch2)
		Walk_with_channel(t2, ch2)
	}()

	var v2 int
	for v1 := range ch1 {
		v2 = <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func Insert(data int, t *Tree) {
	if data < t.Value {
		if t.Left != nil {
			Insert(data, t.Left)
		} else {
			t.Left = &Tree{Value: data}
		}
	} else {
		if t.Right != nil {
			Insert(data, t.Right)
		} else {
			t.Right = &Tree{Value: data}
		}
	}
}

func make_tree(a []int, t *Tree) *Tree {
	for _, i := range a {
		Insert(i, t)
	}
	return t
}

func main() {
	t1 := &Tree{Value: 0}
	a1 := []int{7, -2, 8, -9, 4, 5}
	t1 = make_tree(a1, t1)

	t2 := &Tree{Value: 8}
	a2 := []int{4, -2, 0, -9, 7, 5}
	t2 = make_tree(a2, t2)

	fmt.Println("--------------t1-------------")
	Walk(t1)
	fmt.Println("--------------t2-------------")
	Walk(t2)

	fmt.Println("--------- Are those Same? ---------")
	if same(t1, t2) {
		fmt.Println("Same tree!")
	} else {
		fmt.Println("Different tree!")
	}
}
