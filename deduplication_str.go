package main

import "fmt"

func RemoveDuplicates(input []string) []string {
	m := make(map[string]bool)

	var result []string
	for _, item := range input {
		if _, ok := m[item]; !ok {
			m[item] = true
			result = append(result, item)
		}
	}

	return result
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	result := RemoveDuplicates(input)
	for _, item := range result {
		fmt.Println(item)
	}
}
