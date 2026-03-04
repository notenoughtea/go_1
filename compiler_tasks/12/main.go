package main

import "fmt"

type X struct {
	Val int
}

func (x X) S() {
	fmt.Println(x.Val)
}

func main() {
	x := X{10}
	defer x.S() // 10 покажет, значение зафиксировалось в этот момент из за defer
	x.Val = 256
}
