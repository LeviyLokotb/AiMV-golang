package main

import (
	"ecapacity/capacity"
	"ecapacity/data"
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Print("\n\n-=[ Вычисление электроёмкости квадратной пластины ]=-\n\n")
	fmt.Println("Введите новое значение, или 'd' (default) для значения по умолчанию")

	SetSome("Сторона пластины (м): ", data.SetA, data.A*2)

	SetSome("Число разбиений: ", data.SetN, data.N)
	fmt.Println("  (", data.N*data.N, " ячеек)")

	SetSome("Потенциал на поверхности пластины (В): ", data.SetV, data.V)

	SetSome("Параметр регуляризации Тихонова: ", data.SetALPHA, data.ALPHA)

	C := capacity.ECapacityPlate()
	fmt.Printf("Результат:\n* Электроёмкость C = %f  пФ\n", C)
	C *= 1e-10
	Ccircle := 16 * data.E0 * data.A / math.SqrtPi
	fmt.Printf("Электроёмкость равновеликого круга: %f пФ\n", Ccircle*1e9)
	fmt.Println("Отношение к электроёмкости равновеликого круга: ", 1/(C/Ccircle))

	prompt := "!"
	fmt.Println("\n'q' or 'exit' to quit")
	for ; ; fmt.Scan(&prompt) {
		switch prompt {
		case "q", "quit", "exit", "\n":
			return
		}
	}
}

//

func SetSome(title string, Set func(float64), els interface{}) {
	fmt.Print(title)
	var sA string
	for {
		fmt.Scan(&sA)
		if sA == "d" {
			fmt.Println(title, els)
			return
		} else {
			A, err := strconv.ParseFloat(sA, 64)
			if err == nil {
				Set(A)
				return
			}
		}
	}
}
