package main

import "fmt"

// Вам дан массив paths, где paths[i] = [cityA_i, cityB_i] означает,
// что существует прямой путь из cityA_i в cityB_i.
// Верните конечный город — то есть город, из которого нет исходящих путей в другие города.

// Гарантируется, что граф, заданный paths,
// представляет собой линию без циклов,
// поэтому конечный город существует и он единственный.

// Пример 1:

// Ввод: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
// Вывод: "Sao Paulo"
// Пояснение: начав из "London", вы дойдёте до "Sao Paulo",
// который является конечным городом.
// Маршрут: "London" -> "New York" -> "Lima" -> "Sao Paulo".

// Пример 2:

// Ввод: paths = [["B","C"],["D","B"],["C","A"]]
// Вывод: "A"
// Пояснение: возможные маршруты:
// "D" -> "B" -> "C" -> "A"
// "B" -> "C" -> "A"
// "C" -> "A"
// "A"
// Очевидно, что конечный город — "A".

// Пример 3:

// Ввод: paths = [["A","Z"]]
// Вывод: "Z".

func destCity(paths [][]string) string {
	store := map[string][]string{}

	for _, v := range paths {
		store[v[0]] = []string{}
		store[v[1]] = []string{}
	}

	for _, v := range paths {
		store[v[0]] = append(store[v[0]], v[1])
	}

	for key, v := range store {
		if len(v) == 0 {
			return key
		}
	}
	return ""
}

func main() {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(destCity(paths))
}
