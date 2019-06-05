package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int, 10)
	count := 1000
	var sum int
	sum = 0
	go func() {
		for i := 0; i < count; i++ {
			ch <- rand.Intn(100)
			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < count; i++ {
		sum += <-ch
		fmt.Println("메인루틴 : ", i)
		fmt.Println("Sum : ", sum)
	}

	fmt.Println("finished!!!")
}
