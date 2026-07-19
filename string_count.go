package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(cnt *int){
	*cnt++;
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Iteraction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}
