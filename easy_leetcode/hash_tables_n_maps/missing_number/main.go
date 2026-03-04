package main

import (
	"fmt"
	"sort"
)

// Дан массив nums, содержащий n различных чисел в диапазоне [0, n].
// Нужно вернуть единственное число из этого диапазона,
// отсутствующее в массиве.

// Пример 1:
// Ввод: nums = [3,0,1]
// Вывод: 2

// Пояснение:
// Здесь n = 3, так как в массиве три числа.
// Следовательно, полный диапазон — [0,3].
// Недостающим числом является 2, так как оно не встречается в nums.

func missingNumber(nums []int) int {
	sorted := nums
	sort.Ints(sorted)
	l := len(sorted)
	for i := 0; i < l; i++ {
		if sorted[i] == i+1 {
			return sorted[i] - 1
		}
		if sorted[i] == i-1 {
			return sorted[i] + 1
		}
	}
	return sorted[l-1] + 1
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}
