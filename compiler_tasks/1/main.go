package main

import "fmt"

func main() {
	testSlice := make([]*int, 0)

	for _, v := range []int{1, 2, 3, 4} {
		testSlice = append(testSlice, &v)
	}

	for _, v := range testSlice {
		fmt.Println(*v)
	}
}

// до 1.22: 4,4,4,4
// после 1.22: 1,2,3,4
