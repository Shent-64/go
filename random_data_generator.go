package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type generator struct {
	rnd rand.Source
}

func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

func (g *generator) Read(bytes []byte) (n int, err error) {
	for i := range bytes {
		randInt := g.rnd.Int63()
		randByte := byte(randInt)
		bytes[i] = randByte
	}
	return len(bytes), nil
}

func main() {
	generator := New(time.Now().UnixNano())
	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf)
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

}
