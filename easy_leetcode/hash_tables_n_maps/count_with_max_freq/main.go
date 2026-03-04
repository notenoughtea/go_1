package main

import (
	"fmt"
)

// Дан массив nums, состоящий из положительных целых чисел.
// Верните сумму частот тех элементов массива, которые имеют максимальную частоту.

// Частота элемента — это число его вхождений в массив.

// Input: nums = [1,2,2,3,1,4]
// Output: 4
// Explanation: The elements 1 and 2 have a frequency of 2 which is the maximum frequency in the array.
// So the number of elements in the array with maximum frequency is 4.

func maxFrequencyElements(nums []int) int {
	store := make(map[int]int)
	maxFreq := 0
	for _, val := range nums {
		store[val]++
		if store[val] > maxFreq {
			maxFreq = store[val]
		}
	}
	final := 0
	for _, freq := range store {
		if freq == maxFreq {
			final += freq
		}
	}
	return final
}

func main() {
	nums := []int{10, 12, 11, 9, 6, 19, 11}
	fmt.Println(maxFrequencyElements(nums))
}
