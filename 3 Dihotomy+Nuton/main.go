package main

import (

	//"dihonuton/fx"
	"dihonuton/fx"
	"dihonuton/graphdraw"
	"fmt"
	"math"
)

func equalSign(a, b float64) bool {
	return a > 0 && b > 0 || a < 0 && b < 0
}

var dihotomy_counter int

func dihotomy(a, b, n float64) (float64, float64, float64) {
	dihotomy_counter++
	n *= 2.0

	centerX := (a + b) / 2
	centerY := fx.F(centerX)
	//fmt.Println("cY = ", centerY)
	aY := fx.F(a)
	//fmt.Println("aY = ", aY)
	//fmt.Println("bY = ", bY)
	// Проверяем епсилон (длина отрезка)
	accurate := (b - a) / n
	if accurate <= fx.Eps2 {
		return a, b, centerX
	}

	// сразу проверяем на 0
	if centerY == 0 {
		return a, b, centerX
	}
	// если знаки на I половине одинаковые, ищем на II, иначе на I
	if !equalSign(aY, centerY) {
		return dihotomy(a, centerX, n)
	} else {
		return dihotomy(centerX, b, n)
	}
}

func nuton(x0 float64) float64 {
	x := x0
	for {
		f := fx.F(x)
		dx := fx.DF(x)

		if dx == 0 {
			panic("Производная равна 0, ошибка")
		}

		neo := x - f/dx

		//graphdraw.DrawKasLine()

		if math.Abs(x-neo) <= fx.Eps {
			return neo
		}

		x = neo
	}
}

func main() {
	// метод дихотомии

	// Рисуем функцию
	draw1 := graphdraw.DrawFunction()

	fmt.Println("===============\nМетод дихотомии\n===============")
	// Считаем
	dihotomy_counter = 0
	_, _, x1 := dihotomy(fx.A, fx.B, 2)
	fmt.Println("Количество разбиений: ", dihotomy_counter)
	// Выводим решения
	fmt.Println("\nx = ", x1)
	fmt.Print("\n\n")
	graphdraw.Save(draw1, "dihotomy")

	// Метод Ньютона

	draw2 := graphdraw.DrawFunction()

	fmt.Println("===============\nМетод Ньютона\n===============")
	fmt.Println("Начальное приближение x0 = ", x1)
	x2 := nuton( /*fx.X0*/ x1)

	graphdraw.DrawKasLine(x2, draw2)
	graphdraw.DrawPoint(x2, 0, draw2, 255, 128, 0)
	fmt.Println("\nx = ", x2)

	graphdraw.Save(draw2, "nuton")
}
