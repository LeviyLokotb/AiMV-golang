package main

import (
	"antopt/antiki"
	"antopt/data"
	"fmt"
)

func main() {
	sol := antiki.Calc()
	fmt.Println("Лучшее значение функции: ", sol.Value)
	fmt.Printf("Решение:\nx1=%d\nx2=%d\n\n", sol.X1, sol.X2)
	fmt.Printf("* Достигнуто за %d итераций, силами %d муравьёв", data.ITER, data.ANTS)
}
