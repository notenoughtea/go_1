package main

import (
	"fmt"
)

type myError struct {
	code int
}

func (e myError) Error() string {
	return fmt.Sprintf("code: %d", e.code)
}

func run() error {
	var e *myError
	if false {
		e = &myError{code: 123}
		//Компилятор выкидывает тело if false
		//код внутри не влияет на выполнение
		//но проверяет типы внутри(должно быть корректно по типам, иначе код не скомпилируется)
	}
	return e
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("failed to run, error: ", err)
	} else {
		fmt.Println("success")
	}
}

//Этот код демонстрирует нюанс работы с nil в Go.
// В run() переменная e объявляется как *myError, но не инициализируется (остается nil).
// При возврате e из run() оно приводится к error, но так как e имеет конкретный тип *myError,
// оно неявно становится error, содержащим nil, что не равно nil в Go.
// Поэтому err != nil вернет true, и программа выведет "failed to run, error: <nil>".
