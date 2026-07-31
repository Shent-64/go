package main

import (
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	switch v := a.(type) {
	case int:
		return v * b
	case string:
		return strings.Repeat(v, b)
	default:
		return "Неизвестный тип"
	}
}
