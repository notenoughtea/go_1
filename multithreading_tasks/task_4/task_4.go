//**Потокобезопасный инкремент - Mutex.**

// Задача: Напишите программу,
// где 10 горутин инкрементируют один счётчик,
// защищая его sync.Mutex.
// Что если не обложить мютексом?
// Воспроизвести race condition.
package main

import (
	"fmt"
	"sync"
)

type counter struct {
	c  int
	mu sync.Mutex
}

func (c *counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.c++
	fmt.Println(c.c)
	c.mu.Unlock()
}

func main() {
	counter := &counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go counter.increment(&wg)
	}
	wg.Wait()
}

// без мютексов будет гонка данных
// вот вариант с гонкой:

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type counter struct {
// 	c int
// }

// func (c *counter) increment(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	c.c++
// 	fmt.Println(c.c)
// }

// func main() {
// 	counter := &counter{}
// 	wg := sync.WaitGroup{}

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go counter.increment(&wg)
// 	}
// 	wg.Wait()
// }
