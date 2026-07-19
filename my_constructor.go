package main

import "fmt"

type Item struct {
	NoOption   string
	Parameter1 string
	Parameter2 int
}

func NewItem(opts ...option) *Item {
	i := &Item{
		NoOption:   "usual",
		Parameter1: "default",
		Parameter2: 42,
	}

	for _, opt := range opts {
		opt(i)
	}
	return i
}

type option func(*Item)

func Option1(option1 string) option {
	return func(i *Item) {
		i.Parameter1 = option1
	}
}

func Option2(option2 int) option {
	return func(i *Item) {
		i.Parameter2 = option2
	}
}

func main() {
	item1 := NewItem()
	item2 := NewItem(Option2(70))
	item3 := NewItem(Option1("unusual"), Option2(99))
	item4 := NewItem(Option2(88), Option1("rare"))

	fmt.Println(item1)
	fmt.Println(item2)
	fmt.Println(item3)
	fmt.Println(item4)
}
