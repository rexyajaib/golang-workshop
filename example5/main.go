package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	// 1) Source: stream of bytes (io.Reader)
	src := strings.NewReader("hello\ngo\npipeline\n")

	// 2) Transform: streaming, no full buffering required
	upper := &toUpperReader{r: src}

	// 3) Sink: io.Writer
	io.Copy(os.Stdout, upper)
}

// toUpperReader is a streaming transformer.
// It processes bytes as they flow, not as objects.
type toUpperReader struct {
	r io.Reader
}

func (t *toUpperReader) Read(p []byte) (int, error) {
	n, err := t.r.Read(p)
	for i := 0; i < n; i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] -= 'a' - 'A'
		}
	}
	return n, err
}
