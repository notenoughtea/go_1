package main

import "fmt"

func main() {
	var s *string
	fmt.Println(s == nil) //true
	var i interface{}
	fmt.Println(i == nil) //true
	i = s
	fmt.Println(i == nil) //false
}
