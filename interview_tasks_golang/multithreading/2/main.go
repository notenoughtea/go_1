package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println(Do(context.Background(), []User{{"aaa"}, {"bbb"}, {"ccc"}, {"ddd"}, {"eee"}}))
}

type User struct {
	Name string
}

func fetchByName(ctx context.Context, userName string) (int, error) {
	// Тут происходит сетевой поход, который по userName возвращает userID
	time.Sleep(10 * time.Millisecond) // Имитация сетевого похода
	return rand.Int() % 100000, nil
}

func Do(ctx context.Context, users []User) (map[string]int, error) {
	collected := make(map[string]int)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	chErr := make(chan error, len(users))

	for _, u := range users {
		usr := u
		wg.Add(1)
		go func() {
			defer wg.Done()
			userID, err := fetchByName(ctx, usr.Name)
			if err != nil {
				chErr <- err
				return
			}
			mu.Lock()
			collected[usr.Name] = userID
			mu.Unlock()
		}()
	}

	wg.Wait()
	close(chErr)

	for err := range chErr {
		if err != nil {
			cancel()
			return nil, err
		}
	}

	return collected, nil
}

// обернул мутексом мапу, добавил WG и канал с ошибками
