package ways

import (
	"komivoyager/data"
)

// структура, описывающая путь
type Way struct {
	Path []int     // Полный путь
	Dist []float64 // Массив длин путей
	Len  float64   // Полная длина пути
}

// Начало пути (индекс стартового города)
func NewWay(start int) *Way {
	w := new(Way)
	w.Path = append(w.Path, start)
	w.Len = 0
	return w
}

func CopyWay(from *Way) *Way {
	w := new(Way)
	w.Path = from.Path
	w.Dist = from.Dist
	w.Len = from.Len
	return w
}

// Добавить перемещение к городу с индексом
func (way *Way) Travelto(target int) float64 {
	distance := way.GetDist(target)
	way.Path = append(way.Path, target)
	way.Dist = append(way.Dist, distance)
	way.Len += distance
	return distance
}

func (way *Way) IsGoodCity(target int) bool {
	// Условие: путь существует и это не тот же город (длина пути 0)
	if way.GetDist(target) <= 0 {
		return false
	}

	// Условие: город ранее не посещался
	for _, city := range way.Path {
		if city == target {
			return false
		}
	}

	return true
}

func (way *Way) GetDist(target int) float64 {
	return data.TABLE[way.Last()][target]
}

func (way *Way) Last() int {
	return way.Path[len(way.Path)-1]
}

func (way *Way) IsEnd() bool {
	return len(way.Path) == data.N
}
