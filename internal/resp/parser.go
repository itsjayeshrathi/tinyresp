package resp

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"slices"
	"strconv"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) readCRLFLine() (string, error) {
	line, err := s.r.ReadString('\n')
	if err != nil {
		return "", err
	}
	if len(line) < 2 || line[len(line)-2] != '\r' {
		return "", fmt.Errorf("resp: invalid line coding")
	}
	return line[:len(line)-2], nil
}

func (s *Scanner) ReadSimpleString() (string, error) {
	value, err := s.readCRLFLine()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (s *Scanner) ReadSimpleError() (string, error) {
	value, err := s.readCRLFLine()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (s *Scanner) ReadInteger() (int64, error) {
	value, err := s.readCRLFLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(value, 10, 64)
}

func (s *Scanner) ReadNull() (any, error) {
	value, err := s.readCRLFLine()
	if err != nil {
		return nil, err
	}
	if value != "" {
		return nil, fmt.Errorf("resp: invalid null")
	}
	return nil, nil
}

func (s *Scanner) ReadDouble() (float64, error) {
	value, err := s.readCRLFLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(value, 64)
}

func (s *Scanner) ReadBigNumber() (*big.Int, error) {
	value, err := s.readCRLFLine()
	if err != nil {
		return nil, err
	}
	n, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return nil, fmt.Errorf("resp: invalid bignumber")
	}
	return n, nil
}

func (s *Scanner) ReadBoolean() (bool, error) {
	value, err := s.readCRLFLine()
	if err != nil {
		return false, err
	}
	if len(value) != 1 {
		return false, fmt.Errorf("resp: invalid boolean")
	}
	switch value[0] {
	case 't':
		return true, nil
	case 'f':
		return false, nil
	default:
		return false, fmt.Errorf("resp: invalid boolean")
	}
}

func (s *Scanner) ReadBulkString() (string, error) {
	lenStr, err := s.readCRLFLine()
	if err != nil {
		return "", err
	}
	n, err := strconv.Atoi(lenStr)
	if err != nil {
		return "", err
	}
	if n == -1 {
		return "", nil
	}
	buf := make([]byte, n)
	_, err = io.ReadFull(s.r, buf)
	if err != nil {
		return "", err
	}
	if _, err := s.readCRLFLine(); err != nil {
		return "", err
	}
	return string(buf), nil
}

func (s *Scanner) ReadBulkError() (string, error) {
	lenStr, err := s.readCRLFLine()
	if err != nil {
		return "", err
	}
	n, err := strconv.Atoi(lenStr)
	if err != nil {
		return "", err
	}
	buf := make([]byte, n)
	_, err = io.ReadFull(s.r, buf)
	if err != nil {
		return "", err
	}
	if _, err := s.readCRLFLine(); err != nil {
		return "", err
	}
	return string(buf), nil
}

func (s *Scanner) ReadVerbatimString() (string, error) {
	lenStr, err := s.readCRLFLine()
	if err != nil {
		return "", err
	}
	n, err := strconv.Atoi(lenStr)
	if err != nil {
		return "", err
	}
	buf := make([]byte, n)
	_, err = io.ReadFull(s.r, buf)
	if err != nil {
		return "", err
	}
	if _, err := s.readCRLFLine(); err != nil {
		return "", err
	}
	return string(buf), nil
}

func (s *Scanner) ReadArray() (any, error) {
	var buf []any
	lenStr, err := s.readCRLFLine()
	if err != nil {
		return nil, err
	}
	n, err := strconv.Atoi(lenStr)
	if err != nil {
		return nil, err
	}
	// if array is empty
	if n == 0 {
		return make([]any, 0), nil
	}
	// if array is null
	if n == -1 {
		return nil, nil
	}
	for n != 0 {
		val, err := s.Read()
		if err != nil {
			return nil, err
		} else {
			buf = append(buf, val)
		}
		n--
	}
	return buf, nil
}

func (s *Scanner) ReadMap() (any, error) {
	var buf map[any]any = make(map[any]any)
	lenStr, err := s.readCRLFLine()
	if err != nil {
		return nil, err
	}
	n, err := strconv.Atoi(lenStr)
	if err != nil {
		return nil, err
	}
	// if map is empty
	if n == 0 {
		return buf, nil
	}
	//%<number-of-entries>\r\n<key-1><value-1>...<key-n><value-n>
	// I'm dealing with map and need to foucs on getting
	// am dealing with n * 2 values and after each 2 iteration getting single pair
	for n != 0 {
		key, err := s.Read()
		if err != nil {
			return nil, err
		}
		val, err := s.Read()
		if err != nil {
			return nil, err
		}
		buf[key] = val
		n--
	}
	return buf, nil
}

func (s *Scanner) ReadAttribute() (any, error) {
	var buf map[any]any = make(map[any]any)
	lenStr, err := s.readCRLFLine()
	if err != nil {
		return nil, err
	}
	n, err := strconv.Atoi(lenStr)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return make(map[any]any, 0), nil
	}
	for n != 0 {
		key, err := s.Read()
		if err != nil {
			return nil, err
		}
		val, err := s.Read()
		if err != nil {
			return nil, err
		}
		buf[key] = val
		n--
	}
	return buf, nil
}

func (s *Scanner) ReadSet() (any, error) {
	var buf []any
	lenStr, err := s.readCRLFLine()
	if err != nil {
		return nil, err
	}
	n, err := strconv.Atoi(lenStr)
	if err != nil {
		return nil, err
	}
	// if set is empty
	if n == 0 {
		return make([]any, 0), nil
	}
	// if set is null
	if n == -1 {
		return nil, nil
	}
	for n != 0 {
		val, err := s.Read()
		if err != nil {
			return make([]any, 0), nil
		} else if !slices.Contains(buf, val) {
			buf = append(buf, val)
		}
		n--
	}
	return buf, nil
}

func (s *Scanner) ReadPush() (any, error) {
	var buf []any
	lenStr, err := s.readCRLFLine()
	if err != nil {
		return nil, err
	}
	n, err := strconv.Atoi(lenStr)
	if err != nil {
		return nil, err
	}
	// if set is empty
	if n == 0 {
		return make([]any, 0), nil
	}

	for n != 0 {
		val, err := s.Read()
		if err != nil {
			return make([]any, 0), err
		} else {
			buf = append(buf, val)
		}
		n--
	}
	return buf, nil
}

func (s *Scanner) Read() (any, error) {
	t, err := s.r.ReadByte()

	if err != nil {
		return nil, err
	}

	switch t {

	// +OK\r\n
	case '+':
		return s.ReadSimpleString()

	//-Error message\r\n
	case '-':
		return s.ReadSimpleError()

	//:[<+|->]<value>\r\n
	case ':':
		return s.ReadInteger()

	//$<length>\r\n<data>\r\n
	case '$':
		return s.ReadBulkString()

	//*<number-of-elements>\r\n<element-1>...<element-n>
	case '*':
		return s.ReadArray()

	//_\r\n
	case '_':
		return s.ReadNull()

	//#<t|f>\r\n
	case '#':
		return s.ReadBoolean()

	//,[<+|->]<integral>[.<fractional>][<E|e>[sign]<exponent>]\r\n
	case ',':
		return s.ReadDouble()

	//([+|-]<number>\r\n
	case '(':
		return s.ReadBigNumber()

	//!<length>\r\n<error>\r\n
	case '!':
		return s.ReadBulkError()

	// =<length>\r\n<encoding>:<data>\r\n
	case '=':
		return s.ReadVerbatimString()

	//%<number-of-entries>\r\n<key-1><value-1>...<key-n><value-n>
	case '%':
		return s.ReadMap()

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
		return s.ReadAttribute()

	//~<number-of-elements>\r\n<element-1>...<element-n>
	case '~':
		return s.ReadSet()

	//><number-of-elements>\r\n<element-1>...<element-n>
	case '>':
		return s.ReadPush()
	default:
		return nil, fmt.Errorf("unsupported RESP type: %q", t)
	}
}
