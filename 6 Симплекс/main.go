package main

import (
	"fmt"
	"mat-task/data"
	"mat-task/processing"
)

func main() {
	fmt.Println(data.TZ)

	fmt.Println("\nМатрица симплекс-таблицы:")
	base := processing.GetSimplexTable()
	processing.PrintMx(base)

	solution, mp := processing.Solve(base)
	fmt.Println("\nРешённая таблица:")
	processing.PrintMx(solution)

	tru_solution := processing.ExtractSolution(solution, mp)
	fmt.Println("===\nРешение:")
	for i := 0; i < 4; i++ {
		fmt.Printf("x%d = %f\n", i+1, tru_solution[i])
	}
	fmt.Println("===\nЗатраты: ", -processing.Price(tru_solution[0], tru_solution[1], tru_solution[2], tru_solution[3]))
	fmt.Println("===\nx1 - деталей И1 сделает Б1\nx2 - деталей И2 сделает Б1\nx3 - деталей И1 сделает Б2\nx4 - деталей И2 сделает Б2")
}

///////////////////////////////////////////////
