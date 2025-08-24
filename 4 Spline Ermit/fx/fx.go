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
	Fstart := F(start)
	Fend := F(end)

	// Замена (приведение к значению 0-1)
	t := (x - start) / H
	t2 := t * t
	t3 := t2 * t

	// Среднее значение I производной на интервале (относительно t - для этого ещё домножаем на H)
	//Dx := (Fend - Fstart) / (end - start) / 2
	//Dx *= H

	// Кубический полином (из 4 полиномов)
	return (2*t3-3*t2+1)*Fstart + (t3-2*t2+t)*H*DF(start, F)*H + (-2*t3+3*t2)*Fend + (t3-t2)*H*DF(end, F)*H
}
