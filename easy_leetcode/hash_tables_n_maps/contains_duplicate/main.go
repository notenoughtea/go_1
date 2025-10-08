package main

import "fmt"

// Дан целочисленный массив nums. Верните true, если какое-то значение встречается в массиве хотя бы дважды, и верните false, если каждый элемент встречается ровно один раз.

// Пример 1:
// Ввод: nums = [1,2,3,1]
// Вывод: true
// Объяснение: элемент 1 встречается в индексах 0 и 3.

// Пример 2:
// Ввод: nums = [1,2,3,4]
// Вывод: false
// Объяснение: все элементы различны.

// Пример 3:
// Ввод: nums = [1,1,1,3,3,4,3,2,4,2]
// Вывод: true

func containsDuplicate(nums []int) bool {
	store := make(map[int]int)

	for _, v := range nums {
		if store[v] == 1 {
			return true
		}
		store[v]++
	}
	return false
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(containsDuplicate(nums))
}
