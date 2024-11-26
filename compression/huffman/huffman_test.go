package huffman_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/compression/huffman"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		input  []byte
		output []byte
	}{
		{
			input:  []byte(""),
			output: []byte{},
		},
		{
			input:  []byte("aaabbc"),
			output: []byte{0x1F, 0x00},
		},
		{
			input:  []byte("122333444455555666666"),
			output: []byte{0xAB, 0xB9, 0x20, 0x02, 0xAB, 0xFF, 0xE0},
		},
	}

	for _, test := range tests {
		h := huffman.Compress(test.input)
		if got, want := h.Output(), test.output; !slices.Equal(got, want) {
			t.Errorf("Compress(%q) = %v, want %v", test.input, got, want)
		}
	}
}
