package main

import (
	"io"
	"log"
	"strings"
)

func main() {
	input := "+OK\r\n+PONG\r\n"

	reader := strings.NewReader(input)

	scanner := NewScanner(reader)

	for {
		err := scanner.Read()
		if err != nil {
			continue
		}
		if err == io.EOF {
			break
		}
		log.Fatal(err)
	}
}
