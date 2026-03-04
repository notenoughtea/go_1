package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		panic(123)
	}()

	time.Sleep(time.Hour)
}

//Этот код содержит defer с recover(), но он не сможет поймать panic(123),
// так как паника происходит в горутине,
// а recover() может обработать panic только в той же горутине,
// в которой был объявлен defer. В результате программа завершится с panic: 123
