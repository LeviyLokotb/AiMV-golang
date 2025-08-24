package main

import (
	"fmt"
	"genetic/data"
	"genetic/genetic"
)

func main() {
	x1, x2, fit := genetic.BattleRoyale().GetInfo()
	if fit == -1 {
		fmt.Println("Все умерли. Смертей всего: ", data.Death_Counter)
		fmt.Println("Попробуйте снова.")
		//return
	}
	fmt.Printf("Решение:\nx1 = %d\nx2 = %d\nРешение: %d\n\n", x1, x2, fit)
	fmt.Printf("* Статистика:\nОсобей: %d\nПоколений: %d\nМутаций: %d\nСмертей: %d\n\n", data.Creature_Counter, data.Generations_Counter, data.Mutate_Counter, data.Death_Counter)
}
