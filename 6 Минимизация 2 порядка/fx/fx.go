package fx

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
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
var h = H / 10
var Eps = 1e-6

// Производная (вычисляется)
func DF(x float64, F func(float64) float64) float64 {
	aX, bX := x-h, x+h
	dX := 2 * h
	dY := F(bX) - F(aX)
	return dY / dX
}

// Вторая производная

func DDF(x float64, F func(float64) float64) float64 {
	aX, bX := x-h, x+h
	dX := 2 * h
	dY := DF(bX, F) - DF(aX, F)
	return dY / dX
}

// Частная производная x
func DFDx(x, y float64, F func(float64, float64) float64) float64 {
	return (F(x+h, y) - F(x-h, y)) / (2 * h)
}

// Частная производная y
func DFDy(x, y float64, F func(float64, float64) float64) float64 {
	return (F(x, y+h) - F(x, y-h)) / (2 * h)
}

var X0 = 2.0
var Y0 = 2.0

func gradient(x, y float64) (float64, float64) {
	dfdx := DFDx(x, y, F)
	dfdy := DFDy(x, y, F)
	return dfdx, dfdy
}

// Матрица Гессе
func geser(x, y float64) *mat.Dense {
	m := mat.NewDense(2, 2, nil)
	m.Set(0, 0, (DFDx(x+h, y, F)-DFDx(x-h, y, F))/(2*h)) // d²f/dx²
	m.Set(0, 1, (DFDy(x+h, y, F)-DFDy(x-h, y, F))/(2*h)) // d²f/dxdy
	m.Set(1, 0, (DFDx(x, y+h, F)-DFDx(x, y-h, F))/(2*h)) // d²f/dydx
	m.Set(1, 1, (DFDy(x, y+h, F)-DFDy(x, y-h, F))/(2*h)) // d²f/dy²
	return m
}

// Метод наискорейшего спуска
func Newton(F func(float64, float64) float64) (float64, float64) {
	x, y := X0, Y0
	count := 0
	for i := 0; i < 1000; i++ {
		count++
		// Вычисляем градиент и матрицу Гессе
		gx, gy := gradient(x, y)
		h := geser(x, y)

		// Создаем вектор градиента
		grad := mat.NewVecDense(2, []float64{gx, gy})

		/// Решаем систему H * delta = -grad
		var delta mat.VecDense
		if err := delta.SolveVec(h, grad); err != nil {
			break
		}

		// Обновляем точку
		x -= delta.AtVec(0)
		y -= delta.AtVec(1)

		// Проверка условия остановки
		if math.Abs(gx) < Eps && math.Abs(gy) < Eps {
			break
		}
	}
	fmt.Println("Итераций: ", count)

	return x, y
}
