// Потокобезопасный инкремент - Atomic.**

// Задача: Напишите программу,
// где 10 горутин инкрементируют один счётчик
// без использования мютексов, через атомики.

// Что, если не использовать атомик?
// Что лучше, атомик или мютекс?

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	c int64
}

func (c *counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&c.c, 1)
	fmt.Println(c.c)
}

func main() {
	wg := &sync.WaitGroup{}
	counter := counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go counter.increment(wg)
	}
	wg.Wait()
}

//если не использовать атомик юзаю мютексы
// разница в использовании типов: атомики-примитивные типы и быстрее, мютексы-любые типы и медленнее
