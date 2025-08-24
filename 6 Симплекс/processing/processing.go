package processing

import (
	"fmt"
	. "mat-task/data"
	"math"
)

func GetSimplexTable() [][]float64 {
	/*
		var (
			x1 float64 // количество И1 у Б1
			x2 float64 // количество И2 у Б1
			y1 float64 // количество И1 у Б2
			y2 float64 // количество И2 у Б2
			// остаточные переменные (базисные)
			z1 = B1.Worktime // начальное приближение
			z2 = B2.Worktime // начальное приближение
		)
	*/

	/*
		Система::
			Изделия:
				x1 + y1 = TOTAL_I1
				x2 + y2 = TOTAL_I2
			Изделия:
				x1/B1.I1ph + x2/B1.I2ph <= B1.Worktime
				y1/B2.I1ph + y2/B2.I2ph <= B2.Worktime
			Неотрицательность:
				x1, x2, y1, y2 >= 0
			Функция:
				Price() = ... -> min

		Искуственная целевая функция (минимизировать искуственные переменные):
			W() = z1 + z2 -> min

		Приводим:
				x1 + y1 + z3 = TOTAL_I1
				x2 + y2 + z4 = TOTAL_I2
				x1/B1.I1ph + x2/B1.I2ph + z1 = B1.Worktime
				y1/B2.I1ph + y2/B2.I2ph + z2 = B2.Worktime
			x1, x2, y1, y2, z1, z2 >= 0
			Price() = -(...) -> max
			W() = TOTAL_I1 + TOTAL_I2 - x1 - x2 -y1 -y2

	*/

	// Начальное приближение (небазисные = 0, базисные из формул (уже заданы при инициализации) )
	//Solution := []float64{x1, x2, y1, y2, z1, z2}

	/*
		Симплекс таблица

		Базис	|			x1			|			x2			|			y1			|			y2			|	z1	|	z2	|	z3	|	z4	|		Решение		|
		--------+-----------------------+-----------------------+-----------------------+-----------------------+-------+-------+-------+-------+-------------------+
		Price()	|		B1.I1price		|		B1.I2price		|		B2.I1price		|	B2.I2price			|	0	|	0	|	0	|	0	|			0		| // Коэффициенты функции с обратным знаком
		--------+-----------------------+-----------------------+-----------------------+-----------------------+-------+-------+-------+-------+-------------------+
		W()		|			-1			|			-1			|			-1			|			-1			|	0	|	0	|	0	|	0	|TOTAL_I1 + TOTAL_I2| // Коэффициенты функции с обратным знаком
		--------+-----------------------+-----------------------+-----------------------+-----------------------+-------+-------+-------+-------+-------------------+
		z1		|		1/B1.I1ph		|		1/B1.I2ph		|			0			|			0			|	1	|	0	|	0	|	0	|	B1.Worktime		| // Коэффициенты ограничений
		--------+-----------------------+-----------------------+-----------------------+-----------------------+-------+-------+-------+-------+-------------------+
		z2		|			0			|			0			|		1/B2.I1ph		|		1/B2.I2ph		|	0	|	1	|	1	|	1	|	B2.Worktime		| // Коэффициенты ограничений
		--------+-----------------------+-----------------------+-----------------------+-----------------------+-------+-------+-------+-------+------------------+
		z3		|			1			|			0			|			1			|			0			|	0	|	0	|	1	|	0	|	TOTAL_I1		| // Коэффициенты ограничений
		--------+-----------------------+-----------------------+-----------------------+-----------------------+-------+-------+-------+-------+------------------+
		z4		|			0			|			1			|			0			|			1			|	0	|	0	|	0	|	1	|	TOTAL_I2		| // Коэффициенты ограничений
		--------+-----------------------+-----------------------+-----------------------+-----------------------+-------+-------+-------+-------+------------------+

	*/

	SymplexTable := [][]float64{
		{B1.I1price, B1.I2price, B2.I1price, B2.I2price, 0, 0, 0, 0, 0},
		{-1, -1, -1, -1, 0, 0, 0, 0, -TOTAL_I1 - TOTAL_I2},
		{1 / B1.I1ph, 1 / B1.I2ph, 0, 0, 1, 0, 0, 0, B1.Worktime},
		{0, 0, 1 / B2.I1ph, 1 / B2.I2ph, 0, 1, 1, 1, B2.Worktime},
		{1, 0, 1, 0, 0, 0, 1, 0, TOTAL_I1},
		{0, 1, 0, 1, 0, 0, 0, 1, TOTAL_I2},
	}

	return SymplexTable
}

func Solve(base [][]float64) ([][]float64, map[int]int) {
	rows := len(base)
	cols := len(base[0])
	// Copy
	tbl := base

	// Соответствие переменных (номер столбца) и строки с решением
	sols := make(map[int]int)

	for step := 1; ; step++ {

		// Определяем ведущий столбец (максимальный отрицательный коэффициент в строке функции (в нашем случе W))
		// Если несколько одинаковых, будет выбран первый
		maincol, _ := FindMinInCol(tbl[1])
		// Если выбрать другой элемент на 1 шаге (заменить -1 на 1 to enable)
		if step == -1 {
			maincol = 1
		}

		// Вектор симплексных отношений (соответствует ведущему столбцу)
		simplexdiv := make([]float64, rows)

		for i := range tbl {
			if tbl[i][maincol] > 0 {
				simplexdiv[i] = tbl[i][cols-1] / tbl[i][maincol]
			} else {
				// Если коэффициент не положтельный ставим бесконечность (строка точно не станет ведущей)
				simplexdiv[i] = math.Inf(1)
			}
		}

		// Ищем ведущую строку
		mainrow, _ := FindMinInRow(simplexdiv)

		sols[maincol] = mainrow
		fmt.Printf("\nШаг #%d\n", step)
		fmt.Println("Ведущий элемент: ", tbl[mainrow][maincol])
		fmt.Println(maincol, " : ", mainrow)

		main_el := tbl[mainrow][maincol]
		// Нормализуем ведущую строку по ведущему элементу
		for j := range tbl[mainrow] {
			tbl[mainrow][j] /= main_el
		}

		//PrintMx(tbl)
		fmt.Println("=====")
		// Остальные меняем по правилу прямоугольника
		var cp [][]float64
		for i := range tbl {
			if i == mainrow {
				cp = append(cp, tbl[i])
				continue
			}
			neo := make([]float64, cols)
			for j := range tbl[i] {
				if j == maincol {
					neo[j] = tbl[i][j]
					continue
				}
				// Забавно, в коде описывается гораздо проще чем через мутные аналогии с прямоугольниками
				neo[j] = tbl[i][j]*tbl[mainrow][maincol] - tbl[mainrow][j]*tbl[i][maincol]
				//neo[j] /= tbl[mainrow][maincol]
			}
			cp = append(cp, neo)
		}
		tbl = cp

		// Заменяем элементы ведущего столбца на 0
		for i := range tbl {
			if i == mainrow {
				continue
			}
			tbl[i][maincol] = 0
		}

		PrintMx(tbl)
		// Проверяем решение
		if IsSolution(tbl, sols) {
			break
		}
	}

	return tbl, sols
}

func ExtractSolution(tbl [][]float64, sols map[int]int) map[int]float64 {
	result := make(map[int]float64)
	l := len(tbl[0])
	for key, val := range sols {
		result[key] = tbl[val][l-1]
	}
	for i := 0; i < 4; i++ {
		if result[i] == 0 {
			result[i] = 0
		}
	}

	return result
}

//////////////////////////////////////////////////////////////////////////////

// Функция для максимизации (-затраты) -> max
func Price(x1, x2, y1, y2 float64) float64 {
	return -(B1.I1price*x1 + B1.I2price*x2 + B2.I1price*y1 + B2.I2price*y2)
}

// -1* Искуственная целевая функция -> max
func W(x1, x2, y1, y2 float64) float64 {
	//-(z1+z2)
	return -(TOTAL_I1 + TOTAL_I2) + x1 + x2 + y1 + y2
}

func TestODR(x1, x2, y1, y2, z1, z2 float64) bool {

	o1 := (x1+x2+z1 == B1.Worktime)
	o2 := (y1+y2+z2 == B2.Worktime)
	o3 := (B1.I1ph*x1+B2.I1ph == TOTAL_I1)
	o4 := (B1.I2ph*x1+B2.I2ph == TOTAL_I2)
	o5 := (x1 >= 0 && x2 >= 0 && y1 >= 0 && y2 >= 0 && z1 >= 0 && z2 >= 0)
	return o1 && o2 && o3 && o4 && o5
}
