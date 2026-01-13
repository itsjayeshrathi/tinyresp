package main

import (
	"bufio"
	"fmt"
	"io"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Read() error {
	t, err := s.r.ReadByte()

	if err != nil {
		return err
	}

	switch t {
	case '+':
		line, err := s.r.ReadString('\n')
		if err != nil {
			return err
		}
		fmt.Println(line)
	default:
		return fmt.Errorf("unsupported RESP type: %q", t)
	}
	return nil
}
