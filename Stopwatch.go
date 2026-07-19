package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start  time.Time
	splits []time.Duration
}

func (s *Stopwatch) Start() {
	s.start = time.Now()
	s.splits = nil
}

func (s *Stopwatch) SaveSplit() {
	s.splits = append(s.splits, time.Since(s.start))
}

func (s *Stopwatch) GetSplits() []time.Duration {
	return s.splits
}

func main() {
	sw := Stopwatch{}
	sw.Start()
	time.Sleep(1 * time.Second)
	sw.SaveSplit()
	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetSplits())
}
