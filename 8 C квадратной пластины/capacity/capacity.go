package capacity

import (
	"ecapacity/data"
	"ecapacity/tihonov"
	"math"
)

func ECapacityPlate() float64 {
	// Размер ячейки (-a до a)
	h := 2 * data.A / float64(data.N)
	area := h * h

	// Координаты центров
	x := make([]float64, data.N)
	y := make([]float64, data.N)
	for i := 0; i < data.N; i++ {
		x[i] = -data.A + h*(float64(i)+0.5)
		y[i] = -data.A + h*(float64(i)+0.5)
	}

	// Коэффициент для диагональных элементов
	diagonal := 4 * math.Log(math.Tan(3*math.Pi/8)) * h

	// Строим матрицу A системы (размером N*N на N*N)
	size := data.N * data.N
	A := make([][]float64, size)
	for i := range A {
		A[i] = make([]float64, size)
	}

	for i := 0; i < data.N; i++ {
		for j := 0; j < data.N; j++ {
			// Получаем индекс строки
			row := i*data.N + j
			for k := 0; k < data.N; k++ {
				for l := 0; l < data.N; l++ {
					col := k*data.N + l
					dx := x[i] - x[k]
					dy := y[j] - y[l]
					delta := 1e-6 // Небольшое значение чтобы не получить div by 0
					r := math.Sqrt(dx*dx + dy*dy + delta*delta)

					// Диагональные коэффициенты вычисяются аналитически (заранее)
					if row == col || row == size-col {
						A[row][col] = diagonal
						continue
					}
					// Остальные - по приближённой формуле
					A[row][col] = (area / r) * (1 + (area / (24 * r * r)))
					//A[row][col] = area / (4 * math.Pi * data.E0 * r)
				}
			}
		}
	}
	//tihonov.PrintMatrix(A)
	// Вектор правой части (везде V)
	B := make([]float64, size)
	for i := range B {
		B[i] = data.V
	}

	// Решение СЛАУ A * X = B
	// Используем метод Гаусса-Джордана с регуляризацией Тихонова (из 2 темы)
	sols := tihonov.TihonovReg(A, B, data.ALPHA)
	//tihonov.PrintVector(sols)
	// Вычисление полного заряда их всех решений
	Q := 0.0
	for _, sol := range sols {
		Q += sol * area
	}

	// Электроёмкость
	C := Q / data.V

	return C
}
