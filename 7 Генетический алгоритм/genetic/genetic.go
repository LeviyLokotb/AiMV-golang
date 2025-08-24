package genetic

import (
	"fmt"
	"genetic/data"
)

func BattleRoyale() *Genom {
	// Первые виды
	var firsts []*Genom
	for i := 0; i < data.FIRSTS; i++ {
		firsts = append(firsts, NewGenom())
	}

	var best = NewGenom()
	best.fitness = -1
	for generation := 1; generation <= data.GENERATIONS_MAX; generation++ {
		fmt.Println("gen ", generation, ": ")
		data.Generations_Counter++
		// Скрещиваем
		var next_gen []*Genom
		for f, first := range firsts {
			for s, second := range firsts {
				if f == s {
					continue
				}
				next_gen = append(next_gen, Crossover(first, second)...)
			}
		}
		// Пусть в каждом поколении в род приходят новые особи
		for i := 0; i < data.OUTSIDE_STRANGERS; i++ {
			next_gen = append(next_gen, NewGenom())
		}
		/*
			fmt.Println("[*]")
			for _, g := range next_gen {
				fmt.Println(g.GetInfo())
			}
			fmt.Println("[*]")
		*/
		// Сортируем и убираем мёртвых
		next_gen = BubbleSortGen(next_gen)

		// Достаём лучшего
		for _, gen := range next_gen {
			if gen.fitness > best.fitness {
				best = gen
			}
		}

		// Возвращаем либо переходим на следующий цикл
		// Все вымерли
		if len(next_gen) == 0 {
			return nil
		}
		// Остался один - скрещивается сам с собой (он всё равно лучший, авось мутация прокнет)
		if len(next_gen) < data.FIRSTS {
			firsts = next_gen[0 : len(next_gen)-1]
			continue
		}
		// Скрещиваем лучших
		firsts = next_gen[0:data.FIRSTS]
	}
	return best
}
