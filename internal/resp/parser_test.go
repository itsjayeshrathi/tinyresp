package resp

import (
	"strings"
	"testing"
)

func TestRead_SimpleString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "correct string",
			input:   "+abc\r\n",
			wantErr: false,
		},
		{
			name:    "wrong format",
			input:   "+abc\n",
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "+\r\n",
			wantErr: false,
		},
		{
			name:    "number",
			input:   "+1000\r\n",
			wantErr: false,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(tt.input))
			val, err := scanner.Read()

			if !tt.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error, got value: %#v", val)
			}
		})
	}
}

func TestRead_SimpleError(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "non numeric",
			input:   ":abc\r\n",
			wantErr: true,
		},
		{
			name:    "empty integer",
			input:   ":\r\n",
			wantErr: true,
		},
		{
			name:    "partially numeric",
			input:   ":12x\r\n",
			wantErr: true,
		},
		{
			name:    "only sign",
			input:   ":+\r\n",
			wantErr: true,
		},
		{
			name:    "number",
			input:   ":+1000\r\n",
			wantErr: false,
		}, {
			name:    "float",
			input:   ":1999.22\r\n",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(tt.input))
			val, err := scanner.Read()

			if !tt.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expeced error, got value: %v", val)
			}
		})
	}
}

func TestRead_SimpleInteger(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "non numeric",
			input:   ":abc\r\n",
			wantErr: true,
		},
		{
			name:    "empty integer",
			input:   ":\r\n",
			wantErr: true,
		},
		{
			name:    "partially numeric",
			input:   ":12x\r\n",
			wantErr: true,
		},
		{
			name:    "only sign",
			input:   ":+\r\n",
			wantErr: true,
		},
		{
			name:    "number",
			input:   ":+1000\r\n",
			wantErr: false,
		}, {
			name:    "float",
			input:   ":1999.22\r\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(tt.input))
			val, err := scanner.Read()

			if !tt.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error, got value: %#v", val)
			}
		})
	}
}

func TestRead_Null(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "non numeric",
			input:   ":abc\r\n",
			wantErr: true,
		},
		{
			name:    "empty integer",
			input:   ":\r\n",
			wantErr: true,
		},
		{
			name:    "partially numeric",
			input:   ":12x\r\n",
			wantErr: true,
		},
		{
			name:    "only sign",
			input:   ":+\r\n",
			wantErr: true,
		},
		{
			name:    "number",
			input:   ":+1000\r\n",
			wantErr: false,
		}, {
			name:    "float",
			input:   ":1999.22\r\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(tt.input))
			val, err := scanner.Read()

			if !tt.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error, got value: %#v", val)
			}
		})
	}
}

func TestRead_Double(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "non numeric",
			input:   ":abc\r\n",
			wantErr: true,
		},
		{
			name:    "empty integer",
			input:   ":\r\n",
			wantErr: true,
		},
		{
			name:    "partially numeric",
			input:   ":12x\r\n",
			wantErr: true,
		},
		{
			name:    "only sign",
			input:   ":+\r\n",
			wantErr: true,
		},
		{
			name:    "number",
			input:   ":+1000\r\n",
			wantErr: false,
		}, {
			name:    "float",
			input:   ":1999.22\r\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(tt.input))
			val, err := scanner.Read()

			if !tt.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error, got value: %#v", val)
			}
		})
	}
}

func TestRead_BigNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "non numeric",
			input:   ":abc\r\n",
			wantErr: true,
		},
		{
			name:    "empty integer",
			input:   ":\r\n",
			wantErr: true,
		},
		{
			name:    "partially numeric",
			input:   ":12x\r\n",
			wantErr: true,
		},
		{
			name:    "only sign",
			input:   ":+\r\n",
			wantErr: true,
		},
		{
			name:    "number",
			input:   ":+1000\r\n",
			wantErr: false,
		}, {
			name:    "float",
			input:   ":1999.22\r\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(tt.input))
			val, err := scanner.Read()

			if !tt.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error, got value: %#v", val)
			}
		})
	}
}

func TestRead_Boolean(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "non numeric",
			input:   ":abc\r\n",
			wantErr: true,
		},
		{
			name:    "empty integer",
			input:   ":\r\n",
			wantErr: true,
		},
		{
			name:    "partially numeric",
			input:   ":12x\r\n",
			wantErr: true,
		},
		{
			name:    "only sign",
			input:   ":+\r\n",
			wantErr: true,
		},
		{
			name:    "number",
			input:   ":+1000\r\n",
			wantErr: false,
		}, {
			name:    "float",
			input:   ":1999.22\r\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(tt.input))
			val, err := scanner.Read()

			if !tt.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error, got value: %#v", val)
			}
		})
	}
}
