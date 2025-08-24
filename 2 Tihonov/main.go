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

// Решение
func GaussJordan(a [][]float64, b []float64) []float64 {
	var matrix [][]float64
	for i := 0; i < len(a); i++ {
		neo := a[i]
		neo = append(neo, b[i])
		matrix = append(matrix, neo)
	}
	//printMatrix(matrix)

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

func tihonov(a [][]float64, b []float64, alpha float64) []float64 {
	// поворачиваем A
	AT := MTranspose(a)
	// произведение матриц (новая A)
	ATxA := MxM(AT, a)
	// Доавляем альфа (только для диагонали, якобы умножили на единичеую иатрицу)
	for i := range ATxA {
		ATxA[i][i] += alpha
	}
	// Новый B
	ATxB := MxV(AT, b)
	return GaussJordan(ATxA, ATxB)
}

func main() {
	// Расширенная матрица системы
	matrix := [][]float64{
		{3, 0, -1},
		{2, -5, 1},
		{2, -2, 6},
	}

	b := []float64{-4, 9, 8}

	alpha := 0.1

	fmt.Println("Исходная матрица:")
	printMatrix(matrix)
	fmt.Println("Правые части:")
	printVector(b)

	// Параметр регуляризации
	Bnorma := Vnorma(b)
	fmt.Println("Границы параметра регуляризации:")
	fmt.Println(Bnorma*1e-3, " ~ ", Bnorma*1e-8)

	fmt.Print("\nВведите α: ")
	fmt.Scan(&alpha)

	solution := tihonov(matrix, b, alpha)

	// Выводим все
	fmt.Println("\nРешение системы:")
	for i, x := range solution {
		fmt.Printf("x%d = %f\n", i+1, x)
	}

}
