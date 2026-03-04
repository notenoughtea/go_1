package main

import "fmt"

// написать функцию которая вырезает символы
// из строки после определенного количества их повторов
// cat("awsdeabbd",1) => «awsdeb»
// cat("abbacdddba",2) => «abbacdd"

func cat(str string, n int) string {
	runes := []rune(str)
	result := []rune{}
	store := map[rune]int{}
	for i := 0; i < len(runes); i++ {
		num, ok := store[runes[i]]
		if num == n {
			continue
		}
		if !ok {
			store[runes[i]] = 1
			result = append(result, runes[i])
		} else {
			store[runes[i]] += 1
			result = append(result, runes[i])
		}
	}
	return string(result)
}

func main() {
	fmt.Println(cat("awsdeabbd", 1))
	fmt.Println(cat("abbacdddba", 2))
	fmt.Println(cat("абвгдааб", 1))

}
