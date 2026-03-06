package exer7

import "io"

type MyStringReader struct {
	s string
	i int
}

func NewMyStringReader(s string) *MyStringReader {
	return &MyStringReader{s, 0}
}

func (r *MyStringReader) Read(p []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(p, r.s[r.i:])
	r.i += n
	return
}
