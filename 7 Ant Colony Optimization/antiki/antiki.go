package antiki

import (
	. "antopt/data"
	"math"
	"math/rand"
)

type Solution struct {
	X1    int
	X2    int
	Value int
}

func Calc() Solution {
	// Решения
	sols := make([]Solution, 0)
	best := Solution{0, 0, -1}
	// Уровень "феромона" для каждого значения (больше фкромона -> больше вероятность выбора)
	// Начальный уровень = 1 для всех
	var X1Feromon []float64
	for range X1s {
		X1Feromon = append(X1Feromon, 1)
	}
	var X2Feromon []float64
	for range X2s {
		X2Feromon = append(X2Feromon, 1)
	}
	// Начало алгоритма
	for it := 0; it < ITER; it++ {
		for ant := 0; ant < ANTS; ant++ {
			// Получаем случайное (с влиянием феромона) решение
			i1 := GetRandomValue(X1Feromon, FeromonPOWER)
			i2 := GetRandomValue(X2Feromon, FeromonPOWER)

			x1, x2 := X1s[i1], X2s[i2]
			//fmt.Println(x1, " : ", x2)
			if !IsGoodSol(x1, x2) {
				continue
			}
			cur := Solution{x1, x2, F(x1, x2)}
			sols = append(sols, cur)
			// Получаем лучшее решение
			if cur.Value > best.Value {
				best = cur
			}
		}
		// Добавляем феромон для удачных решений прапорционально значению
		// (-> max, больше феромона за большие значения функции)
		for _, sol := range sols {
			X1Feromon[sol.X1] += FeromonUP * float64(sol.Value)
			X2Feromon[sol.X2] += FeromonUP * float64(sol.Value)
		}
		// Испаряем феромон (стирание неэффективных решений)
		for i := range X1Feromon {
			X1Feromon[i] *= (1 - FeromonDOWN)
		}

		for i := range X2Feromon {
			X2Feromon[i] *= (1 - FeromonDOWN)
		}
	}
	return best
}

// Выбор случайного значения на основе феромона
func GetRandomValue(feromon []float64, power float64) int {
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	// Вероятности
	probabilities := make([]float64, len(feromon))
	total := 0.0
	for i := range probabilities {
		probabilities[i] = math.Pow(feromon[i], power)
		total += probabilities[i]
	}
	// Нормализуем
	for i := range probabilities {
		probabilities[i] /= total
	}
	// Генерируем и получаем результат
	r := rand.Float64()
	//fmt.Println(r)
	sum := 0.0
	for i := 0; i < len(probabilities)-1; i++ {
		sum += probabilities[i]
		if r > sum && r < sum+probabilities[i+1] {
			return i
		}
	}
	return len(probabilities) - 1
}
