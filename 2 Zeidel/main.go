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

func zeidel(matrix [][]float64, b []float64, eps float64) []float64 {
	n := len(b)
	// Изначальный массив решений (заполнен 0)
	solution := make([]float64, n)
	// Новое приближение
	neo := make([]float64, n)
	count := 0
	for {
		// Считаем
		count++
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < i; j++ {
				sum += matrix[i][j] * neo[j]
			}

			for j := i + 1; j < n; j++ {
				sum += matrix[i][j] * solution[j]
			}
			neo[i] = (b[i] - sum) / matrix[i][i]
		}

		// Проверяем эпсилон
		norm := 0.0
		for x := range neo {
			norm += math.Abs(neo[x] - solution[x])
		}

		fmt.Println(norm)

		if norm <= eps {
			break
		}

		solution = neo
	}
	fmt.Println("count: ", count)
	return solution
}

func main() {
	// Расширенная матрица системы
	matrix := [][]float64{
		{3, 0, -1},
		{2, -5, 1},
		{200, -2, 6},
	}

	b := []float64{-4, 9, 8}

	eps := 1e-2

	fmt.Println("Исходная матрица:")
	printMatrix(matrix)
	fmt.Println("Правые части:")
	printVector(b)

	solution := zeidel(matrix, b, eps)

	// Выводим все
	fmt.Println("Решение системы:")
	for i, x := range solution {
		fmt.Printf("x%d = %f\n", i+1, x)
	}

}
