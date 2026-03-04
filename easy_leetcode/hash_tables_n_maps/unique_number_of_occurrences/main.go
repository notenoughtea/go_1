package main

import "fmt"

// Дан массив целых чисел arr.
// Нужно вернуть true, если количество вхождений каждого значения в массиве уникально
// (не повторяется у разных чисел), и false в противном случае.

// Input: arr = [1,2,2,1,1,3]
// Output: true
// Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

func uniqueOccurrences(arr []int) bool {
	store := make(map[int]int)

	for _, val := range arr {
		if s, ok := store[val]; ok {
			store[val] = s + 1
		} else {
			store[val] = 1
		}
	}
	have := make(map[int]bool)
	for _, freq := range store {
		if have[freq] {
			return false
		}
		have[freq] = true
	}
	return true
}

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
}
