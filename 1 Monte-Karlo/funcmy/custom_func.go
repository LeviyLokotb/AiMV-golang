package funcmy

import (
	"log"
	"math"
)

////////////////////////////////////////////////
/* config: */

func F(x float64) (y float64) {
	sqr := x * x
	// здесь задаётся функция
	y = math.Log(x+math.Sqrt(sqr-0.25)) / (2 * sqr)
	// y = math.Log((x+math.Sqrt(sqr-0.25)) / (2 * sqr))
	//y = -sqr + 1
	return
}

var A float64 = 0.5
var B float64 = 1.7
var EPS float64 = 1e-5

///////////////////////////////////////////////

func IsPointInside(x, y float64) (bool, int) {
	F := F(x)
	// функция не определена в точке
	if math.IsNaN(F) {
		return false, 0
	}
	// определяем принадлежит ли точка (учитываем > < 0)
	if y > 0 {
		if y <= F {
			return true, 1
		} else {
			return false, 0
		}
	} else if y < 0 {
		if y >= F {
			return true, -1
		} else {
			return false, 0
		}

	}
	return true, 1
}

// приблизительный максимум
func MaxY() float64 {
	max := math.Inf(-1)
	for i := A; i <= B; i += 0.01 {
		f := F(i)
		if math.IsNaN(f) {
			log.Println("WARNING: infinity in y\n\t| max set in 100")
			return 100
		}
		if f > max {
			max = f
		}
	}
	// округляем до большего целого
	//log.Println("max: ", max)
	return max * 1.1
}

// приближённый минимум
func MinY() float64 {
	min := math.Inf(1)
	for i := A; i <= B; i += 0.01 {
		f := F(i)
		if math.IsNaN(f) {
			log.Println("WARNING: infinity in y\n\t| min set in -100")
			return -100
		}
		if f < min {
			min = f
		}
	}
	// округляем до меньшего целого
	//log.Println("min: ", min)
	return min * 1.1
}

func SetA(neo float64) {
	A = neo
}
func SetB(neo float64) {
	B = neo
}
func SetEPS(neo float64) {
	EPS = neo
}
