package main

import (
	"fmt"
	"math"
	"newton-rafson/fu"
)

/*
	func printM(m [][]float64) {
		for i := range m {
			for j := range m[i] {
				fmt.Printf("[%f]", m[i][j])
			}
			fmt.Println()
		}
	}
*/
func Jacobi(x, y float64) [][]float64 {
	return [][]float64{
		{fu.Dx1(x), fu.Dy1(y)},
		{fu.Dx2(x), fu.Dy2(y)},
	}
}

// Определитель 2х2
func det(m [][]float64) float64 {
	return m[1][1]*m[0][0] - m[1][0]*m[0][1]
}

// Обратная матрица 2х2
func AntiMx(m [][]float64) [][]float64 {
	det := det(m)
	return [][]float64{
		{-m[1][1] / det, m[1][0] / det},
		{m[0][1] / det, -m[0][0] / det},
	}
}

func MxV(m [][]float64, v []float64) []float64 {
	result := make([]float64, len(m))
	for i := range m {
		sum := 0.0
		for j := range m[0] {
			sum += m[i][j] * v[j]
		}
		result[i] = sum
	}
	return result
}

func MxM(a, b [][]float64) [][]float64 {
	result := [][]float64{
		{0, 0},
		{0, 0},
	}
	for i := range a {
		for j := range b[0] {
			for k := range a[0] {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

var count int = 0

func NewtonRafson(x0, y0, eps float64) (float64, float64) {
	// приближение
	x, y := x0, y0
	for {
		count++
		fx := fu.F1(x, y)
		fy := fu.F2(x, y)

		J := Jacobi(x, y)
		det := det(J)
		if det == 0 {
			panic("0 определитель")
		}
		/*
			fmt.Println("Jacobi:")
			printM(J)
			fmt.Println("-Jacobi:")
			printM(AntiMx(J))

			fmt.Println("Test J * J^-1:")
			printM(MxM(J, AntiMx(J)))
		*/

		// Получаем СЛАУ J * [dx, dy] = - [fx, fy]
		// [dx, dy] = J^-1 * F
		/*
			delta := MxV(AntiMx(J), []float64{fx, fy})
			dx, dy := delta[0], delta[1]
		*/
		dx := (-fx*J[1][1] + fy*J[0][1]) / det
		dy := (fx*J[1][0] - fy*J[0][0]) / det
		x += dx
		y += dy
		//fmt.Println("x = ", x, "\ny = ", y)

		// проверяем епсилон
		if max(math.Abs(fx), math.Abs(fy)) <= eps {
			break
		}

	}
	return x, y
}

func main() {
	eps := 1e-4
	x0, y0 := 1.0, 1.0

	fmt.Println("====================\nМетод Ньютона-Рафсона\n====================")
	fmt.Println("eps = ", eps)
	fmt.Println("x0 = ", x0)
	fmt.Println("y0 = ", y0)

	x, y := NewtonRafson(x0, y0, eps)

	fmt.Println("\nРезультат: \nx = ", x, "\ny = ", y)
	fmt.Println("Достигнуто за ", count, " итераций.")
	fmt.Println(fu.F1(x, y))
	fmt.Println(fu.F2(x, y))
}
