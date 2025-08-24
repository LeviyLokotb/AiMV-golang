package tihonov

import (
	"fmt"
	"math"
	"sync"
)

// Решение методом Гаусса-Джордана
func GaussJordan(a [][]float64, b []float64) []float64 {
	//PrintMatrix(a)
	// Объединяем матрицы в прямоугольную
	matrix := make([][]float64, 0)
	for i := 0; i < len(a); i++ {
		neo := a[i]
		neo = append(neo, b[i])
		matrix = append(matrix, neo)
	}
	//PrintMatrix(matrix)

	rows := len(matrix)
	cols := len(matrix[0])
	//PrintMatrix(matrix)

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
		//PrintMatrix(matrix)

		var wg sync.WaitGroup
		// Вычитание текущей строки из нижних строк
		wg.Add(rows - row - 1)
		for r := row + 1; r < rows; r++ {
			go func(r, col int) {
				defer wg.Done()
				// число на которое домножаем перед вычитанием
				multi := matrix[r][col]
				for c := 0; c < cols; c++ {
					// первый элемент должен обнулиться
					matrix[r][c] -= multi * matrix[row][c]
				}
			}(r, col)
		}
		wg.Wait()
		//PrintMatrix(matrix)
	}
	//PrintMatrix(matrix)

	// Обратный ход
	// Массив решений
	solution := make([]float64, cols-1)

	var wg sync.WaitGroup

	// идём с конца
	wg.Add(rows)
	for i := rows - 1; i >= 0; i-- {
		// res = b
		go func(i int) {
			defer wg.Done()
			res := matrix[i][cols-1]
			for k := i + 1; k < cols-1; k++ {
				// Вычитаем a*x (там где x не вычислен будет 0, так и задумано)
				res -= solution[k] * matrix[i][k]
			}
			solution[i] = res
		}(i)
	}
	wg.Wait()
	//PrintVector(solution)
	return solution
}

func TihonovReg(a [][]float64, b []float64, alpha float64) []float64 {
	//PrintMatrix(a)
	fmt.Println()
	size := len(b)
	// Вычисление AtA
	AtA := make([][]float64, size)
	for i := range AtA {
		AtA[i] = make([]float64, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				AtA[i][j] += a[k][i] * a[k][j]
			}
			if i == j {
				AtA[i][j] += alpha
			}
		}
	}

	// Вычисление Atb
	Atb := make([]float64, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			Atb[i] += a[j][i] * b[j]
		}
	}
	return GaussJordan(AtA, Atb)
}
