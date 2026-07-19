package main

import (
	"fmt"

	mypkg "github.com/Shent-64/go"
)

func main() {
	if sum := mypkg.Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
