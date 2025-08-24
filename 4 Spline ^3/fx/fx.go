package fx

import (
	"math"
)

// Функция, которую мы хотим построить
func F1(x float64) float64 {
	return math.Exp(x)
}

func F2(x float64) float64 {
	return math.Exp(-x)
}

func F3(x float64) float64 {
	return math.Sinh(x)
}

func F4(x float64) float64 {
	return math.Cosh(x)
}

func F5(x float64) float64 {
	return math.Sin(x)
}

func F6(x float64) float64 {
	return math.Cos(x)
}

func F7(x float64) float64 {
	return math.Log(x)
}

var A = 1.0
var B = 1.2
var H = 0.04
var Eps = 1e-4

// Производная (вычисляется)
func DF(x float64, F func(float64) float64) float64 {
	h := H / 10
	aX, bX := x-h, x+h
	dX := 2 * h
	dY := F(bX) - F(aX)
	return dY / dX
}

// Вторая производная
func DDFv2(x float64, F func(float64) float64) float64 {
	h := H / 10
	aY, bY := DF(x-h, F), DF(x+h, F)
	dX := 2 * h
	dY := bY - aY
	return dY / dX
}

func DDF(x float64, F func(float64) float64) float64 {
	h := H / 10
	aX, bX := x-h, x+h
	dX := 2 * h
	dY := DF(bX, F) - DF(aX, F)
	return dY / dX
}

//var X0 = 1.0

func CubeSpline(x float64, F func(float64) float64) float64 {
	// приводим x к кратному 0.04
	start := x //float64(int(x*100/4)) * 0.04
	//x += H
	end := x + H
	A := (end - x)
	B := (x - start)
	FA := F(start)
	FB := F(end)
	//DFA := fx.DF(start, F)
	DDFA := DDF(start, F)
	//DFB := fx.DF(end, F)
	DDFB := DDF(end, F)
	A /= H
	B /= H
	return A*FA + B*FB + ((A*A*A-A)*DDFA+(B*B*B-B)*DDFB)*H*H/6.0
}
