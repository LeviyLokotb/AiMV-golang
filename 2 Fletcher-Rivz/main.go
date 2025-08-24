package main

import (
	"fmt"
	"math"
)

// Функция для вывода матрицы на экран
func printMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("[%.2f]", col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func printVector(v []float64) {
	for _, cell := range v {
		fmt.Printf("[%.2f]\n", cell)
	}
	fmt.Println()
}

// Произведение матрицы на вектор
func Multipue(a [][]float64, b []float64) []float64 {
	n := len(b)
	c := make([]float64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c[i] += a[i][j] * b[j]
		}
	}
	return c
}

// Разность векторов
func Diff(a, b []float64) []float64 {
	c := make([]float64, len(a))
	for i := range c {
		c[i] = a[i] - b[i]
	}
	return c
}

// Скалярное произведение векторов
func ScalarMulti(a, b []float64) float64 {
	result := 0.0
	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}
	return result
}

func FletcherRivz(a [][]float64, b []float64, eps float64) []float64 {
	n := len(b)
	x := make([]float64, n)
	d := make([]float64, n)

	// градиент = A * x0 - b
	grad := Multipue(a, x)
	grad = Diff(grad, b)
	// направление поиска (антиградиент)
	for i := range d {
		d[i] = grad[i]
	}

	// Итерации:
	for k := 0; k < 100; k++ {
		// Ищем коэффициент альфа
		Ad := Multipue(a, d)
		alpha := ScalarMulti(grad, grad) / ScalarMulti(d, Ad)
		fmt.Println("alpha = ", alpha)

		// Обновляем решение
		for i := range x {
			x[i] += alpha * d[i]
		}
		fmt.Println("x:")
		printVector(x)

		// Новый градиент (вычисляется так же, но с новым x)
		gradNew := make([]float64, n)
		gradNew = Multipue(a, x)
		gradNew = Diff(gradNew, b)
		fmt.Println("gradient:")
		printVector(gradNew)

		// Проверяем эпсилон
		norm := math.Sqrt(ScalarMulti(gradNew, gradNew))
		fmt.Println(norm)
		if norm <= eps {
			break
		}

		// Ищем коэффициент бета
		beta := ScalarMulti(gradNew, gradNew) / ScalarMulti(grad, grad)
		fmt.Println("beta = ", beta)

		// Обновляем направление поиска
		for i := 0; i < n; i++ {
			d[i] = gradNew[i] + beta*d[i]
		}

		// Обновляем градиент
		grad = gradNew
	}
	return x
}

func main() {
	// Расширенная матрица системы
	matrix := [][]float64{
		{3, 0, -1},
		{2, -5, 1},
		{2, -2, 6},
	}
	b := []float64{-4, 9, 8}
	/*
			{3, 0, -1},
			{2, -5, 1},
			{2, -2, 6},
		}
		b := []float64{-4, 9, 8}
	*/
	/*
			{4, 1, 1},
			{1, 3, 1},
			{1, 1, 5},
		}
		b := []float64{7, 5, 10}
	*/

	eps := 1e-1

	fmt.Println("Исходная матрица:")
	printMatrix(matrix)
	fmt.Println("Правые части:")
	printVector(b)

	solution := FletcherRivz(matrix, b, eps)

	// Выводим все
	fmt.Println("Решение системы:")
	for i, x := range solution {
		fmt.Printf("x%d = %f\n", i+1, x)
	}
}
