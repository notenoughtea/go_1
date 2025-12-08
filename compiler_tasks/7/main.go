package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	var buf *bytes.Buffer // избежать  провала nil в f: 	buf := &bytes.Buffer{} или 	buf := new(bytes.Buffer)
	f(buf)
	if buf != nil {
		fmt.Printf(buf.String())
	}
	fmt.Println("Main completed")
}

func f(out io.Writer) {
	if out != nil {
		_, err := out.Write([]byte("Hello world\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
