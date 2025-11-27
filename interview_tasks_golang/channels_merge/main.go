package main

import "fmt"

// асинхронное слияние каналов

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4

	close(ch1)
	close(ch2)

	ch3 := asyncMerge(ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

// todo
func asyncMerge[T int](ch1 chan T, ch2 chan T) chan T {
	ch3 := make(chan T, 30)
	go func() {
		defer close(ch3)

		ch1Open, ch2Open := true, true

		for ch1Open || ch2Open {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1Open = false
					ch1 = nil
					continue
				}
				ch3 <- v
			case v, ok := <-ch2:
				if !ok {
					ch2Open = false
					ch2 = nil
					continue
				}
				ch3 <- v
			}
		}
	}()
	return ch3
}
