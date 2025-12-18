// Что выведет следующая программа и почему?
package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	*person = Person{ //тут пишем по указателю теперь
		Name: "Alice",
	}
}

func main() {
	person := &Person{
		Name: "Bob",
	}
	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
}

/*
Ответ:
Bob
Bob

Как модифицировать программу, чтобы вывелось:
Bob
Alice
*/
