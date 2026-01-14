package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) ReadLine() (string, error) {
	line, err := s.r.ReadString('\n')

	if err != nil {
		return "", err
	}

	if len(line) < 2 || line[len(line)-2] != '\r' {
		return "", fmt.Errorf("protocol error: expected CRLF")
	}

	return line, nil
}

func (s *Scanner) ReadBulkString() (int, string, error) {
	return 0, "", nil
}

func (s *Scanner) Read() error {
	t, err := s.r.ReadByte()

	if err != nil {
		return err
	}

	switch t {

	case '+':
		line, err := s.ReadLine()
		if err != nil {
			return err
		}
		value := line[:len(line)-2]
		fmt.Printf("SimpleString: %s\n", value)

	case '-':
		line, err := s.ReadLine()
		if err != nil {
			return err
		}
		value := line[:len(line)-2]

		fmt.Printf("SimpleError: %s\n", value)

	//:[<+|->]<value>\r\n
	case ':':
		line, err := s.ReadLine()
		if err != nil {
			return err
		}
		value := line[:len(line)-2]

		i, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return err
		}
		fmt.Printf("Integer: %d\n", i)

	//$<length>\r\n<data>\r\n
	case '$':
	//*<number-of-elements>\r\n<element-1>...<element-n>
	case '*':

	default:
		return fmt.Errorf("unsupported RESP type: %q", t)
	}
	return nil
}

func Split(r rune) bool {
	return r == '\n'
}
