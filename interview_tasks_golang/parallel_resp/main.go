package main

// Написать код, который будет выводить коды ответов на НТТР-запросы по двум URL адресам (например главная
// 	страница Google и главная страница Avito)

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	addrSlice := []string{
		"google.com",
		"avito.ru",
	}

	for i := 0; i < len(addrSlice); i++ {
		wg.Add(1)
		go func() {
			m, err := http.Get("http://" + addrSlice[i])
			if err != nil {
				panic(err)
			}
			fmt.Println(m.StatusCode)
			wg.Done()
		}()
	}
	wg.Wait()
}
