package tihonov

import (
	"fmt"
	"math"
)

// Функция для вывода матрицы на экран
func PrintMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("[%.2f]", col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintVector(v []float64) {
	for _, cell := range v {
		fmt.Printf("[%.2f]\n", cell)
	}
	fmt.Println()
}

func MxM(a, b [][]float64) [][]float64 {
	c := make([][]float64, len(a))
	for i := range a {
		c[i] = make([]float64, len(a[0]))
		for j := range a[i] {
			for k := range a[i] {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func Vnorma(a []float64) float64 {
	c := 0.0
	for i := range a {
		c += a[i] * a[i]
	}
	return math.Sqrt(c)
}

func MpM(a, b [][]float64) [][]float64 {
	c := a
	for i := range a {
		for j := range a[i] {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func MTranspose(a [][]float64) [][]float64 {
	c := make([][]float64, len(a[0]))
	for i := range a {
		c[i] = make([]float64, len(a))
		for j := range a[i] {
			c[i][j] = a[j][i]
		}
	}
	return c
}

func MxV(a [][]float64, b []float64) []float64 {
	n := len(b)
	c := make([]float64, len(b))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i] += a[i][j] * b[j]
		}
	}
	return c
}
