package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/itsjayeshrathi/tinyresp/internal/resp"
)

func main() {
	input := "+OK\r\n+PONG\r\n-ERR unknown command 'asdf'\r\n!21\r\nSYNTAX invalid syntax\r\n:0\r\n:-10000\r\n_\r\n$5\r\nhello\r\n+OK\r\n=15\r\ntxt:Some string\r\n"

	reader := strings.NewReader(input)

	scanner := resp.NewScanner(reader)

	for {
		any, err := scanner.Read()
		fmt.Println(any)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
