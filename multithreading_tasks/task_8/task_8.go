// Практическая задача:
// Конкурентный HTTP-запрос к списку URL-адресов с использованием горутин и синхронизации

// > **Задача:**
// **Напишите функцию FetchURLs(urls []string) map[string]string, которая:**

// *Принимает слайс URL-адресов.
// Конкурентно делает HTTP-запросы к каждому URL.
// Собирает результаты (код ответа и часть тела) в map[string]string, где:
// ключ — URL
// значение — содержимое ответа (ограниченное, например, 100 символами)
// Использует sync.WaitGroup и sync.Mutex для защиты записи в map.
// В случае ошибки записывает "error" как значение.*

// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// var urlList = []string{
// 	"https://api.chucknorris.io/jokes/random?category=animal",
// 	"https://api.chucknorris.io/jokes/random?category=career",
// 	"https://api.chucknorris.io/jokes/random?category=celebrity",
// 	"https://api.chucknorris.io/jokes/random?category=dev",
// 	"https://api.chucknorris.io/jokes/random?category=explicit",
// 	"https://api.chucknorris.io/jokes/random?category=fashion",
// 	"https://api.chucknorris.io/jokes/random?category=food",
// 	"https://api.chucknorris.io/jokes/random?category=history",
// 	"https://api.chucknorris.io/jokes/random?category=money",
// 	"https://api.chucknorris.io/jokes/random?category=movie",
// 	"https://api.chucknorris.io/jokes/random?category=music",
// 	"https://api.chucknorris.io/jokes/random?category=political",
// 	"https://api.chucknorris.io/jokes/random?category=religion",
// 	"https://api.chucknorris.io/jokes/random?category=science",
// 	"https://api.chucknorris.io/jokes/random?category=sport",
// 	"https://api.chucknorris.io/jokes/random?category=travel",
// }

// type store struct {
// 	st map[string]string
// 	mt sync.Mutex
// }

// type JokeResponse struct {
// 	Value string `json:"value"`
// }

// func FetchURLs(ctx context.Context, urls []string) map[string]string {

// 	wg := sync.WaitGroup{}

// 	store := store{
// 		st: make(map[string]string),
// 	}

// 	for i := 0; i < len(urls); i++ {
// 		wg.Add(1)
// 		go func() {
// 			select {
// 			case <-ctx.Done():
// 				store.mt.Lock()
// 				store.st[urls[i]] = "canceled"
// 				store.mt.Unlock()
// 				return
// 			default:
// 			}
// 			store.mt.Lock()
// 			defer store.mt.Unlock()

// 			client := &http.Client{
// 				Timeout: 3 * time.Millisecond,
// 			}

// 			resp, err := client.Get(urls[i])
// 			if err != nil {
// 				panic(err)
// 			}
// 			body, err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				panic(err)
// 			}
// 			var joke JokeResponse
// 			json.Unmarshal(body, &joke)
// 			store.st[urls[i]] = resp.Status + " " + joke.Value + "\n"
// 			defer resp.Body.Close()
// 			defer wg.Done()
// 			fmt.Println("HTTP Status Code:", resp.StatusCode)
// 		}()
// 	}

// 	wg.Wait()
// 	return store.st
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	fmt.Println(FetchURLs(ctx, urlList))
// }

// а тут с worker pool

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var urlList = []string{
	"https://api.chucknorris.io/jokes/random?category=animal",
	"https://api.chucknorris.io/jokes/random?category=career",
	"https://api.chucknorris.io/jokes/random?category=celebrity",
	"https://api.chucknorris.io/jokes/random?category=dev",
	"https://api.chucknorris.io/jokes/random?category=explicit",
	"https://api.chucknorris.io/jokes/random?category=fashion",
	"https://api.chucknorris.io/jokes/random?category=food",
	"https://api.chucknorris.io/jokes/random?category=history",
	"https://api.chucknorris.io/jokes/random?category=money",
	"https://api.chucknorris.io/jokes/random?category=movie",
	"https://api.chucknorris.io/jokes/random?category=music",
	"https://api.chucknorris.io/jokes/random?category=political",
	"https://api.chucknorris.io/jokes/random?category=religion",
	"https://api.chucknorris.io/jokes/random?category=science",
	"https://api.chucknorris.io/jokes/random?category=sport",
	"https://api.chucknorris.io/jokes/random?category=travel",
}

type JokeResponse struct {
	Value string `json:"value"`
}

type store struct {
	st map[string]string
	mt sync.Mutex
}

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan string, s *store) {
	defer wg.Done()
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	for url := range jobs {
		select {
		case <-ctx.Done():
			s.mt.Lock()
			s.st[url] = "canceled"
			s.mt.Unlock()
			continue
		default:
		}

		resp, err := client.Get(url)
		if err != nil {
			s.mt.Lock()
			s.st[url] = "error"
			s.mt.Unlock()
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			s.mt.Lock()
			s.st[url] = "error"
			s.mt.Unlock()
			continue
		}

		var joke JokeResponse
		if err := json.Unmarshal(body, &joke); err != nil {
			s.mt.Lock()
			s.st[url] = "error"
			s.mt.Unlock()
			continue
		}

		result := joke.Value
		if len(result) > 100 {
			result = result[:100]
		}

		s.mt.Lock()
		s.st[url] = result
		s.mt.Unlock()
	}
}

func FetchURLs(ctx context.Context, urls []string) map[string]string {
	s := &store{
		st: make(map[string]string),
	}

	jobs := make(chan string)
	wg := sync.WaitGroup{}

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, jobs, s)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	wg.Wait()
	return s.st
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := FetchURLs(ctx, urlList)
	for url, content := range result {
		fmt.Printf("%s -> %s\n", url, content)
	}
}
