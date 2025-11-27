package main

import (
	"fmt"
	"sync"
	"time"
)

type SafePrice struct {
	mu sync.Mutex
	m  map[string]int
}

func RunProcessor(wg *sync.WaitGroup, prices []*SafePrice) {
	go func() {
		defer wg.Done()
		for _, p := range prices {
			p.mu.Lock()
			for k, v := range p.m {
				p.m[k] = v + 1
			}
			fmt.Println(p.m)
			p.mu.Unlock()
		}
	}()
}

func RunWriter() <-chan *SafePrice {
	ch := make(chan *SafePrice)

	go func() {
		defer close(ch)

		current := map[string]int{
			"AAPL": 163,
			"USD":  117,
			"EUR":  124,
			"NVDA": 234,
		}

		for i := 0; i < 4; i++ {
			newMap := make(map[string]int, len(current))
			for k, v := range current {
				newMap[k] = int(float64(v) * 1.3)
			}
			current = newMap

			sp := &SafePrice{m: newMap}
			ch <- sp

			time.Sleep(100 * time.Millisecond)
		}
	}()

	return ch
}

func main() {
	p := RunWriter()
	var prices []*SafePrice

	for i := range p {
		i.mu.Lock()
		fmt.Println(i.m)
		i.mu.Unlock()
		prices = append(prices, i)
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		RunProcessor(&wg, prices)
	}

	wg.Wait()
}

// - Поправить код чтобы не было дедлока. ---закрываю канал
// Что будет выведено? Поправить бесконечный цикл. ---читаю через for range
// Выводятся мапы с одинаковыми значениями. Почему? ---перезаписывается мапа
// Поправить чтобы была видна история изменений. ---вывожу в консоль сразу из канала

// - Расскомментировать код в main и RunProcessor.
// Что будет выведено? Поправить дедлок (проблема в вейтгруппе которая передается не поинтером).
// Теперь код отработает без проблем?
// Будет паника при параллельном изменении мап (изменение слайса мап без лока, гонка данных).
// Как поправить? Добавить мьютекс при чтении и записи в мапу.
