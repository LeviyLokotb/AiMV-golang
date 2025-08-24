package fx

import (
	"fmt"
	"math"
)

// Функция, которую мы хотим построить
func F(x, y float64) float64 {
	/*
		xx := x * x
		Imxx := (1 - x)
		ymxx := y - xx
		return Imxx*Imxx + 100*ymxx*ymxx
	*/
	//return x*(x+1) + y*(y-1)
	return 0.26*(x*x+y*y) - 0.48*x*y
	/*
		a := 1 + math.Pow(x+y+1, 2)*(19-14*x+3*x*x-14*y+6*x*y+3*y*y)
		b := 30 + math.Pow(2*x-3*y, 2)*(18-32*x+12*x*x+48*y-36*x*y+27*y*y)
		return a * b
	*/

}

var A = 1.0
var B = 1.2
var H = 0.04
var Eps = 1e-6

// Производная (вычисляется)
func DF(x float64, F func(float64) float64) float64 {
	h := H / 10
	aX, bX := x-h, x+h
	dX := 2 * h
	dY := F(bX) - F(aX)
	return dY / dX
}

// Вторая производная

func DDF(x float64, F func(float64) float64) float64 {
	h := H / 10
	aX, bX := x-h, x+h
	dX := 2 * h
	dY := DF(bX, F) - DF(aX, F)
	return dY / dX
}

// Частная производная x
func DFDx(x, y float64, F func(float64, float64) float64) float64 {
	h := H / 10
	return (F(x+h, y) - F(x-h, y)) / (2 * h)
}

// Частная производная y
func DFDy(x, y float64, F func(float64, float64) float64) float64 {
	h := H / 10
	return (F(x, y+h) - F(x, y-h)) / (2 * h)
}

var X0 = 2.0
var Y0 = 2.0

func gradient(x, y float64) (float64, float64) {
	dfdx := DFDx(x, y, F)
	dfdy := DFDy(x, y, F)
	return dfdx, dfdy
}

// Метод наискорейшего спуска
func SteepestDescent(F func(float64, float64) float64) ([]float64, []float64) {
	var xHistory, yHistory []float64
	x, y := X0, Y0
	xHistory = append(xHistory, x)
	yHistory = append(yHistory, y)
	count := 0
	for i := 0; i < 1000; i++ {
		count++
		// Вычисляем градиент
		gx, gy := gradient(x, y)

		// Функция для одномерной оптимизации вдоль направления антиградиента
		lineSearch := func(alpha float64) float64 {
			return F(x-alpha*gx, y-alpha*gy)
		}

		// Метод дихотомии для поиска оптимального шага
		alpha := goldenSectionSearch(lineSearch, 0, 2)

		// Обновляем координаты
		x -= alpha * gx
		y -= alpha * gy

		// Сохраняем историю
		xHistory = append(xHistory, x)
		yHistory = append(yHistory, y)

		// Проверка условия остановки
		if math.Sqrt(gx*gx+gy*gy) < Eps {
			break
		}
		fmt.Println("Итераций: ", count)
	}

	return xHistory, yHistory
}

// Золотое сечение для поиска шага
func goldenSectionSearch(f func(float64) float64, a, b float64) float64 {
	phi := (1 + math.Sqrt(5)) / 2
	for math.Abs(b-a) > 1e-6 {
		x1 := b - (b-a)/phi
		x2 := a + (b-a)/phi
		if f(x1) < f(x2) {
			b = x2
		} else {
			a = x1
		}
	}
	return (a + b) / 2
}
