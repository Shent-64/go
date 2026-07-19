package main

import "fmt"

func main() {
	slice := make([]int, 100, 100)
	slice1 := make([]int, 20, 100)
	for i := 1; i <= 100; i++ {
		slice[i-1] = i
	}

	fmt.Println(slice, "\n")
	slice = append(slice[:10], slice[len(slice)-10:]...)

	fmt.Println(slice, "\n")

	for i := 0; i < len(slice); i++ {
		slice1[i] = slice[len(slice)-i-1]
	}

	slice = slice1

	fmt.Println(slice)
}
