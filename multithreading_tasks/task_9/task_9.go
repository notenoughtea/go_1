// Fan-In (объединение данных из нескольких каналов)
// Задача: Напишите функцию, которая объединяет два входных канала в один выходной.

package main

import (
	"fmt"
	"sync"
)

func Merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			for val := range c {
				out <- val
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in1 := make(chan int)
	in2 := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			in1 <- i
		}
		close(in1)
	}()

	go func() {
		for i := 10; i <= 12; i++ {
			in2 <- i
		}
		close(in2)
	}()

	for val := range Merge(in1, in2) {
		fmt.Println(val)
	}
}
