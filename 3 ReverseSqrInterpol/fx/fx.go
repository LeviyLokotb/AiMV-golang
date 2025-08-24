package fx

import (
	"math"
)

// Функция, которую мы хотим построить
func F(x float64) float64 {
	return x*x - math.Cos(x)
	//return math.Exp(x) + math.Exp(-3*x) - 4
	//return math.Sin(x)
}

var A = -2.0
var B = 2.0
var Eps = 1e-4

// Производная (вычисляется)
func DF(x float64) float64 {
	aX, bX := x-0.01, x+0.01
	dX := 0.02
	dY := F(bX) - F(aX)
	return dY / dX
}

var X0 = 1.0
