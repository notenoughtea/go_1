package main

import "fmt"

func main() {
	digits := []int{1, 2, 3, 4, 5}
	for _, d := range digits {
		defer fmt.Println(d)
	}
}

// так как defer работает по принципу стэка, ряд показан наоборот: 5 4 3 2 1
