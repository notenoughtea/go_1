package main

import "fmt"

// Задача: определить, является ли строка панграммой — то есть содержит ли она все 26 букв английского алфавита хотя бы по одному разу.

// Пример 1:
// Ввод: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Вывод: true
// Пояснение: строка включает все буквы английского алфавита.

// Пример 2:
// Ввод: sentence = "leetcode"
// Вывод: false
// Пояснение: в строке не хватает многих букв.

var alphabet = []string{
	"a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z",
}

func checkIfPangram(s string) bool {
	store := make(map[string]int)
	for _, v := range alphabet {
		store[v] = 0
	}
	fmt.Println(store)
	runes := []rune(s)
	for _, v := range runes {
		store[string(v)]++
	}
	for _, v := range store {
		if v == 0 {
			return false
		}
	}
	return true
}

func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	fmt.Println(checkIfPangram(sentence))
}
