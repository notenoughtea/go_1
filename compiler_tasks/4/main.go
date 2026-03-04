package main

import "fmt"

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	q := &m["one"]
	// покажет ошибку так как нельзя брать адрес по ключу в мапе, он меняется
	// можно хранить в map указатели
	// 	m := map[string]*int{
	//     "one": new(int),
	// }

	fmt.Println(*q)
}
