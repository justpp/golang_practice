package util

import (
	"testing"
)

func TestPrintFo(t *testing.T) {
	var tests = []struct {
		name string
	}{
		{name: "first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintFo()
		})
	}
}

func BenchmarkPrintFo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintFo()
	}
}
