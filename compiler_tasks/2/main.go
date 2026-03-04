package main

import "fmt"

func main() {
	x := []int{}     //len 0 cap 0
	x = append(x, 0) //len 1 cap 1
	x = append(x, 1) //len 2 cap 2
	x = append(x, 2) //len 3 cap 4

	y := append(x, 3) //len 3 cap 4

	z := append(x, 4) //len 3 cap 4

	fmt.Println(y, z) //x 0124 y 0124
}
