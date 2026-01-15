package main

import (
	"io"
	"log"
	"strings"
	 "github.com/itsjayeshrathi/tinyresp/internal/resp"
)

func main() {
	input := "+OK\r\n+PONG\r\n-ERR unknown command 'asdf'\r\n:0\r\n:10000\r\n$5\r\nhello\r\n"

	reader := strings.NewReader(input)

	scanner := NewScanner(reader)

	for {
		err := scanner.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
