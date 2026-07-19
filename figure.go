package main

type figures int

const (
	square figures = iota
	circle
	triangle
)

func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(a float64) float64 {
			return a * a
		}, true
	case circle:
		return func(r float64) float64 {
			return 3.14 * r * r
		}, true
	case triangle:
		return func(a float64) float64 {
			return (a * a) / 2
		}, true
	default:
		return nil, false
	}
}

func main() {
	ar, ok := area(circle)
	if !ok {
		println("Ошибка")
		return
	}
	x := 5.0
	myArea := ar(x)
	println(myArea)
}
