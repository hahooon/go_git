package main

import "fmt"

func fibonacci() func() int {
	f0 := 0
	f1 := 0
	return func() int {
		f2 := f0 + f1
		f0 = f1
		f1 = f2
		if f1 == 0 {
			f1 = 1
		}
		return f2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
