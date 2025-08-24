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

// Решение
func GaussJordan(matrix [][]float64) []float64 {
	rows := len(matrix)
	cols := len(matrix[0])

	// Прямой ход
	for row := 0; row < rows; row++ {
		// Выбор элемента
		// (найти максимальный элемент в столбце (ниже текущей строки) и поменять эту строку с текущей)
		col := row
		maxRow := row
		for r := row + 1; r < rows; r++ {
			if math.Abs(matrix[r][col]) > math.Abs(matrix[maxRow][col]) {
				maxRow = r
			}
		}

		// Меняем текущую строку и строку с ведущим элементом
		matrix[row], matrix[maxRow] = matrix[maxRow], matrix[row]

		// Нормализация текущей строки по ведущему элементу
		norm := matrix[row][col]
		if norm == 0 {
			panic("SLAU has inf or 0 solutions")
		}
		for c := col; c < cols; c++ {
			matrix[row][c] /= norm
		}
		//printMatrix(matrix)

		// Вычитание текущей строки из нижних строк
		for r := row + 1; r < rows; r++ {
			// число на которое домножаем перед вычитанием
			multi := matrix[r][col]
			for c := 0; c < cols; c++ {
				// первый элемент должен обнулиться
				matrix[r][c] -= multi * matrix[row][c]
			}
		}
		//printMatrix(matrix)
	}
	//printMatrix(matrix)

	// Обратный ход
	// Массив решений
	solution := make([]float64, cols-1)

	// идём с конца
	for i := rows - 1; i >= 0; i-- {
		// res = b
		res := matrix[i][cols-1]
		for k := i + 1; k < cols-1; k++ {
			// Вычитаем a*x (там где x не вычислен будет 0, так и задумано)
			res -= solution[k] * matrix[i][k]
		}
		solution[i] = res
	}
	return solution
}

func main() {
	// Расширенная матрица системы
	matrix := [][]float64{
		{3, 0, -1, -4},
		{2, -5, 1, 9},
		{200, -2, 6, 8},
	}

	fmt.Println("Исходная матрица:")
	printMatrix(matrix)

	// Получаем массив решений
	solution := GaussJordan(matrix)

	// Выводим все
	fmt.Println("Решение системы:")
	for i, x := range solution {
		fmt.Printf("x%d = %f\n", i+1, x)
	}
}
