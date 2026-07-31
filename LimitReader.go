package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type limitReader struct {
	r io.Reader
	n int
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if l.n < len(p) {
		p = p[:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= n
	return
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &limitReader{r: r, n: n}
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
