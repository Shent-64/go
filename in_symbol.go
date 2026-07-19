package main

import "fmt"

func main() {
	sentence := "Hello, World!"

	frequency := make(map[rune]int)
	for _, v := range sentence {
		frequency[v]++
	}

	for k, v := range frequency {
		fmt.Printf("Знак %c встречается %d раз \n", k, v)
	}
}
