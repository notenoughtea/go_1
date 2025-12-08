package main

import "fmt"

func main() {
	a := []int{0, 0, 0}
	fmt.Println(a) // 0, 0, 0

	mod1(a)
	fmt.Println(a) // 1, 0, 0

	mod2(a)
	fmt.Println(a) // 1, 0, 3
}

func mod1(s []int) {
	if len(s) > 0 {
		s[0] = 1
	}
}

func mod2(s []int) {
	s[len(s)-1] = 3
	s = append(s, 4, 5) //s указывает на новый слайс так как 3 и 5 не поместились в старый
}
