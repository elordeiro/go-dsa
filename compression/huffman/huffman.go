package huffman

import (
	"github.com/elordeiro/goext/containers/pq"
)

type Huffman struct {
	root   *node
	freq   []int
	codes  []code
	output []byte
}

type node struct {
	freq        int
	char        byte
	left, right *node
}

type code struct {
	code   int
	length int
}

func Compress(input []byte) *Huffman {
	h := &Huffman{}
	h.freq = make([]int, 256)
	for _, b := range input {
		h.freq[b]++
	}
	h.buildTree()
	h.buildCodes()
	h.compress(input)
	return h
}

func (h *Huffman) buildTree() {
	var nodes []*node
	for i, f := range h.freq {
		if f > 0 {
			nodes = append(nodes, &node{freq: f, char: byte(i)})
		}
	}
	pq := pq.NewPQFunc(func(n1, n2 *node) bool {
		return n1.freq < n2.freq
	}, nodes...)

	for pq.Len() > 1 {
		n1 := pq.Pop()
		n2 := pq.Pop()
		n := &node{freq: n1.freq + n2.freq, left: n1, right: n2}
		pq.Push(n)
	}

	if pq.Len() == 0 {
		return
	}
	h.root = pq.Pop()
}

func (h *Huffman) buildCodes() {
	if h.root == nil {
		return
	}
	h.codes = make([]code, 256)
	var build func(*node, int, int)
	build = func(n *node, c, l int) {
		if n.left == nil && n.right == nil {
			h.codes[n.char] = code{c, l}
			return
		}
		build(n.left, c<<1, l+1)
		build(n.right, c<<1|1, l+1)
	}
	build(h.root, 0, 0)
}

func (h *Huffman) compress(input []byte) {
	h.output = []byte{}
	var b, n int
	for _, c := range input {
		b = (b << h.codes[c].length) | h.codes[c].code
		n += h.codes[c].length
		for n > 8 {
			h.output = append(h.output, byte(b>>(n-8)))
			n -= 8
		}
	}
	if n > 0 {
		h.output = append(h.output, byte(b<<(8-n)))
	}
}

func (h *Huffman) Output() []byte {
	return h.output
}

func (h *Huffman) Codes() []code {
	return h.codes
}

func (h *Huffman) Freq() []int {
	return h.freq
}

func (h *Huffman) Root() *node {
	return h.root
}
