package main

import (
	"fmt"
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
func Jacobi(fs []float64) [][]float64 {
	var result [][]float64
	for i := 0; i < 6; i++ {
		var neo []float64
		for j := 0; j < 6; j++ {
			neo = append(neo, fu.Dy(y, i)/fu.Dx(x, j))
		}
		result = append(result, neo)
	}

	return result
}

// Матрица исключая строку i и столбец j
func Am(m [][]float64, i, j int) [][]float64 {
	var result [][]float64
	for r := range m {
		if r == i {
			continue
		}
		var neo []float64
		for c := range m[r] {
			if c == j {
				continue
			}
			neo = append(neo, m[r][c])
		}
		result = append(result, neo)
	}
	return result
}

// Определитель по 1 столбцу
func Det(m [][]float64) float64 {
	var det float64
	for i := 0; i < 6; i++ {
		neo := Am(m, 1, i)
		det += m[i][1] * Det(neo)
	}
	return det
}

// Обратная матрица 2х2
func AntiMx(m [][]float64) [][]float64 {
	det := Det(m)
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

func NewtonRafson(x0s, eps float64) (float64, float64) {
	// приближение
	fs := x0s
	for {
		count++

		// Записываем массив приближений

		J := Jacobi(x, y)
		det := Det(J)
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

		delta := MxV(AntiMx(J), fs)

		for i := range fs {
			fs[i] += delta[i]
		}

		//fmt.Println("x = ", x, "\ny = ", y)

		// проверяем епсилон

	}
	return x, y
}

func main() {
	eps := 1e-4
	x0s := []float64{1.0, 1.0, 1.0, 1.0, 1.0, 1.0}

	fmt.Println("====================\nМетод Ньютона-Рафсона\n====================")
	fmt.Println("eps = ", eps)
	fmt.Println("x0 = ", x0)
	fmt.Println("y0 = ", y0)

	x, y := NewtonRafson(x0s, eps)

	fmt.Println("\nРезультат: \nx = ", x, "\ny = ", y)
	fmt.Println("Достигнуто за ", count, " итераций.")

}
