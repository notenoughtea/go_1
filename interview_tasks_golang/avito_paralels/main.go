package main

//Сколько будет выполняться программа? Как оптимизировать?

//Ответ: мутекс на каунт и запускать "запросы" параллельно в вейт группе

import (
	"fmt"
	"sync"
	"time"
)

const numRequests = 10000

var wg sync.WaitGroup
var mu sync.Mutex
var count int

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Println("Счет:", count)
	wg.Done()
}

func main() {
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go networkRequest()
	}

	fmt.Println(count)
}
