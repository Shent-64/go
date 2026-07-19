package main

import "fmt"

var Global = 5

func change_Global() {
	Global = 50
	fmt.Println(Global)
	defer func() {
		Global = 5
	}()
}
