package fu

import (
	"math"
)

// Функции, описывающие СНАУ:
func F1(x, y float64) float64 {
	return math.Sin(y) + 2*x - 2
}

func F2(x, y float64) float64 {
	return math.Cos(x-1) + y - 0.7
}

// Производная (считаем вручную, к сожалению)
func Dx1(x float64) float64 {
	return 2
}
func Dx2(x float64) float64 {
	return -math.Sin(x - 1)
}
func Dy1(y float64) float64 {
	return math.Cos(y)
}
func Dy2(y float64) float64 {
	return 1
}
