package main

import "fmt"

func add(s []string) {
	s = append(s, "x")
}

func main() {
	s := []string{"a", "b", "c"}
	add(s[1:2])    //выбираем b и добавляем после нее x(s[1:2] смотрит на старый массив)
	fmt.Println(s) // a b x
}
