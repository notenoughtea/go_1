// > 📦 ЗАДАЧА: Батчевая обработка
// >
// >
// > **Условие**:
// >
// > Реализуй функцию `StartBatchProcessor(ctx context.Context, input <-chan int)`, которая:
// >
// > - Собирает числа из канала `input` в батчи по максимум 5 элементов.
// > - Если в течение 2 секунд батч не собран — обрабатывает то, что есть.
// > - Обработка батча — это просто `fmt.Println("Processed batch:", batch)`.
// > - Выход из функции должен происходить при отмене контекста (`ctx.Done()`).
// >
// > **Дополнительно**:
// >
// > - Отмена должна происходить либо через `context.WithTimeout`, либо вручную через `cancel()` — попробовать оба варианта
// >
// >  Начальный код с вызовом(доработать)

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func StartBatchProcessor(ctx context.Context, input <-chan int) {
	batch := []int{}

	for {
		select {
		case i, ok := <-input:
			if !ok {
				if len(batch) > 0 {
					fmt.Println("Processed batch:", batch)
				}
				return
			}
			batch = append(batch, i)
			if len(batch) == 5 {
				fmt.Println("Processed batch:", batch)
				batch = []int{}
			}
		case <-ctx.Done():
			if len(batch) > 0 {
				fmt.Println("Processed batch:", batch)
			}
			return
		}
	}

}

func main() {
	// инициализация канала
	/* создание контекста  */

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	input := make(chan int, 1)

	go StartBatchProcessor(ctx, input)

	go func() {
		for i := 0; i < 20; i++ {
			dur := time.Duration(rand.Intn(300))
			input <- i
			time.Sleep(dur * time.Millisecond)
		}
	}()

	defer close(input)
	// сбор данных
	<-ctx.Done()
	fmt.Println("Main: processing stopped")
}
