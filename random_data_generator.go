package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Hasher interface {
	io.Writer
	Hash() byte
}

type hash struct {
	result byte
}

func hashbyteNew(_init byte) Hasher {
	return &hash{
		result: _init,
	}
}

func (h *hash) Write(bytes []byte) (n int, err error) {
	for _, b := range bytes {
		h.result = (h.result^b)<<1 + b%2
	}
	return len(bytes), nil
}

func (h hash) Hash() byte {
	return h.result
}

type generator struct {
	rnd rand.Source
}

func runbiteNew(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

func (g *generator) Read(bytes []byte) (n int, err error) {
	for i := 0; i < len(bytes); i += 8 {
		randInt := g.rnd.Int63()
		binary.LittleEndian.PutUint64(bytes[i:i+8], uint64(randInt))
	}
	return len(bytes), nil
}

func main() {
	generator := runbiteNew(time.Now().UnixNano())
	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf)
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

	hasher := hashbyteNew(0)
	hasher.Write(buf)
	fmt.Printf("Hash: %v \n", hasher.Hash())

}
