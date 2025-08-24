package main

import (
	"RSImin/fx"
	"RSImin/graphdraw"
	"fmt"
	"math"
)

func RSImin(minX, maxX, eps float64) (float64, float64) {
	var minY, maxY, steps float64
	cX := (maxX + minX) / 2
	for count := 0; count < 100; count++ {
		steps++

		minY = fx.F(minX)
		//fmt.Println("minY: ", minY)
		maxY = fx.F(maxX)
		//fmt.Println("maxY: ", maxY)

		cY := fx.F(cX)

		// вершина параболы, проходящей через 3 точки (minX, cX, maxX)
		newX := (cY*maxY*minX)/((minY-cY)*(minY-maxY)) + (minY*maxY*cX)/((cY-minY)*(cY-maxY)) + (minY*cY*maxX)/((maxY-minY)*(maxY-cY))
		// Избегааем /0
		if math.IsNaN(newX) {
			maxX += 0.001
			minX -= 0.002
			steps--
			continue
		}
		newF := fx.F(newX)

		if newF < minY && newF < maxY {
			// Если функция в найденном x меньше чем в других точках, устанавливаем как максимум (3 точка)
			minX, cX, maxX = cX, maxX, newX
		} else {
			// иначе это минимум (1 точка)
			minX, cX, maxX = newX, minX, cX
		}

		// остановка
		if math.Abs(maxX-minX) < eps {
			break
		}
	}

	return math.Abs(maxX-minX) / 2, steps
}

func main() {
	// метод дихотомии

	// Рисуем функцию
	draw1 := graphdraw.DrawFunction()

	fmt.Println("===============\nОбратная квадратичная интерполяция для поиска минимума\n===============")
	// Считаем
	min, count := RSImin(fx.A, fx.B, fx.Eps)

	fmt.Print("\na: ", fx.A)
	fmt.Print("\nb: ", fx.B)
	fmt.Print("\nМинимум функции: ", min, " : ", fx.F(min))

	fmt.Print("\nДостигнуто за ", count, " шагов.")

	graphdraw.DrawXmark(min, draw1, 255, 0, 0)

	fmt.Print("\n\n")
	graphdraw.Save(draw1, "RSIforMin")

}
