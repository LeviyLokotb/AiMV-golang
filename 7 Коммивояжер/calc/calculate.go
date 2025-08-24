package calc

import (
	"fmt"
	. "komivoyager/data"
	. "komivoyager/ways"
	"math"
)

// Вычисление всех путей для заданного города
func Calculate(capital int) []*Way {
	// Функция FindWays рекурсивная с многопоточностью,
	// поэтому здесь используем канал
	ch := make(chan []*Way)
	way := NewWay(capital)
	go FindWays(way, ch, 0)
	return <-ch
}

// ищем все пути точки (рекурсивно)
func FindWays(current *Way, ch chan []*Way, layer int) {
	layer++
	if current.IsEnd() {
		ch <- []*Way{current}
		return
	}
	/*
		// Выводим путь в процессе
		fmt.Print(" :> ")
		for _, city := range current.Path {
			fmt.Printf("%d -> ", city)
		}
		fmt.Println()
	*/
	ways := make([]*Way, 0)
	chans := make([]chan []*Way, 0)
	for i := range TABLE[current.Last()] {
		if current.IsGoodCity(i) {
			neway := CopyWay(current)
			neway.Travelto(i)
			lch := make(chan []*Way)
			chans = append(chans, lch)
			go FindWays(neway, lch, layer)
		}
	}
	for _, lch := range chans {
		res := <-lch
		ways = append(ways, res...)
	}
	/*
		for _, w := range ways {
			fmt.Println((*w).Path)
		}
		fmt.Println()
	*/
	ch <- ways
}

func BestWay(ways []*Way) ([]*Way, int) {
	min := math.Inf(1)
	for _, way := range ways {
		if way.Len < min {
			min = way.Len
		}
	}
	var bests []*Way
	for _, way := range ways {
		if way.Len == min {
			bests = append(bests, way)
		}
	}
	/*
		// Вывод всех путей (отладка)
		fmt.Println("Все пути:")
		for _, way := range ways {
			for _, city := range way.Path {
				fmt.Print(" :> ")
				fmt.Printf("%d -> ", city)
			}
			fmt.Println(way.Path[0])
		}
	*/
	return bests, len(ways)
}

// Выводим информацию из результата
func Formatting(ways []*Way, all_ways int) {
	// Всего путей
	if all_ways == 0 {
		fmt.Println("Не удалось построить пути")
		return
	}
	fmt.Printf("Путей найдено: %d\n", all_ways)
	// Длина
	if len(ways) <= 0 {
		fmt.Println("Подходящих путей не найдено")
		return
	}
	fmt.Printf("Длина лучшего пути: %f\n", ways[0].Len)
	fmt.Printf("Путей такой длины: %d\n", len(ways))
	// Вывод путей
	fmt.Println("Эти пути:")
	for _, way := range ways {
		fmt.Print(" :> ")
		for _, city := range way.Path {

			fmt.Printf("%d -> ", city+1)
		}
		fmt.Println(way.Path[0] + 1)
	}

}
