// Fan-Out (разделение работы между несколькими воркерами)
// Задача: Создайте функцию,
// которая принимает канал с задачами и
// распределяет их между N горутинами
package main

import (
	"fmt"
	"sync"
)

func Merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		a <- 1
		a <- 2
		a <- 3
		close(a)
	}()

	go func() {
		b <- 10
		b <- 20
		b <- 30
		close(b)
	}()

	for v := range Merge(a, b) {
		fmt.Println(v)
	}
}

// Разделить канал на n каналов,
// которые получают сообщения в циклическом порядке.
// func Split(ch <-chan int, n int) []<-chan int {

//   // Создаем пул из n каналов
//   cs := make([]chan int)
//   for i := 0; i < n; i++ {
//     cs = append(cs, make(chan int))
//   }

//   // Распределяет работу в круговом порядке
//   // среди указанного числа каналов,
//   // пока основной канал не будет закрыт.
//   // При закрытии основного канала закрывает
//   // все каналы и возвращается.
//   toChannels := func(ch <-chan int, cs []chan<- int) {

//     // Закрываем каждый канал,
//     // когда выполнение заканчивается.
//     defer func(cs []chan<- int) {
//       for _, c := range cs {
//         close(c)
//       }
//     }(cs)

//     // Направляем сообщения из
//     // основного канала ch
//     // в каналы из пула cs
//     for {
//       for _, c := range cs {
//         select {
//         case val, ok := <-ch:
//           if !ok {
//             return
//           }

//           c <- val
//         }
//       }
//     }
//   }

//   go toChannels(ch, cs)

//   return cs
// }
