package fx

import (
	"fmt"
	"math"
)

// Функция, которую мы хотим построить
func F(x float64) float64 {

	x = math.Abs(x)
	if x < 0.5 {
		return 1
	} else if x > 0.5 {
		return 0
	} else {
		return 0.5
	}

	//return math.Cos(x)
	//return math.Tan(x)
	//return x + 1

}

//////////////////////////////////////////

// Интервал
var A = -math.Pi/2.0 + 0.1
var B = math.Pi/2.0 - 0.1

// Шаг
var H = 0.001

// Точность
var Eps = 1e-12

// Период
var T = 2 * math.Pi

// Степень
var M = 10.0

// 2*pi / T
var W = 1.0

//////////////////////////////////////////

func Wre() {
	W = 2.0 * math.Pi / T
}

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

// для одного интервала
func simpson(a, b float64, f func(float64, float64) float64, k float64) float64 {
	center := (a + b) / 2.0
	h := b - a
	return (h / 6.0) * (f(a, k) + 4*f(center, k) + f(b, k))
}

func simpson1(a, b float64, f func(float64) float64) float64 {
	center := (a + b) / 2.0
	h := b - a
	return (h / 6.0) * (f(a) + 4*f(center) + f(b))
}

// адаптивный (автоматическое деление на неравные интервалы)
func adapt_simpson(a, b float64, f func(float64, float64) float64, k float64) float64 {
	// находим интеграл на всём интервале
	J := simpson(a, b, f, k)
	// а потом для каждой половины
	center := (a + b) / 2
	J1 := simpson(a, center, f, k)
	J2 := simpson(center, b, f, k)
	Jsum := J1 + J2
	// если соответствует точности
	if math.Abs(Jsum-J)/15 <= Eps {
		return Jsum
	}
	// иначе повторяем рекурсивно
	return adapt_simpson(a, center, f, k) + adapt_simpson(center, b, f, k)
}

func adapt_simpson1(a, b float64, f func(float64) float64) float64 {
	// находим интеграл на всём интервале
	J := simpson1(a, b, f)
	// а потом для каждой половины
	center := (a + b) / 2
	J1 := simpson1(a, center, f)
	J2 := simpson1(center, b, f)
	Jsum := J1 + J2
	// если соответствует точности
	if math.Abs(Jsum-J)/15.0 <= Eps {
		return Jsum
	}
	// иначе повторяем рекурсивно
	return adapt_simpson1(a, center, f) + adapt_simpson1(center, b, f)
}

func FxSin(x, k float64) float64 {
	return F(x) * math.Sin(x*W*k)
}

func FxCos(x, k float64) float64 {
	return F(x) * math.Cos(x*W*k)
}

// Вычисляем a0 заранее
var a0 = 2 / T * adapt_simpson1(0, T, F)

func Furje(x float64) float64 {
	var ak, bk float64

	// Коэффициенты

	result := a0 / 2.0

	for k := 1.0; k <= M; k++ {
		// чётная
		if chet {
			bk = 0.0
		} else {
			bk = 2.0 / T * adapt_simpson(0, T, FxSin, k)
		}
		// нечётная
		if nechet {
			ak = 0.0
		} else {
			ak = 2.0 / T * adapt_simpson(0, T, FxCos, k)
		}

		result += (ak*math.Cos(k*W*x) + bk*math.Sin(k*W*x))
	}

	return result
}

var chet, nechet = false, false

func ChetOrNot() {
	chet = Chet()
	nechet = Nechet()
	if chet {
		fmt.Print("\nФункция чётная\n")
	}
	if nechet {
		fmt.Print("\nФункция нечётная\n")
	}
	if !chet && !nechet {
		fmt.Print("\nФункция не чётная и не нечётная\n")
	}
}

// Проверка на чётность
func Chet() bool {
	result := true
	for k := A; k <= B; k += H {
		if F(k) != F(-k) {
			result = false
		}
	}
	return result
}

// Проверка на нечётность
func Nechet() bool {
	result := true
	for k := A; k <= B; k += H {
		if F(k) != -F(-k) {
			result = false
		}
	}
	return result
}
