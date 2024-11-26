package huffman_test

import (
	"log"
	"os"
	"testing"

	"github.com/elordeiro/goext/compression/huffman"
)

func BenchmarkCompress1(b *testing.B) {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		huffman.Compress(file)
	}
}
