package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

func GetFile(ctx context.Context, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:

	}

	if strings.HasPrefix(name, "invalid") {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	b := make([]byte, 10)
	n, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("getting file %q: %w", name, err)
	}

	return b[:n], nil
}

// GetFilesOld пример функции, которую нужно оптимизировать.
// Менять эту функцию не нужно, она нужна чтоб
// сравнить поведение двух функций после оптимизации
func GetFilesOld(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	result = make(map[string][]byte, len(names))
	for _, name := range names {
		result[name], err = GetFile(ctx, name)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// GetFilesNew эту функцию можно менять, за исключением её
// сигнатуры
func GetFilesNew(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	result = make(map[string][]byte, len(names))

	ctx1, cancel := context.WithCancel(ctx)
	defer cancel()
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	if len(names) == 0 {
		return nil, nil
	}

	chErr := make(chan error, len(names))
	for _, name := range names {
		n := name
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := GetFile(ctx1, n)
			if err != nil {
				chErr <- err
				return
			}
			mu.Lock()
			result[n] = r
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

	return result, nil
}

func main() {
	start := time.Now()
	files, err := GetFilesNew(context.TODO(), "1", "2")
	if err != nil {
		log.Fatalln(time.Since(start), err)
	}

	fmt.Println(time.Since(start))
	for name := range files {
		fmt.Println(name)
	}
}

//вообще хз что ту нужно было, по тз, но вот что я сделал:
// обернул мутексом резалт чтобы небыло гонки, добавил WG чтобы дождаться всех горутин
