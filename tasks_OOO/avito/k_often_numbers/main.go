package main

import (
	"fmt"
	"sort"
)

// Дан массив целых чисел nums и целое число k.
// Нужно написать функцию, которая вынимает из массива k
// наиболее часто встречающихся элементов.

// Пример
// # ввод
// nums = [1,1,1,2,2,3]
// k = 2
// # вывод (в любом порядке)
// [1, 2]

func CountOftenElements(elems []int, n int) []int {
	store := map[int]int{}
	for _, x := range elems {
		store[x]++
	}

	type pair struct {
		val  int
		freq int
	}
	pairs := make([]pair, 0, len(store))
	for k, v := range store {
		pairs = append(pairs, pair{val: k, freq: v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	result := make([]int, 0, n)
	for i := 0; i < n && i < len(pairs); i++ {
		result = append(result, pairs[i].val)
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 8, 8, 8, 8, 8}
	k := 2
	fmt.Println(CountOftenElements(nums, k))
}
