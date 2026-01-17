package resp

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
	// +OK\r\n
	case '+':
	//-Error message\r\n
	case '-':
	//:[<+|->]<value>\r\n
	case ':':
	//$<length>\r\n<data>\r\n
	case '$':
	//*<number-of-elements>\r\n<element-1>...<element-n>
	case '*':
	//_\r\n
	case '_':
	//#<t|f>\r\n
	case '#':
	//,[<+|->]<integral>[.<fractional>][<E|e>[sign]<exponent>]\r\n
	case ',':
	//([+|-]<number>\r\n
	case '(':
	case '!':
	//!<length>\r\n<error>\r\n
	case '=':
	//%<number-of-entries>\r\n<key-1><value-1>...<key-n><value-n>
	case '%':
	/*|1\r\n
	    +key-popularity\r\n
	    %2\r\n
	        $1\r\n
	        a\r\n
	        ,0.1923\r\n
	        $1\r\n
	        b\r\n
	        ,0.0012\r\n
	*2\r\n
	    :2039123\r\n
	    :9543892\r\n*/
	case '|':
	//~<number-of-elements>\r\n<element-1>...<element-n>
	case '~':
	//><number-of-elements>\r\n<element-1>...<element-n>
	case '>':
	default:
		return fmt.Errorf("unsupported RESP type: %q", t)
	}
	return nil
}
