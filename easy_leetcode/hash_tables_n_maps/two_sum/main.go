package main

import "fmt"

// Дан массив целых чисел nums и целое число target.
// Нужно вернуть индексы двух чисел,
// сумма которых равна target.

// Можно предположить,
// что у каждого входного массива есть ровно одно решение,
// и один и тот же элемент нельзя использовать дважды.

// Ответ можно вернуть в любом порядке.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

func twoSum(nums []int, target int) []int {
	store := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := store[target-nums[i]]; ok {
			return []int{j, i}
		}
		store[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	s := twoSum(nums, target)
	fmt.Println(s)
}
