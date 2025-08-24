package main

import (
	"fmt"
	"komivoyager/calc"
	"komivoyager/data"
)

// Задача Комммивояжёра, метод ветвей и границ
func main() {
	PrintTbl(data.TABLE)
	capital := InputInt("Из какого города начинаем: ") - 1
	// Поиск всех путей -> Выбор лучших -> Вывод информации
	calc.Formatting(calc.BestWay(calc.Calculate(capital)))
}

func InputInt(text string) int {
	fmt.Print(text)
	var tmp int
	fmt.Scan(&tmp)
	return tmp
}

func PrintTbl(t [data.N][data.N]float64) {
	for _, row := range t {
		fmt.Println(row)
	}
}
