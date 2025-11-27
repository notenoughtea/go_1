package main

// - Есть функция unpredictableFunc, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).

// - Нужно написать обёртку predictableFunc,
// которая будет работать с заданным фиксированным таймаутом
// (например, 1 секунду).

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос) .
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd

}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести
// Сигнатуру функцию обёртки менять можно.

func predictableFunc() int64 {
	ch := make(chan int64)
	timer := time.NewTicker(1000 * time.Millisecond)
	deadline := time.Now().Add(1000 * time.Second)
	left := time.Until(deadline)
	if left < 0 {
		left = 0
	}
	go func() {
		resp := unpredictableFunc()
		ch <- resp
		defer timer.Stop()
		close(ch)
	}()
	select {
	case resp := <-ch:
		fmt.Println("Результат: ", resp)
		fmt.Println("Функция выполнена за ", left)
		return resp
	case <-timer.C:
		fmt.Println("Время вышло")
	}
	return 0
}

func main() {
	fmt.Println("started")
	fmt.Println(predictableFunc())
}
